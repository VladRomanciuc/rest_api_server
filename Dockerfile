# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Vlad <vlad@vlad.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy Go Modules dependency requirements file
COPY go.mod .

# Copy Go Modules expected hashes file
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy all the app sources (recursively copies files and directories from the host into the image)
COPY . .

# Example to set environment variables
ENV PORT 8000

# Build the app
RUN go build 

# Run the app
CMD ["./api"]
