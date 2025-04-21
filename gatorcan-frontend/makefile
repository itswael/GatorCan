# Makefile for Gatorcan Frontend

# Config
DOCKERFILE = Dockerfile.gatorcan_frontend
IMAGE_NAME = gatorcan-frontend
CONTAINER_NAME = gatorcan-frontend-container
PORT = 3000

# Load env vars from .env file
include .env
export

.PHONY: build run stop clean

build:
	@echo "Building Docker image: $(IMAGE_NAME)"
	docker build \
		-f $(DOCKERFILE) \
		--build-arg VITE_AWS_REGION=$(VITE_AWS_REGION) \
		--build-arg VITE_AWS_ACCESS_KEY_ID=$(VITE_AWS_ACCESS_KEY_ID) \
		--build-arg VITE_AWS_SECRET_ACCESS_KEY=$(VITE_AWS_SECRET_ACCESS_KEY) \
		--build-arg VITE_AWS_SESSION_TOKEN=$(VITE_AWS_SESSION_TOKEN) \
		-t $(IMAGE_NAME) .

run:
	@echo "Running container: $(CONTAINER_NAME)"
	docker run -d --name $(CONTAINER_NAME) \
		-p $(PORT):80 \
		$(IMAGE_NAME)

stop:
	@echo "Stopping container: $(CONTAINER_NAME)"
	docker stop $(CONTAINER_NAME) || true

clean: stop
	@echo "Removing image: $(IMAGE_NAME)"
	docker rmi $(IMAGE_NAME) || true
