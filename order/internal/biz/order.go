package biz

import (
	"context"
	"time"

	cartV1 "order/api/cart/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type PlaceOrderReq struct {
	Owner      string       `json:"owner"`
	Name       string       `json:"name"`
	Currency   string       `json:"Currency"`
	AddressId  uint32       `json:"address_id"`
	Email      string       `json:"email"`
	OrderItems []*OrderItem `json:"order_items"`
}
type OrderResult struct {
	OrderId uint32 `json:"order_id"`
}

type PlaceOrderResp struct {
	Order OrderResult `json:"order"`
}

type OrderItem struct {
	Item *cartV1.CartItem `json:"item"`
	Cost float32          `json:"cost"`
}
type Order struct {
	OrderItems []*OrderItem `json:"order_items"`
	OrderId    uint32       `json:"order_id"`
	Owner      string       `json:"owner"`
	Name       string       `json:"name"`
	Currency   string       `json:"currency"`
	AddressId  uint32       `json:"address_id"`
	Email      string       `json:"email"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
}
type ListOrderReq struct {
	Owner string `json:"owner"`
	Name  string `json:"name"`
}
type ListOrderResp struct {
	Orders []*Order `json:"orders"`
}

type MarkOrderPaidReq struct {
	Owner   string `json:"owner"`
	Name    string `json:"name"`
	OrderId uint32 `json:"order_id"`
}

type MarkOrderPaidResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    uint32 `json:"code"`
}

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
