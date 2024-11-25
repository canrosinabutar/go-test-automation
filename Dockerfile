# Start with a Golang base image
FROM golang:1.22.1 as builder

ENV JWT_SECRET=mysecretkey

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server/main.go

# Stage 2: Create the final image
FROM alpine:latest

# Install necessary packages
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8989 to the outside world
EXPOSE 8989

# Command to run the executable
CMD ["./main"]


