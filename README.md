#Building Protobuf
### Generating the ports gRPC service

```protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./ports.proto```

#Running client in Docker
```
docker run -p:8081:8081 --network=grpc_client_server_services -e JSON_FILE=/data/ports.json -e GRPC_SERVER_ADDRESS="ports-service:9090" -v /Users/johste04/code/grpc_client_server/ports.json:/data/ports.json clientapi
```
