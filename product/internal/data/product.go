package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"product/internal/biz"
	"product/internal/data/models"
	"product/internal/pkg/token"
)

type productRepo struct {
	data *Data
	log  *log.Helper
}

func (p *productRepo) CreateProduct(ctx context.Context, req *biz.CreateProductRequest) (*biz.CreateProductReply, error) {
	payload, err := token.ExtractPayload(ctx)
	if err != nil {
		return nil, err
	}

	if req.Owner != payload.Owner || req.Username != payload.Name {
		return nil, errors.New("invalid token")
	}

	product, err := p.data.db.CreateProduct(ctx, models.CreateProductParams{
		Owner:       req.Owner,
		Username:    req.Username,
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
	products, err := p.data.db.ListProducts(ctx, models.ListProductsParams{
		Page:     int64((req.Page - 1) * req.PageSize),
		PageSize: int64(req.PageSize),
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

func (p *productRepo) GetProduct(ctx context.Context, id uint32) (*biz.GetProductResp, error) {
	product, err := p.data.db.GetProduct(ctx, int32(id))
	if err != nil {
		return nil, err
	}
	return &biz.GetProductResp{Product: &biz.Product{
		Id:          uint32(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Picture:     product.Picture,
		Price:       product.Price,
		Categories:  product.Categories,
	}}, nil
}

func (p *productRepo) SearchProducts(ctx context.Context, req *biz.SearchProductsReq) (*biz.SearchProductsResp, error) {

	products, err := p.data.db.SearchProducts(ctx, &req.Query)
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
	return &biz.SearchProductsResp{
		Result: productsResp,
	}, nil
}

// NewProductRepo .
func NewProductRepo(data *Data, logger log.Logger) biz.ProductRepo {
	return &productRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
