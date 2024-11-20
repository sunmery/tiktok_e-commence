package data

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	productV1 "order/api/product/v1"
	"order/internal/data/models"

	"fmt"
	"github.com/go-kratos/kratos/v2/log"

	"order/internal/biz"
)

type orderRepo struct {
	data *Data
	log  *log.Helper
}

// PlaceOrder TODO db trans
func (o *orderRepo) PlaceOrder(ctx context.Context, req *biz.PlaceOrderReq) (*biz.PlaceOrderResp, error) {
	userAddress, err := o.data.db.CreatUserAddress(ctx, modules.CreatUserAddressParams{
		UserID:        req.UserId,
		StreetAddress: req.Address.StreetAddress,
		City:          req.Address.City,
		State:         req.Address.State,
		Country:       req.Address.Country,
		ZipCode:       req.Address.ZipCode,
	})
	if err != nil {
		// 没有则创建搞地址
		if errors.Is(err, pgx.ErrNoRows) {
			userAddress, err = o.data.db.CreatUserAddress(ctx, modules.CreatUserAddressParams{
				UserID:        req.UserId,
				StreetAddress: req.Address.StreetAddress,
				City:          req.Address.City,
				State:         req.Address.State,
				Country:       req.Address.Country,
				ZipCode:       req.Address.ZipCode,
			})
			if err != nil {
				return nil, err
			}
		}
		return nil, err
	}

	order, orderErr := o.data.db.CreateOrder(ctx, modules.CreateOrderParams{
		Email:        req.Email,
		UserID:       req.UserId,
		UserCurrency: req.UserCurrent,
		AddressID:    userAddress.ID,
	})
	if orderErr != nil {
		return nil, orderErr
	}

	fmt.Printf("orders: '%+v'", order)
	product, err := o.data.productClient.GetProduct(ctx, &productV1.GetProductReq{
		Id: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	fmt.Printf("product: '%+v'", product)

	return &biz.PlaceOrderResp{
		Order: biz.OrderResult{
			OrderId: order.ID,
		},
	}, nil
}

func (o *orderRepo) ListOrder(ctx context.Context, req *biz.ListOrderReq) (*biz.ListOrderResp, error) {
	// TODO implement me
	panic("implement me")
}

func (o *orderRepo) MarkOrderPaid(ctx context.Context, req *biz.MarkOrderPaidReq) (*biz.MarkOrderPaidResp, error) {
	// TODO implement me
	panic("implement me")
}

// NewOrderRepo .
func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
