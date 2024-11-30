package biz

import (
	"context"
	"time"

	v1 "addresses/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Address represents a user's address information.
type Address struct {
	UserID        string `json:"user_id"`
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	ZipCode       int32  `json:"zip_code"`
}

type Reply struct {
	Message string `json:"message"`
	Code    uint   `json:"code"`
}

// AddressesRepo is a Greater repo.
type AddressesRepo interface {
	CreateAddress(ctx context.Context, req Address) (Reply, error)
	UpdateAddress(ctx context.Context, req Address) (Reply, error)
	DeleteAddress(ctx context.Context, req Address) (Reply, error)
	GetAddresses(ctx context.Context, userId string) (Reply, error)
}

// AddressesUsecase is a Addresses usecase.
type AddressesUsecase struct {
	repo AddressesRepo
	log  *log.Helper
}

// NewAddressesUsecase new a Addresses usecase.
func NewAddressesUsecase(repo AddressesRepo, logger log.Logger) *AddressesUsecase {
	return &AddressesUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateAddress creates a Addresses, and returns the new Addresses.
func (uc *AddressesUsecase) CreateAddress(ctx context.Context, req Address) (Reply, error) {
	uc.log.WithContext(ctx).Infof("CreateAddresses: %v", req)
	return uc.repo.CreateAddress(ctx, req)
}
