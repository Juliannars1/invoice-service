# Use the official Golang image as a base image
FROM golang:1.17

# Set the working directory
WORKDIR /app

# Copy the Go module files and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go binary
RUN go build -o invoice-service cmd/main.go

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["./invoice-service"]
