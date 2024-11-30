package data

import (
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"

	"user/internal/conf"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewCache, NewCasdoor)

// Data .
type Data struct {
	rdb *redis.Client
	cs  *casdoorsdk.Client
}

// NewData .
func NewData(
	logger log.Logger,

	rdb *redis.Client,
	cs *casdoorsdk.Client,
) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{

		rdb: rdb,
		cs:  cs,
	}, cleanup, nil
}

func NewCache(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Protocol:     3,
		Addr:         c.Redis.Addr,
		Username:     c.Redis.Username,
		Password:     c.Redis.Password,
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
