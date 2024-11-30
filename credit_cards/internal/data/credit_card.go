package data

import (
	"context"
	"credit_cards/internal/data/modules"
	"credit_cards/internal/pkg/token"
	"errors"
	"fmt"
	"net/http"

	"credit_cards/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type creditCardsRepo struct {
	data *Data
	log  *log.Helper
}

func (c *creditCardsRepo) GetCreditCard(ctx context.Context, req *biz.GetCreditCardsRequest) ([]*biz.CreditCards, error) {
	payload, err := token.ExtractPayload(ctx)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	if req.Owner != payload.Owner || req.Username != payload.Name {
		return nil, errors.New("invalid token")
	}

	cards, err := c.data.db.GetCreditCards(ctx, models.GetCreditCardsParams{
		Owner:            payload.Owner,
		Username:         payload.Name,
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
	payload, err := token.ExtractPayload(ctx)
	if err != nil {
		return nil, errors.New("invalid token")
	}
	fmt.Printf("req: %+v \n", req)
	fmt.Printf("payload %+v \n", payload)
	if req.Owner != payload.Owner || req.Username != payload.Name {
		fmt.Println("OK")
		return nil, errors.New("invalid token")
	}

	_, err = c.data.db.CreateCreditCard(ctx, models.CreateCreditCardParams{
		Owner:                     payload.Owner,
		Username:                  payload.Name,
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
	payload, err := token.ExtractPayload(ctx)
	if err != nil {
		return nil, err
	}

	if req.Owner != payload.Owner || req.Username != payload.Name {
		return nil, errors.New("invalid token")
	}

	_, err = c.data.db.UpdateCreditCard(ctx, models.UpdateCreditCardParams{
		CreditCardNumber:          &req.CreditCardNumber,
		CreditCardCvv:             &req.CreditCardCvv,
		CreditCardExpirationYear:  &req.CreditCardExpirationYear,
		CreditCardExpirationMonth: &req.CreditCardExpirationMonth,
		Username:                  payload.Name,
	})

	return &biz.CreditCardsReply{
		Message: "OK",
		Code:    http.StatusOK,
	}, nil
}

func (c *creditCardsRepo) DeleteCreditCard(ctx context.Context, id int32) (*biz.CreditCardsReply, error) {
	payload, err := token.ExtractPayload(ctx)
	if err != nil {
		return nil, err
	}

	_, err = c.data.db.DeleteCreditCard(ctx, models.DeleteCreditCardParams{
		Username: payload.Name,
		ID:       id,
	})

	return &biz.CreditCardsReply{
		Message: "OK",
		Code:    http.StatusOK,
	}, nil
}

func (c *creditCardsRepo) ListCreditCards(ctx context.Context, req *biz.ListCreditCardsRequest) ([]*biz.CreditCards, error) {
	payload, err := token.ExtractPayload(ctx)
	if err != nil {
		return nil, err
	}
	if req.Owner != payload.Owner || req.Username != payload.Name {
		return nil, errors.New("invalid token")
	}
	cards, err := c.data.db.ListCreditCards(ctx, models.ListCreditCardsParams{
		Owner:    payload.Owner,
		Username: payload.Name,
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
