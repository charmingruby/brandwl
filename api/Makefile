DOCKER_IMAGE_NAME := brandwl

###################
# Docker          #
###################
.PHONY: docker-build
docker-build:
	docker build -t ${DOCKER_IMAGE_NAME} .

.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 ${DOCKER_IMAGE_NAME}

###################
# App             #
###################
.PHONY: run
run:
	go run cmd/api/main.go

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/brandwl ./cmd/api/main.go