package main

import (
	"context"
	"net/url"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

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
	code := req.PathParameters["shortenCode"]
	if code == "" {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Headers: map[string]string{
				"access-control-allow-origin":  "*",
				"access-control-allow-headers": "content-type",
				"access-control-allow-methods": "OPTIONS,GET",
			},
			Body: "missing shortenCode",
		}, nil
	}

	out, err := h.db.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(h.tableName),
		Key: map[string]types.AttributeValue{
			"shortCode": &types.AttributeValueMemberS{Value: code},
		},
		ConsistentRead: aws.Bool(true),
	})
	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"access-control-allow-origin":  "*",
				"access-control-allow-headers": "content-type",
				"access-control-allow-methods": "OPTIONS,GET",
			},
			Body: "failed to load short url",
		}, nil
	}
	if out.Item == nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 404,
			Headers: map[string]string{
				"access-control-allow-origin":  "*",
				"access-control-allow-headers": "content-type",
				"access-control-allow-methods": "OPTIONS,GET",
			},
			Body: "not found",
		}, nil
	}

	longAttr, ok := out.Item["longUrl"].(*types.AttributeValueMemberS)
	if !ok || longAttr.Value == "" {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"access-control-allow-origin":  "*",
				"access-control-allow-headers": "content-type",
				"access-control-allow-methods": "OPTIONS,GET",
			},
			Body: "invalid stored url",
		}, nil
	}
	if _, err := url.Parse(longAttr.Value); err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"access-control-allow-origin":  "*",
				"access-control-allow-headers": "content-type",
				"access-control-allow-methods": "OPTIONS,GET",
			},
			Body: "invalid stored url",
		}, nil
	}

	_ = h.bestEffortUpdate(ctx, code)

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 302,
		Headers: map[string]string{
			"location":                     longAttr.Value,
			"cache-control":                "no-store",
			"access-control-allow-origin":  "*",
			"access-control-allow-headers": "content-type",
			"access-control-allow-methods": "OPTIONS,GET",
		},
		Body: "",
	}, nil
}

func (h *handler) bestEffortUpdate(ctx context.Context, code string) error {
	now := time.Now().UTC().Format(time.RFC3339)
	_, err := h.db.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(h.tableName),
		Key: map[string]types.AttributeValue{
			"shortCode": &types.AttributeValueMemberS{Value: code},
		},
		UpdateExpression: aws.String("ADD hits :one SET lastVisited = :now"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":one": &types.AttributeValueMemberN{Value: "1"},
			":now": &types.AttributeValueMemberS{Value: now},
		},
	})
	return err
}
