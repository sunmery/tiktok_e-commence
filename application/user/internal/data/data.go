package data

import (
	"context"
	"fmt"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"

	"github.com/jackc/pgx/v5/pgxpool"
	"user/internal/conf"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewDB, NewCache, NewCasdoor)

// Data .
type Data struct {
	*Queries
	rdb *redis.Client
	cs  *casdoorsdk.Client
}

// NewData .
func NewData(
	logger log.Logger,
	db *pgxpool.Pool,
	rdb *redis.Client,
	cs *casdoorsdk.Client,
) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		Queries: New(db),
		rdb:     rdb,
		cs:      cs,
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
func NewCasdoor(cc *conf.Casdoor) *casdoorsdk.Client {
	client := casdoorsdk.NewClient(
		cc.Casdoor.Server.Endpoint,
		cc.Casdoor.Server.ClientId,
		cc.Casdoor.Server.ClientSecret,
		cc.Casdoor.Certificate,
		cc.Casdoor.Server.Organization,
		cc.Casdoor.Server.Application,
	)

	return client
}
