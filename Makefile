###################
# App             #
###################
.PHONY: run
run:
	go run cmd/api/main.go

.PHONY: docs
docs:
	swag init -g swagger.go -d ./internal/infra/transport -o ./docs --parseDependency --parseInternal

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/brandwl ./cmd/api/main.go