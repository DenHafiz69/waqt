# Variables
IMAGE_NAME := waqt:latest
CONTAINER_NAME := waqt-app
DOCKERFILE := deployments/Dockerfile
PORT := 8080

.PHONY: help build run stop clean restart logs

help: ## Show available commands
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2}'

build: ## Build the Docker image
	docker build -f $(DOCKERFILE) -t $(IMAGE_NAME) .

run: ## Run the container with mounted .env
	docker run -d \
		--name $(CONTAINER_NAME) \
		-p $(PORT):$(PORT) \
		-v $$(pwd)/.env:/app/.env:ro \
		$(IMAGE_NAME)

stop: ## Stop and remove the running container
	-docker stop $(CONTAINER_NAME)
	-docker rm $(CONTAINER_NAME)

logs: ## Tail container logs
	docker logs -f $(CONTAINER_NAME)

restart: stop run ## Stop, re-run the container

clean: stop ## Stop container and delete the image
	-docker rmi $(IMAGE_NAME)
