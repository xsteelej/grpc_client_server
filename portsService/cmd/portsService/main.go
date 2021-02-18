package main

import (
	"github.com/xsteelej/grpc_client_server/grpc"
	"github.com/xsteelej/grpc_client_server/portsService/internal/database"
	"github.com/xsteelej/grpc_client_server/portsService/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
)

const grpcPortEnvVar = "GRPC_PORT"
const defaultGrpcPort = ":9090"

func main() {
	var wg sync.WaitGroup
	err := startGrpcServer(&wg, getEnv(grpcPortEnvVar,defaultGrpcPort))
	if err != nil {
		log.Printf("Error: %s", err.Error())
		os.Exit(1)
	}
	wg.Wait()
}

func startGrpcServer(wg *sync.WaitGroup, port string) error {
	wg.Add(1)
	go func() {
		svr := grpc.NewServer()
		portsDB.RegisterPortsDatabaseServer(svr,service.NewServiceService(database.NewMemMap()))
		lis, _ := net.Listen("tcp",port)
		log.Printf("Ports Service gRPC listening on port: %s\n",port)
		startShutdownListener(wg, svr)
		err := svr.Serve(lis)
		if err != nil {
			log.Printf("%s\n",err.Error())
		}
	}()
	return nil
}

func startShutdownListener(wg *sync.WaitGroup, server *grpc.Server) {
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt)
	go func() {
		for range shutdownChan {
			server.GracefulStop()
			log.Println("Server shutdown")
			wg.Done()
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