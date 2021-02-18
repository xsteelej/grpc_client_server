package main

import (
	"context"
	"github.com/grpc_client_server/clientApi/internal/ports"
	"github.com/grpc_client_server/clientApi/internal/rest"
	dbgrpc "github.com/xsteelej/grpc_client_server/grpc"
	"google.golang.org/grpc"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

const grpcServerAddressEnvVar = "GRPC_SERVER_ADDRESS"
const defaultServerAddress = "localhost:9090"
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
	conn, err := grpcClient()
	if err != nil {
		return err
	}
	defer conn.Close()

	dbClient := dbgrpc.NewPortsDatabaseClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	startJsonReader(&wg, ctx, &ports.Sender{dbClient})

	svr := &http.Server{Addr: getEnv(restPortEnvVar, defaultRestPort), Handler: rest.NewServer(dbClient)}
	startRestServer(svr, &wg)
	startShutdownListener(ctx, cancel, svr)
	wg.Wait()

	return nil
}

func startJsonReader(wg *sync.WaitGroup, ctx context.Context, writer io.Writer) error {
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
		ports.ReadJsonFile(ctx, input, writer)
		wg.Done()
	}()

	return nil
}

func grpcClient() (*grpc.ClientConn, error) {
	port := getEnv(grpcServerAddressEnvVar, defaultServerAddress)
	log.Println("Connecting to grpcClient " + port)
	conn, err := grpc.Dial(port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	log.Println("Connected")
	return conn, nil
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
