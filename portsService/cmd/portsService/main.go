package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	portsDB "portsService/grpc"
	"portsService/internal/service"
	"sync"
)

const PortNumberEnvVar = "PORTS_SERVICE_ADDRESS"

func main() {
	var wg sync.WaitGroup
	err := startGrpcServer(&wg, getEnv(PortNumberEnvVar,":9090"))
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
		s := service.Server{}
		portsDB.RegisterPortsDatabaseServer(svr,&s)
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