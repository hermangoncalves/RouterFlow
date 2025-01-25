package api

import "github.com/go-routeros/routeros/v3"

type RouterOsClient interface {
	RunCommand(command string, args ...string) (*routeros.Reply, error)
	Close()
}

type routerOsClient struct {
	Client *routeros.Client
}

func NewRouterOsClient(address, username, password string) (RouterOsClient, error) {
	client, err := routeros.Dial(address, username, password)
	if err != nil {
		return nil, err
	}
	return &routerOsClient{Client: client}, nil
}

func (c *routerOsClient) Close() {
	if c.Client != nil {
		c.Client.Close()
	}
}

func (c *routerOsClient) RunCommand(command string, args ...string) (*routeros.Reply, error) {
	reply, err := c.Client.RunArgs(append([]string{command}, args...))
	if err != nil {
		return nil, err
	}
	return reply, nil
}
