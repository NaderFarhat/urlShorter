build-ShortenFunction:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o $(ARTIFACTS_DIR)/bootstrap ./cmd/shorten

build-RedirectFunction:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o $(ARTIFACTS_DIR)/bootstrap ./cmd/redirect
