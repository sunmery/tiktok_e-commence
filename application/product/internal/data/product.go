package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"product/internal/biz"
)

type productRepo struct {
	data *Data
	log  *log.Helper
}

func (p *productRepo) CreateProduct(ctx context.Context, req *biz.CreateProductRequest) (*biz.CreateProductReply, error) {
	product, err := p.data.CreateProduct(ctx, CreateProductParams{
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Categories:  req.Categories,
	})
	if err != nil {
		return nil, err
	}

	return &biz.CreateProductReply{
		Product: biz.Product{
			Id:          uint32(product.ID),
			Name:        product.Name,
			Description: product.Description,
			Picture:     product.Picture,
			Price:       product.Price,
			Categories:  product.Categories,
		},
	}, nil
}

func (p *productRepo) ListProducts(ctx context.Context, req *biz.ListProductsReq) (*biz.ListProductsResp, error) {
	products, err := p.data.ListProducts(ctx, ListProductsParams{
		CategoryName: &req.CategoryName,
		Page:         int64((req.Page - 1) * req.PageSize),
		PageSize:     int64(req.PageSize),
	})
	if err != nil {
		return nil, err
	}

	productsResp := make([]*biz.Product, len(products))
	for i, product := range products {
		productsResp[i] = &biz.Product{
			Id:          uint32(product.ID),
			Name:        product.Name,
			Description: product.Description,
			Picture:     product.Picture,
			Price:       product.Price,
			Categories:  product.Categories,
		}
	}
	return &biz.ListProductsResp{
		Product: productsResp,
	}, nil
}

func (p *productRepo) GetProductReq(ctx context.Context, id uint32) (*biz.GetProductResp, error) {
	// TODO implement me
	panic("implement me")
}

func (p *productRepo) SearchProducts(ctx context.Context, req *biz.SearchProductsReq) (*biz.SearchProductsResp, error) {
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
