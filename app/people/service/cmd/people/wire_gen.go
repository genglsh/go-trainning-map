// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/genglsh/go-trainning-map/app/people/service/internal/biz"
	"github.com/genglsh/go-trainning-map/app/people/service/internal/conf"
	"github.com/genglsh/go-trainning-map/app/people/service/internal/data"
	"github.com/genglsh/go-trainning-map/app/people/service/internal/server"
	"github.com/genglsh/go-trainning-map/app/people/service/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// Injectors from wire.go:

func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*grpc.Server, func(), error) {
	client := data.NewEntClient(confData, logger)
	dataData, cleanup, err := data.NewData(client, logger)
	if err != nil {
		return nil, nil, err
	}
	peopleRepo := data.NewLabelRepo(dataData, logger)
	peopleUseCase := biz.NewPeopleUseCase(peopleRepo, logger)
	peopleService := service.NewPeopleService(peopleUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, logger, peopleService)
	return grpcServer, func() {
		cleanup()
	}, nil
}
