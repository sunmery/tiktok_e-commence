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
	Owner string      `json:"owner"`
	Name  string      `json:"name"`
	Items []*CartItem `json:"items"`
}

type CreateCartItemRequest struct {
	Owner    string   `json:"owner"`
	Name     string   `json:"name"`
	CartItem CartItem `json:"cart_item"`
}

type CreateCartItemReply struct {
	Message string `json:"message"`
	Code    uint32 `json:"code"`
}

type CartRequest struct {
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

type GetCartReply struct {
	Cart Cart `json:"cart"`
}

type EmptyCartReply struct {
	Message string `json:"message"`
	Code    uint32 `json:"code"`
}

// CartRepo is a Greater repo.
type CartRepo interface {
	CreateCartItem(ctx context.Context, req *CreateCartItemRequest) (*CreateCartItemReply, error)
	EmptyCart(ctx context.Context, req *CartRequest) (*EmptyCartReply, error)
	GetCart(ctx context.Context, req *CartRequest) (*GetCartReply, error)
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

func (uc *CartUsecase) GetCart(ctx context.Context, req *CartRequest) (*GetCartReply, error) {
	uc.log.WithContext(ctx).Infof("GetCart: %v", req)
	return uc.repo.GetCart(ctx, req)
}
func (uc *CartUsecase) EmptyCart(ctx context.Context, req *CartRequest) (*EmptyCartReply, error) {
	uc.log.WithContext(ctx).Infof("EmptyCart: %v", req)
	return uc.repo.EmptyCart(ctx, req)
}
