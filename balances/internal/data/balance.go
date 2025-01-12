package data

import (
	"balances/internal/biz"
	"balances/internal/data/models"
	"context"
	"google.golang.org/grpc/codes"

	"github.com/go-kratos/kratos/v2/log"
)

type balanceRepo struct {
	data *Data
	log  *log.Helper
}

func (b *balanceRepo) CreateBalance(ctx context.Context, req *biz.BalanceRequest) (*biz.BalanceReply, error) {
	balance, err := b.data.db.CreatBalance(ctx, models.CreatBalanceParams{
		Owner:   req.Owner,
		Name:    req.Name,
		Balance: req.Balance,
	})
	if err != nil {
		return nil, err
	}
	return &biz.BalanceReply{
		Message: "OK",
		Code:    uint32(codes.OK),
		Owner:   balance.Owner,
		Name:    balance.Name,
		Balance: balance.Balance,
	}, nil
}

func (b *balanceRepo) UpdateBalance(ctx context.Context, req *biz.BalanceRequest) (*biz.BalanceReply, error) {
	balance, err := b.data.db.UpdateBalance(ctx, models.UpdateBalanceParams{
		Owner:   req.Owner,
		Name:    req.Name,
		Balance: &req.Balance,
	})
	if err != nil {
		return nil, err
	}
	return &biz.BalanceReply{
		Message: "OK",
		Code:    uint32(codes.OK),
		Owner:   balance.Owner,
		Name:    balance.Name,
		Balance: balance.Balance,
	}, nil
}

func (b *balanceRepo) GetBalance(ctx context.Context, req *biz.BalanceRequest) (*biz.BalanceReply, error) {
	balance, err := b.data.db.GetBalance(ctx, models.GetBalanceParams{
		Owner: req.Owner,
		Name:  req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &biz.BalanceReply{
		Message: "OK",
		Code:    uint32(codes.OK),
		Owner:   balance.Owner,
		Name:    balance.Name,
		Balance: balance.Balance,
	}, nil
}

// NewBalanceRepo .
func NewBalanceRepo(data *Data, logger log.Logger) biz.BalanceRepo {
	return &balanceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
