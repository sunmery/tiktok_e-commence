package data

import (
	"context"
	"credit_cards/internal/data/modules"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"net/http"

	"credit_cards/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type creditCardsRepo struct {
	data *Data
	log  *log.Helper
}

func (c *creditCardsRepo) GetCreditCard(ctx context.Context, req *biz.GetCreditCardsRequest) ([]*biz.CreditCards, error) {
	token, ok := jwt.FromContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "token not found")
	}
	fmt.Printf("token %v", token)
	cards, err := c.data.db.GetCreditCards(ctx, modules.GetCreditCardsParams{
		Owner:            req.Owner,
		Username:         req.Username,
		CreditCardNumber: &req.CreditCardNumber,
	})
	if err != nil {
		return nil, err
	}
	cardList := make([]*biz.CreditCards, len(cards))
	for i, card := range cards {
		cardList[i] = &biz.CreditCards{
			Owner:                     card.Owner,
			Username:                  card.Username,
			CreditCardNumber:          card.CreditCardNumber,
			CreditCardCvv:             card.CreditCardCvv,
			CreditCardExpirationYear:  card.CreditCardExpirationYear,
			CreditCardExpirationMonth: card.CreditCardExpirationMonth,
		}
	}
	return cardList, nil
}

func (c *creditCardsRepo) CreateCreditCard(ctx context.Context, req *biz.CreditCards) (*biz.CreditCardsReply, error) {

	_, err := c.data.db.CreateCreditCard(ctx, modules.CreateCreditCardParams{
		Owner:                     req.Owner,
		Username:                  req.Username,
		CreditCardNumber:          req.CreditCardNumber,
		CreditCardCvv:             req.CreditCardCvv,
		CreditCardExpirationYear:  req.CreditCardExpirationYear,
		CreditCardExpirationMonth: req.CreditCardExpirationMonth,
	})
	if err != nil {
		return nil, err
	}
	return &biz.CreditCardsReply{
		Message: "OK",
		Code:    http.StatusOK,
	}, nil
}

func (c *creditCardsRepo) UpdateCreditCard(ctx context.Context, req *biz.CreditCards) (*biz.CreditCardsReply, error) {
	// TODO implement me
	panic("implement me")
}

func (c *creditCardsRepo) DeleteCreditCard(ctx context.Context, id int32) (*biz.CreditCardsReply, error) {
	// TODO implement me
	panic("implement me")
}

func (c *creditCardsRepo) ListCreditCards(ctx context.Context, req *biz.ListCreditCardsRequest) ([]*biz.CreditCards, error) {
	cards, err := c.data.db.ListCreditCards(ctx, modules.ListCreditCardsParams{
		Owner:    req.Owner,
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}
	cardList := make([]*biz.CreditCards, len(cards))
	for i, card := range cards {
		cardList[i] = &biz.CreditCards{
			Owner:                     card.Owner,
			Username:                  card.Username,
			CreditCardNumber:          card.CreditCardNumber,
			CreditCardCvv:             card.CreditCardCvv,
			CreditCardExpirationYear:  card.CreditCardExpirationYear,
			CreditCardExpirationMonth: card.CreditCardExpirationMonth,
		}
	}
	return cardList, nil
}

func NewCreditCardsRepo(data *Data, logger log.Logger) biz.CreditCardsRepo {
	return &creditCardsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}