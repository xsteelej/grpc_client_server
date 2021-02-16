package main

import (
	"context"
	"github.com/grpc_client_server/clientApi/internal/ports"
	"github.com/grpc_client_server/clientApi/internal/rest"
	dbgrpc "github.com/xsteelej/grpc_client_server/grpc"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

const GrpcPortEnvVar = "GRPC_PORT"
const defaultGrpcPort = ":9090"
const RestPortEnvVar = "SERVER_PORT"
const defaultRestPort = ":8081"

func main() {
	conn := grpcClient()
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	dbClient := dbgrpc.NewPortsDatabaseClient(conn)

	var wg sync.WaitGroup
	startJsonReader(wg, ctx, dbClient)

	svr := &http.Server{
		Addr:    getEnv(RestPortEnvVar, defaultRestPort),
		Handler: rest.NewServer(dbClient),
	}
	startRestServer(svr, &wg)
	startShutdownListener(ctx, cancel, svr)
	wg.Wait()
}

func startJsonReader(wg sync.WaitGroup, ctx context.Context, dbClient dbgrpc.PortsDatabaseClient) {
	wg.Add(1)
	go func() {
		ports.ReadJsonFile(ctx, "", dbClient)
		wg.Done()
	}()
}

func grpcClient() *grpc.ClientConn {
	conn, err := grpc.Dial(getEnv(GrpcPortEnvVar, defaultGrpcPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	return conn
}

func startRestServer(svr *http.Server, wg *sync.WaitGroup) error {
	wg.Add(1)
	go func() {
		log.Println("Starting server")
		err := svr.ListenAndServe()
		if err != nil {
			log.Printf("%s", err.Error())
		}
		wg.Done()
	}()
	return nil
}

func startShutdownListener(ctx context.Context, cancel context.CancelFunc, svr *http.Server) {
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt)
	go func() {
		for range shutdownChan {
			svr.Shutdown(ctx)
			cancel()
		}
	}()
}

func getEnv(envvar string, defaultValue string) string {
	value := os.Getenv(envvar)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
