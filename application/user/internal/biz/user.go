package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRequest struct {
	Email           string
	Password        string
	ConfirmPassword string
}
type LoginRequest struct {
	Email    string
	Password string
}

type UserReply struct {
	UserId int32
}

type UserRepo interface {
	CreateUser(context.Context, *UserRequest) (*UserReply, error)
	LoginUser(context.Context, *LoginRequest) (*UserReply, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, req *UserRequest) (*UserReply, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", req)
	return uc.repo.CreateUser(ctx, req)
}
func (uc *UserUsecase) LoginUser(ctx context.Context, req *LoginRequest) (*UserReply, error) {
	uc.log.WithContext(ctx).Infof("LoginUser: %v", req)
	return uc.repo.LoginUser(ctx, req)
}
