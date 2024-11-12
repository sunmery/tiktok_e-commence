package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Product struct {
	ID          uint32
	Name        string
	Description string
	Picture     string
	Price       float32
	Categories  []string
}

type ListProductsReq struct {
	Page         int32  `json:"page"`
	PageSize     int64  `json:"pageSize"`
	CategoryName string `json:"categoryName"`
}
type ListProductsResp struct {
	Product []*Product `json:"product"`
}

type SearchProductsReq struct {
	Query string `json:"query"`
}
type SearchProductsResp struct {
	Result Product `json:"result"`
}

// ProductRepo is a Greater repo.
type ProductRepo interface {
	ListProducts(ctx context.Context, req *ListProductsReq) (*ListProductsResp, error)
	GetProductReq(ctx context.Context, id uint32) (*Product, error)
	SearchProducts(ctx context.Context, req *SearchProductsReq) (*SearchProductsResp, error)
}

// ProductUsecase is a Product usecase.
type ProductUsecase struct {
	repo ProductRepo
	log  *log.Helper
}

// NewProductUsecase new a Product usecase.
func NewProductUsecase(repo ProductRepo, logger log.Logger) *ProductUsecase {
	return &ProductUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ProductUsecase) ListProducts(ctx context.Context, req *ListProductsReq) (*ListProductsResp, error) {
	uc.log.WithContext(ctx).Infof("ListProducts: %v", req)
	return uc.repo.ListProducts(ctx, req)
}
