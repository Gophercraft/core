package core

import "google.golang.org/grpc"

type Manager struct {
	connection grpc.ClientConn
}

type ManagerConfig struct {
	HomeAddress
}

func New(config *ManagerConfig) (manager *Manager) {

}

func (manager *Manager)
