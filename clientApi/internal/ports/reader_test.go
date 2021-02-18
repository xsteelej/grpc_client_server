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
"AEAJM": {
"name": "Ajman",
"city": "Ajman",
"country": "United Arab Emirates",
"alias": [],
"regions": [],
"coordinates": [
55.5136433,
25.4052165
],
"province": "Ajman",
"timezone": "Asia/Dubai",
"unlocs": [
"AEAJM"
],
"code": "52000"
},
"AEAUH": {
"name": "Abu Dhabi",
"coordinates": [
54.37,
24.47
],
"city": "Abu Dhabi",
"province": "Abu Z¸aby [Abu Dhabi]",
"country": "United Arab Emirates",
"alias": [],
"regions": [],
"timezone": "Asia/Dubai",
"unlocs": [
"AEAUH"
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
}
