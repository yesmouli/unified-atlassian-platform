# Define variables for the Go binary and Docker image
BINARY=unified-atlassian-platform
DOCKER_IMAGE=unified-atlassian-platform:latest

# Build the Go binary
build:
	go build -o $(BINARY) .

# Run tests
test:
	go test ./...

# Clean up generated files
clean:
	rm -f $(BINARY)

# Build and run the Docker image
docker-build:
	docker build -t $(DOCKER_IMAGE) .

docker-run:
	docker-compose up

docker-clean:
	docker-compose down
