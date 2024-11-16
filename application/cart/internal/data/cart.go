package data

import (
	"cart/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type cartRepo struct {
	data *Data
	log  *log.Helper
}

func (c *cartRepo) CreateCartItem(ctx context.Context, req *biz.CreateCartItemRequest) (*biz.CreateCartItemReply, error) {
	err := c.data.CreateCartItem(ctx, CreateCartItemParams{
		UserID:    req.UserId,
		ProductID: int32(req.CartItem.ProductId),
		Quantity:  int32(req.CartItem.Quantity),
	})
	if err != nil {
		return nil, err
	}
	return &biz.CreateCartItemReply{}, nil
}

func (c *cartRepo) GetCart(ctx context.Context, req *biz.GetCartRequest) (*biz.GetCartReply, error) {
	cartInfo, err := c.data.GetCart(ctx, &req.UserId)
	if err != nil {
		return nil, err
	}
	log.Infof("cartInfo%+v", cartInfo)
	carts := make([]*biz.CartItem, len(cartInfo))
	for i, row := range cartInfo {
		carts[i] = &biz.CartItem{
			ProductId: uint32(row.ProductID),
			Quantity:  uint32(row.Quantity),
		}
	}
	return &biz.GetCartReply{
		Cart: biz.Cart{
			UserId: "",
			Items:  carts,
		},
	}, nil
}

func (c *cartRepo) EmptyCart(ctx context.Context, req *biz.EmptyCartRequest) (*biz.EmptyCartReply, error) {
	_, err := c.data.DeleteCart(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &biz.EmptyCartReply{}, nil
}

// NewCartRepo .
func NewCartRepo(data *Data, logger log.Logger) biz.CartRepo {
	return &cartRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
