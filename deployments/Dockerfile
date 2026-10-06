# ------------------------------
# Build Stage
# ------------------------------
FROM golang:trixie AS builder

WORKDIR /src

# Leverage caching for dependency downloads
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Compile statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /bin/server ./cmd/waqt/main.go

# ------------------------------
# Runtime Stage
# ------------------------------
FROM debian:trixie-slim

RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy binary and runtime data
COPY --from=builder /bin/server /app/server
COPY --from=builder /src/data /app/data

# Run as non-root user
USER nobody:nogroup

EXPOSE 8080

ENTRYPOINT ["/app/server"]
