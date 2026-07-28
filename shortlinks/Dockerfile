# Build stage
FROM golang:1.23 as builder

# Build the project
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build

# Run the build
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/go-shortlinks .
COPY . .
RUN chmod +x ./go-shortlinks
# override with variables if needed
ENV CONFIG_PATH=/app/config.yaml

ENTRYPOINT ["./go-shortlinks"]

