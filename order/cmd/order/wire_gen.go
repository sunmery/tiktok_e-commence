// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"order/internal/biz"
	"order/internal/conf"
	"order/internal/data"
	"order/internal/server"
	"order/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, registry *conf.Registry, logger log.Logger) (*kratos.App, func(), error) {
	registrar := server.NewRegistrar(registry)
	pool := data.NewDB(confData)
	client := data.NewCache(confData)
	discovery, err := data.NewDiscovery(registry)
	if err != nil {
		return nil, nil, err
	}
	productCatalogServiceClient, err := data.NewProductServiceClient(discovery, logger)
	if err != nil {
		return nil, nil, err
	}
	cartServiceClient, err := data.NewCartServiceClient(discovery, logger)
	if err != nil {
		return nil, nil, err
	}
	dataData, cleanup, err := data.NewData(logger, pool, client, productCatalogServiceClient, cartServiceClient)
	if err != nil {
		return nil, nil, err
	}
	orderRepo := data.NewOrderRepo(dataData, logger)
	orderUsecase := biz.NewOrderUsecase(orderRepo, logger)
	orderServiceService := service.NewOrderServiceService(orderUsecase)
	grpcServer := server.NewGRPCServer(confServer, orderServiceService, logger)
	httpServer := server.NewHTTPServer(confServer, orderServiceService, logger)
	app := newApp(logger, registrar, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
