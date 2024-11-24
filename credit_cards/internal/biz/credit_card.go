package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// CreditCards is a CreditCards model.
type CreditCards struct {
	Owner                     string `json:"owner"`
	Username                  string `json:"username"`
	CreditCardNumber          string `json:"credit_card_number"`
	CreditCardCvv             int32  `json:"credit_card_cvv"`
	CreditCardExpirationYear  int32  `json:"credit_card_expiration_year"`
	CreditCardExpirationMonth int32  `json:"credit_card_expiration_month"`
}

type GetCreditCardsRequest struct {
	Owner            string `json:"owner"`
	Username         string `json:"username"`
	CreditCardNumber string `json:"credit_card_number"`
}
type ListCreditCardsRequest struct {
	Owner    string `json:"owner"`
	Username string `json:"username"`
}

type CreditCardsReply struct {
	Message string `json:"message"`
	Code    int32  `json:"code"`
}

// CreditCardsRepo is a Greater repo.
type CreditCardsRepo interface {
	CreateCreditCard(ctx context.Context, req *CreditCards) (*CreditCardsReply, error)
	UpdateCreditCard(ctx context.Context, req *CreditCards) (*CreditCardsReply, error)
	DeleteCreditCard(ctx context.Context, id int32) (*CreditCardsReply, error)
	GetCreditCard(ctx context.Context, req *GetCreditCardsRequest) ([]*CreditCards, error)
	ListCreditCards(ctx context.Context, req *ListCreditCardsRequest) ([]*CreditCards, error)
}

// CreditCardsUsecase is a CreditCards usecase.
type CreditCardsUsecase struct {
	repo CreditCardsRepo
	log  *log.Helper
}

// NewCreditCardsUsecase new a CreditCards usecase.
func NewCreditCardsUsecase(repo CreditCardsRepo, logger log.Logger) *CreditCardsUsecase {
	return &CreditCardsUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CreditCardsUsecase) CreateCreditCards(ctx context.Context, req *CreditCards) (*CreditCardsReply, error) {
	uc.log.WithContext(ctx).Infof("CreateCreditCards: %+v\n", req)
	return uc.repo.CreateCreditCard(ctx, req)
}
func (uc *CreditCardsUsecase) UpdateCreditCards(ctx context.Context, req *CreditCards) (*CreditCardsReply, error) {
	uc.log.WithContext(ctx).Infof("UpdateCreditCards: %+v\n", req)
	return uc.repo.UpdateCreditCard(ctx, req)
}
func (uc *CreditCardsUsecase) DeleteCreditCards(ctx context.Context, id int32) (*CreditCardsReply, error) {
	uc.log.WithContext(ctx).Infof("DeleteCreditCards: %+v\n", id)
	return uc.repo.DeleteCreditCard(ctx, id)
}
func (uc *CreditCardsUsecase) GetCreditCards(ctx context.Context, req *GetCreditCardsRequest) ([]*CreditCards, error) {
	uc.log.WithContext(ctx).Infof("GetCreditCards: %+v\n", req)
	return uc.repo.GetCreditCard(ctx, req)
}
func (uc *CreditCardsUsecase) ListCreditCards(ctx context.Context, req *ListCreditCardsRequest) ([]*CreditCards, error) {
	uc.log.WithContext(ctx).Infof("ListCreditCards: %+v\n", req)
	return uc.repo.ListCreditCards(ctx, req)
}
