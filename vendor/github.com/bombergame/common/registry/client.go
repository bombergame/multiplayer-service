package registry

import (
	"strconv"

	consulapi "github.com/hashicorp/consul/api"
)

//Config contains registry client config parameters
type Config struct {
	RegistryAddress string
	ServiceName     string
	ServiceHost     string
	ServicePort     string
}

//Client is the Consul client wrapper
type Client struct {
	config Config
	client *consulapi.Client
}

//NewClient creates registry client instance
func NewClient(cf Config) *Client {
	return &Client{
		config: cf,
		client: func() *consulapi.Client {
			c, err := consulapi.NewClient(
				func() *consulapi.Config {
					config := consulapi.DefaultConfig()
					config.Address = cf.RegistryAddress
					return config
				}(),
			)
			if err != nil {
				panic(err)
			}
			return c
		}(),
	}
}

//Register performs service registration
func (c *Client) Register() error {
	return c.client.Agent().ServiceRegister(
		&consulapi.AgentServiceRegistration{
			ID:      c.config.ServiceName,
			Name:    c.config.ServiceName,
			Address: c.config.ServiceHost,
			Port: func() int {
				p, err := strconv.Atoi(c.config.ServicePort)
				if err != nil {
					panic(err)
				}
				return p
			}(),
			Check: &consulapi.AgentServiceCheck{},
		},
	)
}

//Deregister deletes service from registry
func (c *Client) Deregister() error {
	return c.client.Agent().ServiceDeregister(c.config.ServiceName)
}
