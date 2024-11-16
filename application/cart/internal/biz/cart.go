package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type CartItem struct {
	ProductId uint32 // 商品id
	Quantity  uint32 // 数量
}

type Cart struct {
	UserId string      `json:"userId"`
	Items  []*CartItem `json:"items"`
}

type CreateCartItemRequest struct {
	UserId   string   `json:"user_id"`
	CartItem CartItem `json:"cart_item"`
}

type CreateCartItemReply struct {
}

type GetCartRequest struct {
	UserId string `json:"user_id"`
}

type GetCartReply struct {
	Cart Cart `json:"cart"`
}

type EmptyCartRequest struct {
	UserId string `json:"user_id"`
}

type EmptyCartReply struct {
}

// CartRepo is a Greater repo.
type CartRepo interface {
	CreateCartItem(ctx context.Context, req *CreateCartItemRequest) (*CreateCartItemReply, error)
	EmptyCart(ctx context.Context, req *EmptyCartRequest) (*EmptyCartReply, error)
	GetCart(ctx context.Context, req *GetCartRequest) (*GetCartReply, error)
}

// CartUsecase is a Cart usecase.
type CartUsecase struct {
	repo CartRepo
	log  *log.Helper
}

// NewCartUsecase new a Cart usecase.
func NewCartUsecase(repo CartRepo, logger log.Logger) *CartUsecase {
	return &CartUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateCartItem 创建购物车商品
func (uc *CartUsecase) CreateCartItem(ctx context.Context, req *CreateCartItemRequest) (*CreateCartItemReply, error) {
	uc.log.WithContext(ctx).Infof("CreateCart: %v", req)
	return uc.repo.CreateCartItem(ctx, req)
}

func (uc *CartUsecase) GetCart(ctx context.Context, req *GetCartRequest) (*GetCartReply, error) {
	// uc.log.WithContext(ctx).Infof("GetCart: %v", req)
	return uc.repo.GetCart(ctx, req)
}
func (uc *CartUsecase) EmptyCart(ctx context.Context, req *EmptyCartRequest) (*EmptyCartReply, error) {
	// uc.log.WithContext(ctx).Infof("EmptyCart: %v", req)
	return uc.repo.EmptyCart(ctx, req)
}
