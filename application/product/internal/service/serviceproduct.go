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

func (s *ProductService) ListProducts(ctx context.Context, req *pb.ListProductsReq) (*pb.ListProductsResp, error) {
	list, err := s.ps.ListProducts(ctx, &biz.ListProductsReq{
		Page:         req.Page,
		PageSize:     req.PageSize,
		CategoryName: req.CategoryName,
	})
	if err != nil {
		return nil, err
	}
	pbProduct := make([]*pb.Product, len(list.Product))
	for i, product := range list.Product {
		pbProduct[i] = &pb.Product{
			Id:          product.ID,
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
	// TODO implement me
	panic("implement me")
}

func (s *ProductService) SearchProducts(ctx context.Context, req *pb.SearchProductsReq) (*pb.SearchProductsResp, error) {
	// TODO implement me
	panic("implement me")
}

func NewServiceProductService(ps *biz.ProductUsecase) *ProductService {
	return &ProductService{ps: ps}
}
