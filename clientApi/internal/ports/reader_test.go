package ports_test

import (
	"context"
	"errors"
	"github.com/grpc_client_server/clientApi/internal/ports"
	dbgrpc "github.com/xsteelej/grpc_client_server/grpc"
	"google.golang.org/grpc"
	"strings"
	"testing"
)

const testJson = `{
"1": {
"name": "Test Port 1",
"city": "Test City 1",
"country": "Test Country 1",
"alias": [],
"regions": [],
"coordinates": [
55.5136433,
25.4052165
],
"province": "Test Province",
"timezone": "Test/zone",
"unlocs": [
"1"
],
"code": "52000"
},
"2": {
"name": "Test Name 2",
"coordinates": [
54.37,
24.47
],
"city": "Test City 2",
"province": "Test Province 2",
"country": "Test Country 2",
"alias": [],
"regions": [],
"timezone": "Test/Zone2",
"unlocs": [
"2"
],
"code": "52001"
}}`

type MockJsonDbClient struct {
	Ports map[string]*dbgrpc.Port
}

func NewMockJsonClient() *MockJsonDbClient {
	return &MockJsonDbClient{Ports: make(map[string]*dbgrpc.Port)}
}

func (m *MockJsonDbClient) Write(ctx context.Context, in *dbgrpc.Port, opts ...grpc.CallOption) (*dbgrpc.WriteResponse, error) {
	m.Ports[in.Id] = in
	return &dbgrpc.WriteResponse{
		Id:      in.Id,
		Success: true,
	}, nil
}

func (m *MockJsonDbClient) Read(ctx context.Context, in *dbgrpc.PortRequest, opts ...grpc.CallOption) (*dbgrpc.Port, error) {
	port, found := m.Ports[in.Id]
	if !found {
		return nil, errors.New("Port not found")
	}
	return port, nil
}

func TestReadJsonFile(t *testing.T) {
	r := strings.NewReader(testJson)
	dbc := NewMockJsonClient()
	ctx, _ := context.WithCancel(context.Background())
	err := ports.ReadJsonFile(ctx, r, &ports.Sender{dbc})
	if err != nil {
		t.Fatal("Error reading json file " + err.Error())
	}
	if len(dbc.Ports) != 2 {
		t.Fatal("Error unexpected number of ports")
	}

	if _, found := dbc.Ports["1"]; !found {
		t.Fatal("Could not find port ID " + "1")
	}

	if _, found := dbc.Ports["2"]; !found {
		t.Fatal("Could not find port ID " + "2")
	}

	aeauhPort := dbc.Ports["2"]

	if aeauhPort.City != "Test City 2" {
		t.Fatal("City not Test City 2")
	}

	if aeauhPort.Coordinates[0] != 54.37 {
		t.Fatal("Coordinate incorrect")
	}

	if aeauhPort.Coordinates[1] != 24.47 {
		t.Fatal("Coordinate incorrect")
	}
}
