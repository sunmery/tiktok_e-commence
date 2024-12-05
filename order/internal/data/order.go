package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"net/http"
	v1 "order/api/cart/v1"
	"order/internal/biz"
	"order/internal/data/models"
)

type orderRepo struct {
	data *Data
	log  *log.Helper
}

// PlaceOrder 创建订单
func (o *orderRepo) PlaceOrder(ctx context.Context, req *biz.PlaceOrderReq) (*biz.PlaceOrderResp, error) {
	order, orderErr := o.data.db.CreateOrder(ctx, models.CreateOrderParams{
		Owner:     req.Owner,
		Name:      req.Name,
		Email:     req.Email,
		AddressID: int32(req.AddressId),
		Currency:  req.Currency,
		// CreatedAt: pgtype.Timestamptz{
		// 	Time:  time.Now().UTC(),
		// 	Valid: true,
		// },
	})
	if orderErr != nil {
		return nil, orderErr
	}

	for _, item := range req.OrderItems {
		orderItems, err := o.data.db.CreateOrderItems(ctx, models.CreateOrderItemsParams{
			OrderID:   order,
			ProductID: int32(item.Item.ProductId),
			Quantity:  int32(item.Item.Quantity),
			Cost:      int32(item.Cost),
		})
		if err != nil {
			return nil, err
		}

		fmt.Printf("orderItems: %+v\n", orderItems)

	}
	return &biz.PlaceOrderResp{
		Order: biz.OrderResult{
			OrderId: uint32(order),
		},
	}, nil

}

// ListOrder 查询用户的订单列表
func (o *orderRepo) ListOrder(ctx context.Context, req *biz.ListOrderReq) (*biz.ListOrderResp, error) {
	orders, err := o.data.db.ListOrders(ctx, models.ListOrdersParams{
		Owner: req.Owner,
		Name:  req.Name,
	})
	if err != nil {
		return nil, err
	}

	orderItemsResponse := make([]*biz.OrderItem, len(orders))

	orderItems := make([]models.OrdersOrderItems, len(orders))
	for _, order := range orders {
		var err error
		orderItems, err = o.data.db.ListOrderItems(ctx, order.ID)
		if err != nil {
			return nil, err
		}
	}

	for i, item := range orderItems {
		orderItemsResponse[i] = &biz.OrderItem{
			Item: &v1.CartItem{
				ProductId: uint32(item.ProductID),
				Quantity:  uint32(item.Quantity),
			},
			Cost: float32(item.Cost),
		}
	}

	resp := make([]*biz.Order, len(orders))
	for i, order := range orders {

		resp[i] = &biz.Order{
			OrderItems: orderItemsResponse,
			OrderId:    uint32(order.ID),
			Owner:      order.Owner,
			Name:       order.Name,
			Currency:   order.Currency,
			AddressId:  uint32(order.AddressID),
			Email:      order.Email,
			CreatedAt:  order.CreatedAt,
			UpdatedAt:  order.UpdatedAt,
		}

	}
	return &biz.ListOrderResp{Orders: resp}, err
}

// MarkOrderPaid 标记订单为已支付
func (o *orderRepo) MarkOrderPaid(ctx context.Context, req *biz.MarkOrderPaidReq) (*biz.MarkOrderPaidResp, error) {
	orderPaid, err := o.data.db.MarkOrderPaid(ctx, models.MarkOrderPaidParams{
		ID:    int32(req.OrderId),
		Owner: req.Owner,
		Name:  req.Name,
	})
	if err != nil {
		return nil, err
	}

	return &biz.MarkOrderPaidResp{
		Status:  orderPaid.Status,
		Message: "OK",
		Code:    http.StatusOK,
	}, nil
}

// NewOrderRepo .
func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
