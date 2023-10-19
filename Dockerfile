# Stage 1: Build the Go application
FROM golang:alpine3.18 AS build-stage

# Set the working directory in the container
WORKDIR /home/project/

# Copy your project files into the container
COPY ./ /home/project/

# Create a directory for the build
RUN mkdir -p /home/build

# Download Go module dependencies
RUN go mod download

# Build the Go application
RUN go build -v -o /home/build/api ./cmd/api

# Stage 2: Create a minimal container for the Go application
FROM gcr.io/distroless/static-debian11

# Copy the compiled binary from the build-stage
COPY --from=build-stage /home/build/api /api

# Copy the templates folder
COPY --from=build-stage /home/project/pkg/templates /pkg/templates

# Copy the .env file
COPY --from=build-stage /home/project/.env /

# Expose the port your application listens on
EXPOSE 1243

# Set the command to run the Go application
CMD ["/api"]
