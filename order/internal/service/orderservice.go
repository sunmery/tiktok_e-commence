package service

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	cartV1 "order/api/cart/v1"
	"order/internal/biz"
	"order/internal/pkg/token"

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
	orderItems := make([]*biz.OrderItem, len(req.OrderItems))
	for i, item := range req.OrderItems {
		orderItems[i] = &biz.OrderItem{
			Item: &cartV1.CartItem{
				ProductId: item.Item.ProductId,
				Quantity:  item.Item.Quantity,
			},
			Cost: item.Cost,
		}
	}

	orderResult, err := s.os.PlaceOrder(ctx, &biz.PlaceOrderReq{
		Owner:      req.Owner,
		Name:       req.Name,
		Currency:   req.Currency,
		AddressId:  req.AddressId,
		Email:      req.Email,
		OrderItems: orderItems,
	})
	if err != nil {
		return nil, err
	}

	return &pb.PlaceOrderResp{
		Order: &pb.OrderResult{OrderId: orderResult.Order.OrderId},
	}, nil
}
func (s *OrderServiceService) ListOrder(ctx context.Context, req *pb.ListOrderReq) (*pb.ListOrderResp, error) {
	result, err := s.os.ListOrder(ctx, &biz.ListOrderReq{
		Owner: req.Owner,
		Name:  req.Name,
	})
	if err != nil {
		return nil, err
	}

	orders := make([]*pb.Order, len(result.Orders))
	for i, order := range result.Orders {
		orders[i] = &pb.Order{
			Owner:      order.Owner,
			Name:       order.Name,
			Email:      order.Email,
			Currency:   order.Currency,
			OrderId:    order.OrderId,
			AddressId:  order.AddressId,
			OrderItems: nil,
			CreatedAt:  timestamppb.New(order.CreatedAt),
			UpdatedAt:  timestamppb.New(order.UpdatedAt),
		}
	}

	return &pb.ListOrderResp{
		Orders: orders,
	}, nil
}
func (s *OrderServiceService) MarkOrderPaid(ctx context.Context, req *pb.MarkOrderPaidReq) (*pb.MarkOrderPaidResp, error) {
	payload, err := token.ExtractPayload(ctx)
	if err != nil {
		return nil, err
	}

	if req.Owner != payload.Owner || req.Name != payload.Name {
		return nil, errors.New("invalid token")
	}

	orderPaid, err := s.os.MarkOrderPaid(ctx, &biz.MarkOrderPaidReq{
		Owner:   payload.Owner,
		Name:    payload.Name,
		OrderId: req.OrderId,
	})
	if err != nil {
		return nil, err
	}
	
	return &pb.MarkOrderPaidResp{
		Status:  orderPaid.Status,
		Code:    orderPaid.Code,
		Message: orderPaid.Message,
	}, nil
}
