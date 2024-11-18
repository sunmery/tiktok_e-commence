package service

import (
	"context"
	// cartV1 "order/api/cart/v1"
	"order/internal/biz"

	pb "order/api/order/v1"
)

type OrderServiceService struct {
	pb.UnimplementedOrderServiceServer
	os *biz.OrderUsecase
}

func NewOrderServiceService(os *biz.OrderUsecase) *OrderServiceService {
	return &OrderServiceService{
		os: os,
	}
}

func (s *OrderServiceService) PlaceOrder(ctx context.Context, req *pb.PlaceOrderReq) (*pb.PlaceOrderResp, error) {
	_, err := s.os.PlaceOrder(ctx, &biz.PlaceOrderReq{
		UserId:      req.UserId,
		UserCurrent: req.UserCurrency,
		Address: biz.Address{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       req.Address.ZipCode,
		},
		Email:      req.Email,
		OrderItems: nil,
	})
	if err != nil {
		return nil, err
	}
	return &pb.PlaceOrderResp{
		Order: nil,
	}, nil
}
func (s *OrderServiceService) ListOrder(ctx context.Context, req *pb.ListOrderReq) (*pb.ListOrderResp, error) {
	return &pb.ListOrderResp{}, nil
}
func (s *OrderServiceService) MarkOrderPaid(ctx context.Context, req *pb.MarkOrderPaidReq) (*pb.MarkOrderPaidResp, error) {
	return &pb.MarkOrderPaidResp{}, nil
}
