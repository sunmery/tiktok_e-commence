package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type BalanceRequest struct {
	Owner   string  `json:"owner"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type BalanceReply struct {
	Message string  `json:"message"`
	Code    uint32  `json:"code"`
	Owner   string  `json:"owner"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

// BalanceRepo is a Greater repo.
type BalanceRepo interface {
	CreateBalance(ctx context.Context, req *BalanceRequest) (*BalanceReply, error)
	UpdateBalance(ctx context.Context, req *BalanceRequest) (*BalanceReply, error)

	GetBalance(ctx context.Context, req *BalanceRequest) (*BalanceReply, error)
}

// BalanceUsecase is a Balance usecase.
type BalanceUsecase struct {
	repo BalanceRepo
	log  *log.Helper
}

// NewBalanceUsecase new a Balance usecase.
func NewBalanceUsecase(repo BalanceRepo, logger log.Logger) *BalanceUsecase {
	return &BalanceUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateBalance creates a Balance, and returns the new Balance.
func (bc *BalanceUsecase) CreateBalance(ctx context.Context, req *BalanceRequest) (*BalanceReply, error) {
	bc.log.WithContext(ctx).Infof("CreateBalance: %v", req)
	return bc.repo.CreateBalance(ctx, req)
}
func (bc *BalanceUsecase) UpdateBalance(ctx context.Context, req *BalanceRequest) (*BalanceReply, error) {
	bc.log.WithContext(ctx).Infof("UpdateBalance: %v", req)
	return bc.repo.UpdateBalance(ctx, req)
}

func (bc *BalanceUsecase) GetBalance(ctx context.Context, req *BalanceRequest) (*BalanceReply, error) {
	bc.log.WithContext(ctx).Infof("GetBalance: %v", req)
	return bc.repo.GetBalance(ctx, req)
}
