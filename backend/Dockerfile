# Build stage
FROM golang:latest AS builder

# Copy the application code
COPY . /sushi-backend

# Set the working directory
WORKDIR /sushi-backend

# Build the Golang application
RUN go build -o app .




# Production stage
FROM debian

RUN apt update && apt install -y ca-certificates

# Create a directory for the application
RUN mkdir /sushi-backend

# Create a directory for the log file
RUN mkdir /sushi-backend/logs

# Copy the built binary from the build stage
COPY --from=builder /sushi-backend/app /sushi-backend/app




# Set the working directory
WORKDIR /sushi-backend


# Run the Golang application
CMD ["./app"]
