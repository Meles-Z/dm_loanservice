# syntax=docker/dockerfile:1.2
FROM golang:1.25 AS builder
WORKDIR /go/src/dm_loanservice
COPY . .

# Use SSH agent forwarding to allow private repo access
RUN --mount=type=ssh go env -w GOPRIVATE=github.com/brianjobling
RUN --mount=type=ssh git config --global --add url."git@github.com:".insteadOf "https://github.com/"

# Ensure the .ssh directory exists and add GitHub to known_hosts
RUN mkdir -p /root/.ssh && chmod 700 /root/.ssh && \
    ssh-keyscan -H github.com >> /root/.ssh/known_hosts

# Configure Git to use SSH without strict host key checking and test connection
RUN --mount=type=ssh \
    git config --global core.sshCommand "ssh -o StrictHostKeyChecking=accept-new" && \
    ssh -T git@github.com || true

# Download dependencies and update go.sum
RUN --mount=type=ssh go mod download
RUN --mount=type=ssh go mod tidy

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o product-service cmd/main.go


FROM alpine:latest
ARG APP_USER=dm

# Create a new user and group
RUN addgroup -S ${APP_USER} && adduser -S ${APP_USER} -G ${APP_USER}

# Install compatibility libraries for any C dependencies
RUN apk add --no-cache libc6-compat ca-certificates

WORKDIR /app/

# Copy the compiled Go binary and necessary files from the builder stage
COPY --from=builder /go/src/dm_loanservice/product-service ./
COPY migrations ./migrations
COPY config-kube.yaml ./config.yaml

# Switch to the new user
USER ${APP_USER}:${APP_USER}

# Run the service
ENTRYPOINT ["./product-service"]
