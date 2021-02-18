package service_test

import (
	"context"
	grpc "github.com/xsteelej/grpc_client_server/grpc"
	"github.com/xsteelej/grpc_client_server/portsService/internal/database"
	"github.com/xsteelej/grpc_client_server/portsService/internal/service"
	"testing"
)

func TestReadingAndRetrievingFromServer(t *testing.T) {
	svr := service.NewServiceService(database.NewMemMap())
	ctx := context.Background()
	svr.Write(ctx,&grpc.Port{
		Id:          "1",
		Name:        "Port number #1",
		City:        "City 1",
		Province:    "Province 1",
		Country:     "Country 1",
		Alias:       []string{"1", "2"},
		Regions:     []string{"region 1", "region 2"},
		Coordinates: []float64{55.0, 10.2},
		Timezone:    "time/zone",
		Unlocs:      []string{"unloc1"},
	})

	readPort, err := svr.Read(ctx,&grpc.PortRequest{Id: "1"})
	if err != nil {
		t.Fatal("Error reading port 1")
	}

	if readPort.Name != "Port number #1" {
		t.Fatal("Not the expected name")
	}

	if readPort.City != "City 1" {
		t.Fatal("Not the expected city")
	}

	if readPort.Coordinates[0] != 55.0 {
		t.Fatal("Incorrect coordinate[0]")
	}

	if readPort.Coordinates[1] != 10.2 {
		t.Fatal("Incorrect coordinate[1]")
	}

}
