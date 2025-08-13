# --- Stage 1: Build ---
# Use an official Go image as a builder.
FROM golang:1.24.5-alpine AS builder

# Set the working directory inside the container.
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code.
COPY . .

# Build the application. CGO_ENABLED=0 is important for a static binary.
# -o /app/chat-app specifies the output path for the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/chat-app .

# --- Stage 2: Final Image ---
# Use a minimal, non-root base image for security.
FROM alpine:3.19

# Set the working directory.
WORKDIR /app

# Copy only the compiled binary from the builder stage.
COPY --from=builder /app/chat-app .

# Expose the port the app runs on.
EXPOSE 8080

# The command to run when the container starts.
CMD ["./chat-app"]