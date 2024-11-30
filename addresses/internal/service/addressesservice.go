package service

import (
	"addresses/internal/biz"
	"context"
	"net/http"

	pb "addresses/api/addresses/v1"
)

type AddressesServiceService struct {
	pb.UnimplementedAddressesServiceServer

	as *biz.AddressUsecase
}

func NewAddressesServiceService(as *biz.AddressUsecase) *AddressesServiceService {
	return &AddressesServiceService{as: as}
}

func (s *AddressesServiceService) CreateAddresses(ctx context.Context, req *pb.CreateAddressesRequest) (*pb.Reply, error) {

	return &pb.Reply{
		Message: "Create Address OK",
		Code:    http.StatusOK,
	}, nil
}
func (s *AddressesServiceService) UpdateAddresses(ctx context.Context, req *pb.UpdateAddressesRequest) (*pb.Reply, error) {
	return &pb.Reply{}, nil
}
func (s *AddressesServiceService) DeleteAddresses(ctx context.Context, req *pb.DeleteAddressesRequest) (*pb.Reply, error) {
	return &pb.Reply{}, nil
}
func (s *AddressesServiceService) GetAddresses(ctx context.Context, req *pb.GetAddressesRequest) (*pb.Addresses, error) {
	return &pb.Addresses{}, nil
}
