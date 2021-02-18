ARG BUILDER_IMAGE="golang:latest"
ARG BASE_IMAGE="alpine"

from ${BUILDER_IMAGE} as builder

run git clone https://github.com/xsteelej/grpc_client_server.git
run cd grpc_client_server/portsService/cmd/portsService && CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/portsService main.go

from ${BASE_IMAGE}

COPY --from=builder /app/portsService /app/portsService

WORKDIR "/app"
ENTRYPOINT ["/app/portsService"]
