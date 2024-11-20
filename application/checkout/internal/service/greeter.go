package service

import (
	"context"

	v1 "checkout/api/checkout/v1"
	"checkout/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	cc *biz.CheckoutUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

func (s *GreeterService) CreateService(ctx context.Context, req *v1.CreateServiceRequest) (*v1.CreateServiceReply, error) {
	s.cc.CreateService(ctx, &biz.CreateServiceRequest)
}
