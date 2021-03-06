//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/tanmingqian/go-tam/app/apiserver/service/internal/biz"
	"github.com/tanmingqian/go-tam/app/apiserver/service/internal/conf"
	"github.com/tanmingqian/go-tam/app/apiserver/service/internal/data"
	"github.com/tanmingqian/go-tam/app/apiserver/service/internal/server"
	"github.com/tanmingqian/go-tam/app/apiserver/service/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
