package service

import (
	"balances/internal/biz"
	"context"

	pb "balances/api/balance/v1"
)

type BalanceService struct {
	pb.UnimplementedBalanceServer
	bc *biz.BalanceUsecase
}

func NewBalanceService(bc *biz.BalanceUsecase) *BalanceService {
	return &BalanceService{bc: bc}
}

func (s *BalanceService) CreateBalance(ctx context.Context, req *pb.BalanceRequest) (*pb.BalanceReply, error) {
	reply, err := s.bc.CreateBalance(ctx, &biz.BalanceRequest{
		Owner:   req.Owner,
		Name:    req.Name,
		Balance: req.Balance,
	})
	if err != nil {
		return nil, err
	}
	return &pb.BalanceReply{
		Message: reply.Message,
		Code:    reply.Code,
		Owner:   reply.Owner,
		Name:    reply.Name,
		Balance: reply.Balance,
	}, nil
}
func (s *BalanceService) UpdateBalance(ctx context.Context, req *pb.BalanceRequest) (*pb.BalanceReply, error) {
	reply, err := s.bc.UpdateBalance(ctx, &biz.BalanceRequest{
		Owner:   req.Owner,
		Name:    req.Name,
		Balance: req.Balance,
	})
	if err != nil {
		return nil, err
	}
	return &pb.BalanceReply{
		Message: reply.Message,
		Code:    reply.Code,
		Owner:   reply.Owner,
		Name:    reply.Name,
		Balance: reply.Balance,
	}, nil
}

func (s *BalanceService) GetBalance(ctx context.Context, req *pb.GetBalanceRequest) (*pb.BalanceReply, error) {
	reply, err := s.bc.GetBalance(ctx, &biz.BalanceRequest{
		Owner: req.Owner,
		Name:  req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &pb.BalanceReply{
		Message: reply.Message,
		Code:    reply.Code,
		Owner:   reply.Owner,
		Name:    reply.Name,
		Balance: reply.Balance,
	}, nil
}
