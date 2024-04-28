# Use a Go base image that meets the minimum version requirement
FROM golang:1.21.9-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the local code to the container's workspace
COPY . .

# Build the Go app
RUN go build -o myservice

# Use a smaller base image
FROM alpine:latest
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/myservice .

RUN apk --no-cache add curl
# Expose the port the app runs on
EXPOSE 8000

# Run the binary
CMD ["./myservice"]
