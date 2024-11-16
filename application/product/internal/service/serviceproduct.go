package service

import (
	"context"
	"product/internal/biz"

	pb "product/api/product/v1"
)

type ProductService struct {
	pb.UnsafeProductCatalogServiceServer
	ps *biz.ProductUsecase
}

func (s *ProductService) SearchProducts(ctx context.Context, req *pb.SearchProductsReq) (*pb.SearchProductsResp, error) {
	products, err := s.ps.SearchProducts(ctx, &biz.SearchProductsReq{Query: req.GetQuery()})
	if err != nil {
		return nil, err
	}
	pbProduct := make([]*pb.Product, len(products.Result))
	for i, product := range products.Result {
		pbProduct[i] = &pb.Product{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Picture:     product.Picture,
			Price:       product.Price,
			Categories:  product.Categories,
		}
	}
	return &pb.SearchProductsResp{
		Results: pbProduct,
	}, nil
}

func (s *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductReply, error) {
	p, err := s.ps.CreateProduct(ctx, &biz.CreateProductRequest{
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Categories:  req.Categories,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateProductReply{
		Product: &pb.Product{
			Id:          p.Product.Id,
			Name:        p.Product.Name,
			Description: p.Product.Description,
			Picture:     p.Product.Picture,
			Price:       p.Product.Price,
			Categories:  p.Product.Categories,
		},
	}, nil
}

func (s *ProductService) ListProducts(ctx context.Context, req *pb.ListProductsReq) (*pb.ListProductsResp, error) {
	list, err := s.ps.ListProducts(ctx, &biz.ListProductsReq{
		Page:         uint(req.Page),
		PageSize:     uint(req.PageSize),
		CategoryName: req.CategoryName,
	})
	if err != nil {
		return nil, err
	}
	pbProduct := make([]*pb.Product, len(list.Product))
	for i, product := range list.Product {
		pbProduct[i] = &pb.Product{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Picture:     product.Picture,
			Price:       product.Price,
			Categories:  product.Categories,
		}
	}
	return &pb.ListProductsResp{
		Products: pbProduct,
	}, nil
}

func (s *ProductService) GetProduct(ctx context.Context, req *pb.GetProductReq) (*pb.GetProductResp, error) {
	product, err := s.ps.GetProduct(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetProductResp{Product: &pb.Product{
		Id:          product.Product.Id,
		Name:        product.Product.Name,
		Description: product.Product.Description,
		Picture:     product.Product.Picture,
		Price:       product.Product.Price,
		Categories:  product.Product.Categories,
	}}, nil
}

func NewServiceProductService(ps *biz.ProductUsecase) *ProductService {
	return &ProductService{ps: ps}
}
