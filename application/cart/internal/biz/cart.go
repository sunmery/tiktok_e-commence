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
	UserId uint32
	Items  []*CartItem
}

type AddItemReq struct {
	UserId   uint32   `json:"user_id"`
	CartItem CartItem `json:"cart_item"`
}

type AddItemResp struct {
	UserId   uint32   `json:"user_id"`
	CartItem CartItem `json:"cart_item"`
}

type GetCartReq struct {
	UserId uint32 `json:"user_id"`
}

type GetCartResp struct {
	Cart Cart `json:"cart"`
}

type EmptyCartReq struct {
	UserId uint32 `json:"user_id"`
}

type EmptyCartResp struct {
}

// CartRepo is a Greater repo.
type CartRepo interface {
	AddItem(ctx context.Context, req *AddItemReq) (*AddItemResp, error)
	GetCart(ctx context.Context, req *GetCartReq) (*GetCartResp, error)
	EmptyCart(ctx context.Context, req *EmptyCartReq) (*EmptyCartResp, error)
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

// AddItem 添加购物车
func (uc *CartUsecase) AddItem(ctx context.Context, req *AddItemReq) (*AddItemResp, error) {
	uc.log.WithContext(ctx).Infof("AddItem: %v", req)
	return uc.repo.AddItem(ctx, req)
}
