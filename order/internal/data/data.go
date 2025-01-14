package data

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"order/internal/conf"
	"order/internal/data/models"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
// var ProviderSet = wire.NewSet(NewData, NewDB, NewCache, NewOrderRepo, NewDiscovery, NewAddressesServiceClient, NewCartServiceClient, NewProductServiceClient)
var ProviderSet = wire.NewSet(NewData, NewDB, NewCache, NewOrderRepo)

// Data .
type Data struct {
	db  *models.Queries
	rdb *redis.Client
	// addressesClient addressesV1.AddressesServiceClient
	// cartClient      cartV1.CartServiceClient
	// productClient   protuctsV1.ProductCatalogServiceClient
}

// NewData .
func NewData(
	logger log.Logger,
	pgx *pgxpool.Pool,
	rdb *redis.Client,
	// addressesClient addressesV1.AddressesServiceClient,
	// productClient protuctsV1.ProductCatalogServiceClient,
	// cartClient cartV1.CartServiceClient,
) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db:  models.New(pgx),
		rdb: rdb,
		// addressesClient: addressesClient,
		// cartClient:      cartClient,
		// productClient:   productClient,
	}, cleanup, nil
}

func NewDB(c *conf.Data) *pgxpool.Pool {
	conn, err := pgxpool.New(context.Background(), c.Database.Source)
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v", err))
	}
	return conn
}

func NewCache(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Protocol: 3,
		Addr:     c.Redis.Addr,
		// Username:     c.Redis.Username,
		// Password:     c.Redis.Password,
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
	})

	return rdb
}

//
// // NewDiscovery 配置服务发现功能
// func NewDiscovery(conf *conf.Registry) (registry.Discovery, error) {
// 	c := consulAPI.DefaultConfig()
// 	c.Address = conf.Consul.Address
// 	c.Scheme = conf.Consul.Scheme
// 	cli, err := consulAPI.NewClient(c)
// 	if err != nil {
// 		return nil, err
// 	}
// 	r := consul.New(cli, consul.WithHealthCheck(false))
// 	return r, nil
// }
//
// // NewAddressesServiceClient 地址微服务
// func NewAddressesServiceClient(d registry.Discovery, logger log.Logger) (addressesV1.AddressesServiceClient, error) {
// 	conn, err := grpc.DialInsecure(
// 		context.Background(),
// 		grpc.WithEndpoint("discovery:///tiktok-e_commence-addresses"),
// 		grpc.WithDiscovery(d),
// 		grpc.WithMiddleware(
// 			recovery.Recovery(),
// 			logging.Client(logger),
// 		),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return addressesV1.NewAddressesServiceClient(conn), nil
// }
//
// // NewCartServiceClient 购物车
// func NewCartServiceClient(d registry.Discovery, logger log.Logger) (cartV1.CartServiceClient, error) {
// 	conn, err := grpc.DialInsecure(
// 		context.Background(),
// 		grpc.WithEndpoint("discovery:///tiktok-e_commence-orders"),
// 		grpc.WithDiscovery(d),
// 		grpc.WithMiddleware(
// 			recovery.Recovery(),
// 			logging.Client(logger),
// 		),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return cartV1.NewCartServiceClient(conn), nil
// }
//
// // NewProductServiceClient 购物车
// func NewProductServiceClient(d registry.Discovery, logger log.Logger) (protuctsV1.ProductCatalogServiceClient, error) {
// 	conn, err := grpc.DialInsecure(
// 		context.Background(),
// 		grpc.WithEndpoint("discovery:///tiktok-e_commence-products"),
// 		grpc.WithDiscovery(d),
// 		grpc.WithMiddleware(
// 			recovery.Recovery(),
// 			logging.Client(logger),
// 		),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return protuctsV1.NewProductCatalogServiceClient(conn), nil
// }
