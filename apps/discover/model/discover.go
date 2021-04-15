package model

import (
	"oi.io/apps/discover/configs"
)

type Discovery struct {
	config    *configs.GlobalConfig
	protected bool
	Registry  *Registry
}
func NewDiscovery(config *configs.GlobalConfig) *Discovery {
	dis := &Discovery{
		protected: false,
		config:    config,
		Registry:  NewRegistry(), //init registry
	}
	return dis
}
//init discovery
//var Discovery *model.Discovery