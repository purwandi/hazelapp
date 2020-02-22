package handler

import (
	"github.com/purwandi/hazelapp/gateway/resolver"
	"github.com/purwandi/hazelapp/gateway/schema"
)

type AppContainer struct {
	Schema   string
	Resolver *resolver.Resolver
}

func NewAppContainer() *AppContainer {
	return &AppContainer{
		Schema:   schema.GetRootSchema(),
		Resolver: resolver.NewResolver(),
	}
}
