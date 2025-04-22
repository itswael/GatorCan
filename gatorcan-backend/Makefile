# Include the .env file, for automatic environment variable loading
-include .env

IMAGE_NAME = gatorcan_backend
IMAGE_TAG = main
ECR_REPO = gatorcan
ECR_TAG = backend

# Build the Docker image with environment variables passed as build args
build:
	docker build \
		--build-arg ADMIN_USER=$(ADMIN_USER)\
		--build-arg ADMIN_PASSWORD=$(ADMIN_PASSWORD)\
		--build-arg AWS_REGION=$(AWS_REGION)\
		--build-arg AWS_TOPIC_ARN=$(AWS_TOPIC_ARN)\
		--build-arg AWS_ACCESS_KEY_ID=$(AWS_ACCESS_KEY_ID)\
		--build-arg AWS_SECRET_ACCESS_KEY=$(AWS_SECRET_ACCESS_KEY)\
		--build-arg AWS_SESSION_TOKEN=$(AWS_SESSION_TOKEN)\
		-t muthu1729/muthu:$(IMAGE_NAME)_$(IMAGE_TAG)\
		-f Dockerfile.$(IMAGE_NAME) .


run:
# Run the container with .env file (at runtime)
	docker run -t -p 8080:8080 --name $(IMAGE_NAME)_$(IMAGE_TAG) muthu1729/muthu:$(IMAGE_NAME)_$(IMAGE_TAG)
	

save:
	docker save -o $(IMAGE_TAG).tar $(IMAGE_NAME):$(IMAGE_TAG)

push:
	docker push muthu1729/muthu:$(IMAGE_NAME)_$(IMAGE_TAG)

clean:
	docker rm -f $(IMAGE_NAME)_$(IMAGE_TAG)
	docker rmi muthu1729/muthu:$(IMAGE_NAME)_$(IMAGE_TAG)

pull:
	docker pull muthu1729/muthu:$(IMAGE_NAME)_$(IMAGE_TAG)

test:
	docker run --env-file .env -p 8080:8080 --rm muthu1729/muthu:$(IMAGE_NAME)_$(IMAGE_TAG)

push-to-ecr:
	aws ecr get-login-password --region us-east-2 | docker login --username AWS --password-stdin 335257738167.dkr.ecr.us-east-2.amazonaws.com
	docker tag muthu1729/muthu:$(IMAGE_NAME)_$(IMAGE_TAG) 335257738167.dkr.ecr.us-east-2.amazonaws.com/$(ECR_REPO):$(ECR_TAG)
	docker push 335257738167.dkr.ecr.us-east-2.amazonaws.com/$(ECR_REPO):$(ECR_TAG)
