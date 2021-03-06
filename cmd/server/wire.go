//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/bagechashu/kratos-layout/app/greeter"
	"github.com/bagechashu/kratos-layout/conf"
	"github.com/bagechashu/kratos-layout/server"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, greeter.ProviderSet, newApp))
}
