package data

import (
	"context"

	"addresses/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) CreateAddress(ctx context.Context, req Address) (Reply, error) {
	
}

func (r *greeterRepo) UpdateAddress(ctx context.Context, req Address) (Reply, error) {

}
func (r *greeterRepo) DeleteAddress(ctx context.Context, req Address) (Reply, error) {

}
func (r *greeterRepo) GetAddresses(ctx context.Context, userId string) (Reply, error) {

}
