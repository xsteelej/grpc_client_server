package database

import (
	"context"
	"errors"
	grpc "github.com/xsteelej/grpc_client_server/grpc"
)

type memmap struct {
	store map[string]*grpc.Port
}

func NewMemMap() *memmap {
	return &memmap{
		make(map[string]*grpc.Port),
	}
}

func (m *memmap) Write(ctx context.Context, port *grpc.Port) (bool, error) {
	m.store[port.Id] = port
	return true, nil
}

func (m *memmap) Read(ctx context.Context, id string) (*grpc.Port, error) {
	if _, ok := m.store[id]; !ok {
		return nil, errors.New(id + " not found")
	}
	return m.store[id], nil
}
 
