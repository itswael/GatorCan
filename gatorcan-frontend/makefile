# Makefile for Gatorcan Frontend

# Config
DOCKERFILE = Dockerfile.gatorcan_frontend
IMAGE_NAME = gatorcan-frontend
CONTAINER_NAME = gatorcan-frontend-container
PORT = 3000
ECR_REPO = gatorcan
ECR_TAG = frontend

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

push-to-ecr:
	@echo "Pushing image to ECR"
	aws ecr get-login-password --region us-east-2 | docker login --username AWS --password-stdin 335257738167.dkr.ecr.us-east-2.amazonaws.com
	docker tag $(IMAGE_NAME) 335257738167.dkr.ecr.us-east-2.amazonaws.com/$(ECR_REPO):$(ECR_TAG)
	docker push 335257738167.dkr.ecr.us-east-2.amazonaws.com/$(ECR_REPO):$(ECR_TAG)

run-from-ecr:
	@echo "Running container from ECR"
	aws ecr get-login-password --region us-east-2 | docker login --username AWS --password-stdin 335257738167.dkr.ecr.us-east-2.amazonaws.com
	docker pull 335257738167.dkr.ecr.us-east-2.amazonaws.com/$(ECR_REPO):$(ECR_TAG)
	docker run -d --name $(CONTAINER_NAME) \
		-p $(PORT):80 \
		335257738167.dkr.ecr.us-east-2.amazonaws.com/$(ECR_REPO):$(ECR_TAG)
