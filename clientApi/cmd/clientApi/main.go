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

const grpcPortEnvVar = "GRPC_PORT"
const defaultGrpcPort = ":9090"
const restPortEnvVar = "SERVER_PORT"
const defaultRestPort = ":8081"
const jsonFileLocationEnvVar = "JSON_FILE"

func main() {
	err := run()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}

func run() error {
	conn := grpcClient()
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	dbClient := dbgrpc.NewPortsDatabaseClient(conn)

	var wg sync.WaitGroup
	startJsonReader(&wg, ctx, dbClient)

	svr := &http.Server{
		Addr:    getEnv(restPortEnvVar, defaultRestPort),
		Handler: rest.NewServer(dbClient),
	}
	startRestServer(svr, &wg)
	startShutdownListener(ctx, cancel, svr)
	wg.Wait()

	return nil
}

func startJsonReader(wg *sync.WaitGroup, ctx context.Context, dbClient dbgrpc.PortsDatabaseClient) error {
	filename := getEnv(jsonFileLocationEnvVar, "")
	if len(filename) == 0 {
		return nil
	}
	input, err := os.Open(filename)
	if err != nil {
		return err
	}

	wg.Add(1)
	go func() {
		ports.ReadJsonFile(ctx, input, &ports.Sender{dbClient})
		wg.Done()
	}()

	return nil
}

func grpcClient() *grpc.ClientConn {
	conn, err := grpc.Dial(getEnv(grpcPortEnvVar, defaultGrpcPort), grpc.WithInsecure())
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
