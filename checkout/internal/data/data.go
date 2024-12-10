package data

import (
	"checkout/internal/conf"
	"checkout/internal/data/models"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	creditCardsV1 "checkout/api/credit_cards/v1"
	orderV1 "checkout/api/order/v1"
	paymentV1 "checkout/api/payment/v1"

	"context"
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewCache, NewCheckoutRepo, NewDiscovery, NewOrderServiceClient, NewPaymentServiceClient, NewCreditCardsServiceClient)

// Data .
type Data struct {
	db                *models.Queries
	rdb               *redis.Client
	orderClient       orderV1.OrderServiceClient
	paymentClient     paymentV1.PaymentServiceClient
	creditCardsClient creditCardsV1.CreditCardsServiceClient
}

// NewData .
func NewData(
	logger log.Logger,
	pgx *pgxpool.Pool,
	rdb *redis.Client,
	orderClient orderV1.OrderServiceClient,
	paymentClient paymentV1.PaymentServiceClient,
	creditCardsClient creditCardsV1.CreditCardsServiceClient,
) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db:                models.New(pgx),
		rdb:               rdb,
		orderClient:       orderClient,
		paymentClient:     paymentClient,
		creditCardsClient: creditCardsClient,
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
		Username: c.Redis.Username,
		Password: c.Redis.Password,
		// DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
	})

	return rdb
}

// NewDiscovery 配置服务发现功能
func NewDiscovery(conf *conf.Registry) (registry.Discovery, error) {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		return nil, err
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r, nil
}

// NewOrderServiceClient 订单微服务
func NewOrderServiceClient(d registry.Discovery, logger log.Logger) (orderV1.OrderServiceClient, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///tiktok-e_commence-orders"),
		grpc.WithDiscovery(d),
		grpc.WithMiddleware(
			recovery.Recovery(),
			logging.Client(logger),
		),
	)
	if err != nil {
		return nil, err
	}
	return orderV1.NewOrderServiceClient(conn), nil
}

// NewPaymentServiceClient 支付微服务
func NewPaymentServiceClient(d registry.Discovery, logger log.Logger) (paymentV1.PaymentServiceClient, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///tiktok-e_commence-payments"),
		grpc.WithDiscovery(d),
		grpc.WithMiddleware(
			recovery.Recovery(),
			logging.Client(logger),
		),
	)
	if err != nil {
		return nil, err
	}
	return paymentV1.NewPaymentServiceClient(conn), nil
}

// NewCreditCardsServiceClient 地址微服务
func NewCreditCardsServiceClient(d registry.Discovery, logger log.Logger) (creditCardsV1.CreditCardsServiceClient, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///tiktok-e_commence-credit_cards"),
		grpc.WithDiscovery(d),
		grpc.WithMiddleware(
			recovery.Recovery(),
			logging.Client(logger),
		),
	)
	if err != nil {
		return nil, err
	}
	return creditCardsV1.NewCreditCardsServiceClient(conn), nil
}
