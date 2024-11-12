package data

import (
	"context"
	"database/sql"
	"errors"
	"user/internal/biz"
	"user/internal/pkg"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (r *userRepo) CreateUser(ctx context.Context, req *biz.UserRequest) (*biz.UserReply, error) {
	user, err := r.data.CreateUser(ctx, CreateUserParams{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, errors.New(pkg.PgErrorCode(err))
	}

	return &biz.UserReply{
		UserId: int32(user.ID),
	}, nil
}

func (r *userRepo) LoginUser(ctx context.Context, request *biz.LoginRequest) (*biz.UserReply, error) {
	user, err := r.data.LoginUser(ctx, request.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("该邮件在数据库中不存在")
		}
		return nil, errors.New(pkg.PgErrorCode(err))
	}
	// TODO: 添加密码校验
	return &biz.UserReply{UserId: int32(user.ID)}, nil
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
