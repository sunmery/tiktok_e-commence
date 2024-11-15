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

func (s *CartService) AddItem(ctx context.Context, req *pb.AddItemReq) (*pb.AddItemResp, error) {
	_, err := s.cs.CreateCartItem(ctx, &biz.CreateCartItemRequest{
		UserId: req.UserId,
		CartItem: biz.CartItem{
			ProductId: req.Item.ProductId,
			Quantity:  uint32(req.Item.Quantity),
		},
	})
	if err != nil {
		return nil, err
	}
	return &pb.AddItemResp{}, nil
}

func (s *CartService) EmptyCart(ctx context.Context, req *pb.EmptyCartReq) (*pb.EmptyCartResp, error) {
	_, err := s.cs.EmptyCart(ctx, &biz.EmptyCartRequest{UserId: req.UserId})
	if err != nil {
		return nil, err
	}
	return &pb.EmptyCartResp{}, nil
}
func (s *CartService) GetCart(ctx context.Context, req *pb.GetCartReq) (*pb.GetCartResp, error) {
	cart, err := s.cs.GetCart(ctx, &biz.GetCartRequest{UserId: int32(req.UserId)})
	if err != nil {
		return nil, err
	}
	cartItems := make([]*pb.CartItem, len(cart.Cart.Items))
	for i, item := range cart.Cart.Items {
		cartItems[i] = &pb.CartItem{
			ProductId: item.ProductId,
			Quantity:  int32(item.Quantity),
		}
	}
	return &pb.GetCartResp{
		Cart: &pb.Cart{
			UserId: cart.Cart.UserId,
			Items:  cartItems,
		},
	}, nil
}
