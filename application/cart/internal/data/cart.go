package data

import (
	"cart/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type cartRepo struct {
	data *Data
	log  *log.Helper
}

func (c cartRepo) AddItem(ctx context.Context, req *biz.AddItemReq) (*biz.AddItemResp, error) {
	// TODO implement me
	panic("implement me")
}

func (c cartRepo) GetCart(ctx context.Context, req *biz.GetCartReq) (*biz.GetCartResp, error) {
	// TODO implement me
	panic("implement me")
}

func (c cartRepo) EmptyCart(ctx context.Context, req *biz.EmptyCartReq) (*biz.EmptyCartResp, error) {
	// TODO implement me
	panic("implement me")
}

// NewCartRepo .
func NewCartRepo(data *Data, logger log.Logger) biz.CartRepo {
	return &cartRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
