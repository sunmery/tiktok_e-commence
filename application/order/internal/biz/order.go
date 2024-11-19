package biz

import (
	"context"

	cartV1 "order/api/cart/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type Address struct {
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	ZipCode       int32  `json:"zipCode"`
}
type PlaceOrderReq struct {
	UserId      string       `json:"user_id"`
	UserCurrent string       `json:"user_current"`
	Address     Address      `json:"address"`
	Email       string       `json:"email"`
	OrderItems  []*OrderItem `json:"order_items"`
}
type OrderResult struct {
	OrderId int32 `json:"order_id"`
}

type PlaceOrderResp struct {
	Order OrderResult `json:"order"`
}

type OrderItem struct {
	Item cartV1.CartItem `json:"item"`
	Cost float32         `json:"cost"`
}
type Order struct {
	OrderItems []*OrderItem
}
type ListOrderReq struct {
	UserId string `json:"user_id"`
}
type ListOrderResp struct {
	Orders []*Order `json:"orders"`
}

type MarkOrderPaidReq struct {
	UserId  string `json:"user_id"`
	OrderId string `json:"order_id"`
}

type MarkOrderPaidResp struct{}

// OrderRepo is a Greater repo.
type OrderRepo interface {
	PlaceOrder(ctx context.Context, req *PlaceOrderReq) (*PlaceOrderResp, error)
	ListOrder(ctx context.Context, req *ListOrderReq) (*ListOrderResp, error)
	MarkOrderPaid(ctx context.Context, req *MarkOrderPaidReq) (*MarkOrderPaidResp, error)
}

// OrderUsecase is a Order usecase.
type OrderUsecase struct {
	repo OrderRepo
	log  *log.Helper
}

// NewOrderUsecase new a Order usecase.
func NewOrderUsecase(repo OrderRepo, logger log.Logger) *OrderUsecase {
	return &OrderUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (o *OrderUsecase) PlaceOrder(ctx context.Context, req *PlaceOrderReq) (*PlaceOrderResp, error) {
	o.log.WithContext(ctx).Debugf("PlaceOrderReq: %+v", req)
	return o.repo.PlaceOrder(ctx, req)
}
func (o *OrderUsecase) ListOrder(ctx context.Context, req *ListOrderReq) (*ListOrderResp, error) {
	return o.repo.ListOrder(ctx, req)
}
func (o *OrderUsecase) MarkOrderPaid(ctx context.Context, req *MarkOrderPaidReq) (*MarkOrderPaidResp, error) {
	return o.repo.MarkOrderPaid(ctx, req)
}
