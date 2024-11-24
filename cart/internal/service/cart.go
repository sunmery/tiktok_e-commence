package service

import (
	"cart/internal/biz"
	"context"
	"net/http"

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
		Owner: req.Owner,
		Name:  req.Name,
		CartItem: biz.CartItem{
			ProductId: req.Item.ProductId,
			Quantity:  uint32(req.Item.Quantity),
		},
	})
	if err != nil {
		return nil, err
	}
	return &pb.AddItemResp{
		Message: "OK",
		Code:    http.StatusOK,
	}, nil
}

func (s *CartService) EmptyCart(ctx context.Context, req *pb.CartReq) (*pb.EmptyCartResp, error) {
	var _, err = s.cs.EmptyCart(ctx, &biz.CartRequest{
		Owner: req.Owner,
		Name:  req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &pb.EmptyCartResp{
		Message: "OK",
		Code:    http.StatusOK,
	}, nil
}
func (s *CartService) GetCart(ctx context.Context, req *pb.CartReq) (*pb.GetCartResp, error) {
	cart, err := s.cs.GetCart(ctx, &biz.CartRequest{
		Owner: req.Owner,
		Name:  req.Name,
	})
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
			Owner: req.Owner,
			Name:  req.Name,
			Items: cartItems,
		},
	}, nil
}
