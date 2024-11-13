package service

import (
	"cart/internal/biz"
	"context"

	pb "cart/api/cart/v1"
)

type CartService struct {
	pb.UnimplementedCartServiceServer
	cs *biz.CartUsecase
}

func NewCartService(cs *biz.CartUsecase) *CartService {
	return &CartService{cs: cs}
}

func (s *CartService) CreateCart(ctx context.Context, req *pb.AddItemReq) (*pb.AddItemResp, error) {
	item, err := s.cs.CreateCart(ctx, &biz.Cart{
		Hello: "",
	})
	return &pb.AddItemResp{}, nil
}

func (s *CartService) DeleteCart(ctx context.Context, req *pb.DeleteCartRequest) (*pb.DeleteCartReply, error) {
	return &pb.DeleteCartReply{}, nil
}
func (s *CartService) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartReply, error) {
	return &pb.GetCartReply{}, nil
}
