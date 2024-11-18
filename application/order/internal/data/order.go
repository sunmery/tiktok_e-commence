package data

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"order/internal/data/models"

	"fmt"
	"github.com/go-kratos/kratos/v2/log"

	"order/internal/biz"
)

type orderRepo struct {
	data *Data
	log  *log.Helper
}

// PlaceOrder TODO
func (o *orderRepo) PlaceOrder(ctx context.Context, req *biz.PlaceOrderReq) (*biz.PlaceOrderResp, error) {
	var userAddress modules.AddressesAddresses
	var err error
	userAddress, err = o.data.db.CreatUserAddress(ctx, modules.CreatUserAddressParams{
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
	fmt.Printf("userAddress: %+v", userAddress)

	order, err := o.data.db.CreateOrder(ctx, modules.CreateOrderParams{
		Email:        &req.Email,
		UserID:       &req.UserId,
		UserCurrency: &req.UserCurrent,
		AddressID:    &userAddress.ID,
	})
	fmt.Printf("orders: '%+v'", order)
	if err != nil {
		return nil, err
	}
	return nil, err
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
