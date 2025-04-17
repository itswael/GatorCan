-include .env
# Define variables
IMAGE_NAME = gatorcan_backend
IMAGE_TAG = main

# Build the Docker image
build:
	docker build -t muthu1729/muthu:$(IMAGE_NAME)_$(IMAGE_TAG) -f Dockerfile.$(IMAGE_TAG) .
		
# Run container
run:
	docker run -t -p 8080:8080 --name $(IMAGE_NAME)_$(IMAGE_TAG)  muthu1729/muthu:$(IMAGE_NAME)_$(IMAGE_TAG)

save:
	docker save -o $(IMAGE_TAG).tar $(IMAGE_NAME):$(IMAGE_TAG)

# Push the Docker image to a ECR
push:
	docker push muthu1729/muthu:$(IMAGE_NAME)_$(IMAGE_TAG)

# Remove the Docker image
clean:
	docker rm $(IMAGE_NAME)_$(IMAGE_TAG)
	docker rmi muthu1729/muthu:$(IMAGE_NAME)_$(IMAGE_TAG)

pull:
	docker pull muthu1729/muthu:$(IMAGE_NAME)_$(IMAGE_TAG)

test:
	docker run -t -p 8080:8080 --name $(IMAGE_NAME)_$(IMAGE_TAG) muthu1729/muthu:$(IMAGE_NAME)_$(IMAGE_TAG)