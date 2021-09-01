package utils

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/pborman/uuid"
	"log"
)

var consulClient *api.Client
var ServiceID string
var ServicePort int

func init() {
	ServiceID = "userService_" + uuid.New()[:8]
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	var err error
	consulClient, err = api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
}

func RegistryService(name string, port int) error {
	ServicePort = port
	service := &api.AgentServiceRegistration{
		ID:      ServiceID,
		Name:    name,
		Tags:    []string{"primary"},
		Port:    port,
		Address: "127.0.0.1",
		Check: &api.AgentServiceCheck{
			Interval: "5s",
			HTTP:     fmt.Sprintf("http://127.0.0.1:%d/health", port),
		},
	}
	return consulClient.Agent().ServiceRegister(service)
}

func DeregisterService() error {
	return consulClient.Agent().ServiceDeregister(ServiceID)
}
