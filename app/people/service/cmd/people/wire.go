// +build wireinject

package main

import (
	"github.com/genglsh/go-trainning-map/app/people/service/internal/biz"
	"github.com/genglsh/go-trainning-map/app/people/service/internal/conf"
	"github.com/genglsh/go-trainning-map/app/people/service/internal/data"
	"github.com/genglsh/go-trainning-map/app/people/service/internal/server"
	"github.com/genglsh/go-trainning-map/app/people/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
