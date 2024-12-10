package data

import (
	"context"
	"credit_cards/internal/data/models"
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

func (c *creditCardsRepo) GetCreditCard(ctx context.Context, req *biz.GetCreditCardsRequest) (*biz.CreditCards, error) {
	payload, err := token.ExtractPayload(ctx)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	if req.Owner != payload.Owner || req.Name != payload.Name {
		return nil, errors.New("invalid token")
	}

	card, err := c.data.db.GetCreditCards(ctx, models.GetCreditCardsParams{
		Owner:  payload.Owner,
		Name:   payload.Name,
		Number: req.Number,
	})
	if err != nil {
		return nil, err
	}

	return &biz.CreditCards{
		Owner:           card.Owner,
		Name:            card.Name,
		Number:          card.Number,
		Cvv:             card.Cvv,
		ExpirationYear:  card.ExpirationYear,
		ExpirationMonth: card.ExpirationMonth,
	}, nil
}
func (c *creditCardsRepo) SearchCreditCards(ctx context.Context, req *biz.GetCreditCardsRequest) ([]*biz.CreditCards, error) {
	payload, err := token.ExtractPayload(ctx)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	if req.Owner != payload.Owner || req.Name != payload.Name {
		return nil, errors.New("invalid token")
	}

	cards, err := c.data.db.SearchCreditCards(ctx, models.SearchCreditCardsParams{
		Owner:  payload.Owner,
		Name:   payload.Name,
		Number: &req.Number,
	})
	if err != nil {
		return nil, err
	}
	cardList := make([]*biz.CreditCards, len(cards))
	for i, card := range cards {
		cardList[i] = &biz.CreditCards{
			Owner:           card.Owner,
			Name:            card.Name,
			Number:          card.Number,
			Cvv:             card.Cvv,
			ExpirationYear:  card.ExpirationYear,
			ExpirationMonth: card.ExpirationMonth,
		}
	}
	return cardList, nil
}

func (c *creditCardsRepo) CreateCreditCard(ctx context.Context, req *biz.CreditCards) (*biz.CreditCardsReply, error) {
	params := models.CreateCreditCardParams{
		Owner:           req.Owner,
		Name:            req.Name,
		Number:          req.Number,
		Cvv:             req.Cvv,
		ExpirationYear:  req.ExpirationYear,
		ExpirationMonth: req.ExpirationMonth,
	}
	fmt.Printf("params:%+v\n", params)
	_, err := c.data.db.CreateCreditCard(ctx, params)
	if err != nil {
		return nil, err
	}
	return &biz.CreditCardsReply{
		Message: "OK",
		Code:    http.StatusOK,
	}, nil
}

func (c *creditCardsRepo) UpdateCreditCard(ctx context.Context, req *biz.CreditCards) (*biz.CreditCardsReply, error) {
	result, err := c.data.db.UpdateCreditCard(ctx, models.UpdateCreditCardParams{
		Number:          &req.Number,
		Cvv:             &req.Cvv,
		ExpirationYear:  &req.ExpirationYear,
		ExpirationMonth: &req.ExpirationMonth,
		Owner:           req.Owner,
		Name:            req.Name,
	})
	if err != nil {
		return nil, err
	}

	return &biz.CreditCardsReply{
		Message: string(result.ID),
		Code:    http.StatusOK,
	}, nil
}

func (c *creditCardsRepo) DeleteCreditCard(ctx context.Context, req *biz.DeleteCreditCardsRequest) (*biz.CreditCardsReply, error) {
	result, err := c.data.db.DeleteCreditCard(ctx, models.DeleteCreditCardParams{
		Owner: req.Owner,
		Name:  req.Name,
		ID:    int32(req.Id),
	})
	if err != nil {
		return nil, err
	}

	return &biz.CreditCardsReply{
		Message: string(result.ID),
		Code:    http.StatusOK,
	}, nil
}

func (c *creditCardsRepo) ListCreditCards(ctx context.Context, req *biz.CreditCardsRequest) ([]*biz.CreditCards, error) {
	cards, lErr := c.data.db.ListCreditCards(ctx, models.ListCreditCardsParams{
		Owner: req.Owner,
		Name:  req.Name,
	})
	if lErr != nil {
		return nil, lErr
	}
	cardList := make([]*biz.CreditCards, len(cards))
	for i, card := range cards {
		cardList[i] = &biz.CreditCards{
			Id:              uint32(card.ID),
			Owner:           card.Owner,
			Name:            card.Name,
			Number:          card.Number,
			Cvv:             card.Cvv,
			ExpirationYear:  card.ExpirationYear,
			ExpirationMonth: card.ExpirationMonth,
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
