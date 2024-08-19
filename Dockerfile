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

# Create a directory for the database
RUN mkdir /sushi-backend/db
COPY ./db /sushi-backend/db

COPY ./static /sushi-backend/static

COPY ./config /sushi-backend/config

# Copy the built binary from the build stage
COPY --from=builder /sushi-backend/app /sushi-backend/app

# Make the binary executable
RUN chmod +x /sushi-backend/app

# Set the working directory
WORKDIR /sushi-backend


# Run the Golang application
CMD ["./app"]
