package main

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type requestBody struct {
	URL string `json:"url"`
}

type responseBody struct {
	ShortCode string `json:"shortCode"`
	ShortURL  string `json:"shortUrl"`
	LongURL   string `json:"longUrl"`
}

type item struct {
	ShortCode   string `dynamodbav:"shortCode"`
	LongURL     string `dynamodbav:"longUrl"`
	CreatedAt   string `dynamodbav:"createdAt"`
	Hits        int64  `dynamodbav:"hits"`
	LastVisited string `dynamodbav:"lastVisited"`
}

type handler struct {
	db        *dynamodb.Client
	tableName string
}

func main() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic(err)
	}
	fn := &handler{
		db:        dynamodb.NewFromConfig(cfg),
		tableName: os.Getenv("TABLE_NAME"),
	}
	if fn.tableName == "" {
		panic("TABLE_NAME is required")
	}

	lambda.Start(fn.handle)
}

func (h *handler) handle(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var body requestBody
	if err := json.Unmarshal([]byte(req.Body), &body); err != nil {
		return jsonResp(400, map[string]string{"error": "invalid JSON body"}), nil
	}

	longURL, err := validateURL(body.URL)
	if err != nil {
		return jsonResp(400, map[string]string{"error": err.Error()}), nil
	}

	baseURL := strings.TrimSpace(os.Getenv("BASE_URL"))
	if baseURL == "" {
		baseURL = inferBaseURL(req)
	}
	baseURL = strings.TrimRight(baseURL, "/")

	now := time.Now().UTC().Format(time.RFC3339)

	for i := 0; i < 10; i++ {
		code, err := generateCode(7)
		if err != nil {
			return jsonResp(500, map[string]string{"error": "failed to generate short code"}), nil
		}

		it := item{
			ShortCode:   code,
			LongURL:     longURL,
			CreatedAt:   now,
			Hits:        0,
			LastVisited: "",
		}
		av, err := attributevalue.MarshalMap(it)
		if err != nil {
			return jsonResp(500, map[string]string{"error": "failed to serialize item"}), nil
		}

		_, err = h.db.PutItem(ctx, &dynamodb.PutItemInput{
			TableName:           aws.String(h.tableName),
			Item:                av,
			ConditionExpression: aws.String("attribute_not_exists(shortCode)"),
		})
		if err != nil {
			var cfe *types.ConditionalCheckFailedException
			if errors.As(err, &cfe) {
				continue
			}
			return jsonResp(500, map[string]string{"error": "failed to store short url"}), nil
		}

		shortURL := fmt.Sprintf("%s/%s", baseURL, code)
		return jsonResp(201, responseBody{ShortCode: code, ShortURL: shortURL, LongURL: longURL}), nil
	}

	return jsonResp(503, map[string]string{"error": "could not allocate a unique short code"}), nil
}

func validateURL(raw string) (string, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", errors.New("url is required")
	}
	u, err := url.Parse(raw)
	if err != nil {
		return "", errors.New("invalid url")
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return "", errors.New("url scheme must be http or https")
	}
	if u.Host == "" {
		return "", errors.New("url host is required")
	}
	return u.String(), nil
}

func inferBaseURL(req events.APIGatewayV2HTTPRequest) string {
	proto := "https"
	if v := req.Headers["x-forwarded-proto"]; v != "" {
		proto = v
	}
	host := req.Headers["host"]
	if host == "" {
		host = req.RequestContext.DomainName
	}
	if host == "" {
		return ""
	}
	return fmt.Sprintf("%s://%s", proto, host)
}

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func generateCode(n int) (string, error) {
	if n <= 0 {
		return "", errors.New("invalid code length")
	}
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	out := make([]byte, n)
	for i := 0; i < n; i++ {
		out[i] = alphabet[int(b[i])%len(alphabet)]
	}
	return string(out), nil
}

func jsonResp(status int, v any) events.APIGatewayV2HTTPResponse {
	b, _ := json.Marshal(v)
	return events.APIGatewayV2HTTPResponse{
		StatusCode: status,
		Headers: map[string]string{
			"content-type":                "application/json",
			"access-control-allow-origin":  "*",
			"access-control-allow-headers": "content-type",
			"access-control-allow-methods": "OPTIONS,POST",
		},
		Body: string(b),
	}
}
