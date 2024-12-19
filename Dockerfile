FROM golang:1.23.2-bullseye AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy
COPY . .
RUN go build -o build/main main.go
RUN apt-get update && apt-get install -y gcc
RUN mkdir -p build/plugins
ENV CGO_ENABLED 1
RUN make build-plugins


FROM debian:bullseye-slim
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
ENV CGO_ENABLED 1
WORKDIR /app
COPY --from=builder /app/build /app
# COPY config.yaml /app/config.yaml # Not recommended uncomment this to add your config to the image
EXPOSE 8080
CMD ["/app/main"]