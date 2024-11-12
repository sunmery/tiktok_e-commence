package data

import (
	"context"

	"product/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type productRepo struct {
	data *Data
	log  *log.Helper
}

func (p productRepo) ListProducts(ctx context.Context, req *biz.ListProductsReq) (*biz.ListProductsResp, error) {
	// TODO implement me
	panic("implement me")
}

func (p productRepo) GetProductReq(ctx context.Context, id uint32) (*biz.Product, error) {
	// TODO implement me
	panic("implement me")
}

func (p productRepo) SearchProducts(ctx context.Context, req *biz.SearchProductsReq) (*biz.SearchProductsResp, error) {
	// TODO implement me
	panic("implement me")
}

// NewProductRepo .
func NewProductRepo(data *Data, logger log.Logger) biz.ProductRepo {
	return &productRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
