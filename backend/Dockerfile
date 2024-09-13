# Step 1: Base Image
FROM golang:1.20-alpine as builder

# Step 2: Set working directory inside container
WORKDIR /app

# Step 3: Copy go.mod and go.sum files for dependency management
COPY go.mod go.sum ./

# Step 4: Install dependencies
RUN go mod download

# Step 5: Copy the entire project into the container
COPY . .

# Step 6: Build the Go application
RUN go build -o unified-atlassian-platform .

# Step 7: Final image for running the application
FROM alpine:latest

WORKDIR /app

# Step 8: Copy the built binary from the builder stage
COPY --from=builder /app/unified-atlassian-platform .

# Step 9: Expose the necessary port
EXPOSE 8080

# Step 10: Run the application
CMD ["./unified-atlassian-platform"]
