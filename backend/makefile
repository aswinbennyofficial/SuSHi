.PHONY: all dev prod kill

all: dev prod kill

dev:
	@echo "Building development Docker image..."
	go build -o app . &&  docker build -f Dockerfile.dev -t breeze5690/sushi-backend:v1 && podman-compose -f docker-compose-dev.yaml up


prod:
	@echo "Building production Docker image..."
	podman build -t breeze5690/sushi-backend-prod:v1 . && podman-compose up

kill:
	podman-compose -f ./docker-compose-dev.yaml  down