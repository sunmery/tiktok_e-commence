package data

import (
	"cart/internal/biz"
	modules "cart/internal/data/models"
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

type cartRepo struct {
	data *Data
	log  *log.Helper
}

func (c *cartRepo) CreateCartItem(ctx context.Context, req *biz.CreateCartItemRequest) (*biz.CreateCartItemReply, error) {
	quantity := int32(req.CartItem.Quantity)
	// 最大上限
	if quantity > 9999 {
		return &biz.CreateCartItemReply{}, errors.New("quantity too big")
	}
	// 最小上限
	if quantity < 0 {
		return nil, errors.New("quantity must be greater than zero")
	}
	result, err := c.data.db.CreateOrUpdateCartItem(ctx, modules.CreateOrUpdateCartItemParams{
		UserID:    req.UserId,
		ProductID: int32(req.CartItem.ProductId),
		Quantity:  quantity,
	})
	if err != nil {
		return nil, err
	}
	log.Infof("result: '%+v'", result)
	return &biz.CreateCartItemReply{}, nil
}

func (c *cartRepo) GetCart(ctx context.Context, req *biz.GetCartRequest) (*biz.GetCartReply, error) {
	cartInfo, err := c.data.db.GetCart(ctx, &req.UserId)
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
	_, err := c.data.db.DeleteCart(ctx, req.UserId)
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
