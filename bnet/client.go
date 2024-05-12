package bnet

import (
	"context"
	"fmt"
	"strings"

	"github.com/Gophercraft/core/bnet/rpc"
	"github.com/Gophercraft/core/bnet/util"
)

type Client struct {
	dial_options []DialOption
	connection   *Connection
	// registered services
	services map[util.ServiceHash]*registered_service
}

type ClientOption func(*Client)

func NewClient(dial_options ...DialOption) (client *Client) {
	client = new(Client)
	client.dial_options = dial_options
	client.services = make(map[util.ServiceHash]*registered_service)
	return
}

func (client *Client) RegisterService(service_desc *rpc.ServiceDesc, service any) (err error) {
	client.services[service_desc.ServiceHash] = &registered_service{
		descriptor: service_desc,
		service:    service,
	}
	return
}

func (client *Client) get_registered_service(hash util.ServiceHash) (service *registered_service) {
	service = client.services[hash]
	return
}

func (client *Client) Connect(portal string) (err error) {
	address := portal
	if !strings.Contains(address, ":") {
		address += ":1119"
	}
	fmt.Println("Dialing", address)
	client.connection, err = DialContext(context.WithValue(context.TODO(), client_context_key, client), address, client.dial_options...)
	if err != nil {
		return
	}

	return
}

func (client *Client) Disconnect() (err error) {
	return client.connection.Close()
}

func (client *Client) Connection() *Connection {
	return client.connection
}
