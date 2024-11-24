package data

import (
	"cart/internal/biz"
	modules "cart/internal/data/models"
	"cart/internal/pkg/token"
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"net/http"
)

type cartRepo struct {
	data *Data
	log  *log.Helper
}

func (c *cartRepo) CreateCartItem(ctx context.Context, req *biz.CreateCartItemRequest) (*biz.CreateCartItemReply, error) {
	payload, err := token.ExtractPayload(ctx)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	if req.Owner != payload.Owner || req.Name != payload.Name {
		return nil, errors.New("invalid token")
	}

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
		Owner:     payload.Owner,
		Name:      payload.Name,
		ProductID: int32(req.CartItem.ProductId),
		Quantity:  quantity,
	})
	if err != nil {
		return nil, err
	}
	log.Infof("result: '%+v'", result)
	return &biz.CreateCartItemReply{
		Message: "OK",
		Code:    http.StatusOK,
	}, nil
}

func (c *cartRepo) GetCart(ctx context.Context, req *biz.CartRequest) (*biz.GetCartReply, error) {
	payload, err := token.ExtractPayload(ctx)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	if req.Owner != payload.Owner || req.Name != payload.Name {
		return nil, errors.New("invalid token")
	}

	cartInfo, err := c.data.db.GetCart(ctx, modules.GetCartParams{
		Owner: &payload.Owner,
		Name:  &payload.Name,
	})
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

			Items: carts,
		},
	}, nil
}

func (c *cartRepo) EmptyCart(ctx context.Context, req *biz.CartRequest) (*biz.EmptyCartReply, error) {
	payload, err := token.ExtractPayload(ctx)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	if req.Owner != payload.Owner || req.Name != payload.Name {
		return nil, errors.New("invalid token")
	}

	_, err = c.data.db.DeleteCart(ctx, modules.DeleteCartParams{
		Owner: payload.Owner,
		Name:  payload.Name,
	})
	if err != nil {
		return nil, err
	}
	return &biz.EmptyCartReply{
		Message: "OK",
		Code:    http.StatusOK,
	}, nil
}

// NewCartRepo .
func NewCartRepo(data *Data, logger log.Logger) biz.CartRepo {
	return &cartRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
