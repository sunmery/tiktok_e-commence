package service

import (
	pb "addresses/api/addresses/v1"
	"addresses/internal/biz"
	"addresses/internal/pkg/token"
	"context"
	"errors"
)

type AddressesServiceService struct {
	pb.UnimplementedAddressesServiceServer

	as *biz.AddressesUsecase
}

func NewAddressesServiceService(as *biz.AddressesUsecase) *AddressesServiceService {
	return &AddressesServiceService{as: as}
}

func (s *AddressesServiceService) CreateAddresses(ctx context.Context, req *pb.Address) (*pb.Address, error) {
	address, err := s.as.CreateAddress(ctx, &biz.Address{
		Owner:         req.Owner,
		Name:          req.Name,
		StreetAddress: req.StreetAddress,
		City:          req.City,
		State:         req.State,
		Country:       req.Country,
		ZipCode:       req.ZipCode,
	})
	if err != nil {
		return nil, err
	}
	return &pb.Address{
		Id:            address.Id,
		Owner:         address.Owner,
		Name:          address.Name,
		City:          address.City,
		State:         address.State,
		Country:       address.Country,
		ZipCode:       address.ZipCode,
		StreetAddress: address.StreetAddress,
	}, nil

}
func (s *AddressesServiceService) UpdateAddresses(ctx context.Context, req *pb.Address) (*pb.Address, error) {
	payload, err := token.ExtractPayload(ctx)
	if err != nil {
		return nil, err
	}

	if req.Owner != payload.Owner || req.Name != payload.Name {
		return nil, errors.New("invalid token")
	}

	address, err := s.as.UpdateAddress(ctx, &biz.Address{
		Id:            req.Id,
		Owner:         payload.Owner,
		Name:          payload.Name,
		StreetAddress: req.StreetAddress,
		City:          req.City,
		State:         req.State,
		Country:       req.Country,
		ZipCode:       req.ZipCode,
	})
	if err != nil {
		return nil, err
	}
	return &pb.Address{
		Id:            address.Id,
		Owner:         address.Owner,
		Name:          address.Name,
		StreetAddress: address.StreetAddress,
		City:          address.City,
		State:         address.State,
		Country:       address.Country,
		ZipCode:       address.ZipCode,
	}, nil

}
func (s *AddressesServiceService) DeleteAddresses(ctx context.Context, req *pb.DeleteAddressesRequest) (*pb.DeleteAddressesReply, error) {
	reply, err := s.as.DeleteAddress(ctx, &biz.DeleteAddressesRequest{
		AddressId: uint32(req.AddressesId),
		Owner:     req.Owner,
		Name:      req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &pb.DeleteAddressesReply{
		Message: reply.Message,
		Id:      reply.Id,
		Code:    reply.Code,
	}, nil
}
func (s *AddressesServiceService) GetAddresses(ctx context.Context, req *pb.GetAddressesRequest) (*pb.GetAddressesReply, error) {
	payload, err := token.ExtractPayload(ctx)
	if err != nil {
		return nil, err
	}

	if req.Owner != payload.Owner || req.Name != payload.Name {
		return nil, errors.New("invalid token")
	}

	addresses, err := s.as.GetAddresses(ctx, &biz.Request{
		Owner: payload.Owner,
		Name:  payload.Name,
	})
	if err != nil {
		return nil, err
	}
	addressList := make([]*pb.Address, len(addresses.Addresses))
	for i, address := range addresses.Addresses {
		addressList[i] = &pb.Address{
			Id:            address.Id,
			Owner:         address.Owner,
			Name:          address.Name,
			City:          address.City,
			State:         address.State,
			Country:       address.Country,
			ZipCode:       address.ZipCode,
			StreetAddress: address.StreetAddress,
		}
	}

	return &pb.GetAddressesReply{
		Addresses: addressList,
	}, nil
}
