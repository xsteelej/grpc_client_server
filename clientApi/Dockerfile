ARG BUILDER_IMAGE="golang:latest"
ARG BASE_IMAGE="alpine"

from ${BUILDER_IMAGE} as builder

run git clone https://github.com/xsteelej/grpc_client_server.git
run cd grpc_client_server/clientApi/cmd/clientApi && CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/clientApi main.go

from ${BASE_IMAGE}

COPY --from=builder /app/clientApi /app/clientApi

WORKDIR "/app"
ENTRYPOINT ["/app/clientApi"]
