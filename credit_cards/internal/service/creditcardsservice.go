package service

import (
	"context"
	pb "credit_cards/api/credit_cards/v1"
	"credit_cards/internal/biz"
	"credit_cards/internal/pkg/token"
	"errors"
)

type CreditCardsServiceService struct {
	pb.UnimplementedCreditCardsServiceServer

	cs *biz.CreditCardsUsecase
}

func NewCreditCardsServiceService(cs *biz.CreditCardsUsecase) *CreditCardsServiceService {
	return &CreditCardsServiceService{cs: cs}
}

func (s *CreditCardsServiceService) CreateCreditCard(ctx context.Context, req *pb.CreditCards) (*pb.CardsReply, error) {
	payload, err := token.ExtractPayload(ctx)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	if req.Owner != payload.Owner || req.Name != payload.Name {
		return nil, errors.New("invalid token")
	}
	result, cErr := s.cs.CreateCreditCards(ctx, &biz.CreditCards{
		Owner:           req.Owner,
		Name:            req.Name,
		Number:          req.Number,
		Cvv:             req.Cvv,
		ExpirationYear:  req.ExpirationYear,
		ExpirationMonth: req.ExpirationMonth,
	})
	if cErr != nil {
		return nil, cErr
	}
	return &pb.CardsReply{
		Message: result.Message,
		Code:    result.Code,
	}, nil
}
func (s *CreditCardsServiceService) UpdateCreditCard(ctx context.Context, req *pb.CreditCards) (*pb.CardsReply, error) {
	payload, err := token.ExtractPayload(ctx)
	if err != nil {
		return nil, err
	}

	if req.Owner != payload.Owner || req.Name != payload.Name {
		return nil, errors.New("invalid token")
	}
	result, err := s.cs.UpdateCreditCards(ctx, &biz.CreditCards{
		Number:          req.Number,
		Cvv:             req.Cvv,
		ExpirationYear:  req.ExpirationYear,
		ExpirationMonth: req.ExpirationMonth,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CardsReply{
		Message: result.Message,
		Code:    result.Code,
	}, nil
}
func (s *CreditCardsServiceService) DeleteCreditCard(ctx context.Context, req *pb.DeleteCreditCardsRequest) (*pb.CardsReply, error) {
	payload, err := token.ExtractPayload(ctx)
	if err != nil {
		return nil, err
	}

	result, dErr := s.cs.DeleteCreditCards(ctx, &biz.DeleteCreditCardsRequest{
		Owner: payload.Owner,
		Name:  payload.Name,
		Id:    req.Id,
	})
	if dErr != nil {
		return nil, dErr
	}
	return &pb.CardsReply{
		Message: result.Message,
		Code:    result.Code,
	}, nil
}
func (s *CreditCardsServiceService) GetCreditCard(ctx context.Context, req *pb.GetCreditCardsRequest) (*pb.GetCreditCardsReply, error) {
	payload, err := token.ExtractPayload(ctx)
	if err != nil {
		return nil, err
	}
	if req.Owner != payload.Owner || req.Name != payload.Name {
		return nil, errors.New("invalid token")
	}
	creditCard, gErr := s.cs.GetCreditCard(ctx, &biz.GetCreditCardsRequest{
		Owner:  payload.Owner,
		Name:   payload.Name,
		Number: req.Number,
	})
	if gErr != nil {
		return nil, gErr
	}

	return &pb.GetCreditCardsReply{
		CreditCards: &pb.CreditCards{
			Number:          creditCard.Number,
			Cvv:             creditCard.Cvv,
			ExpirationYear:  creditCard.ExpirationYear,
			ExpirationMonth: creditCard.ExpirationMonth,
		},
	}, nil
}
func (s *CreditCardsServiceService) SearchCreditCards(ctx context.Context, req *pb.GetCreditCardsRequest) (*pb.SearchCreditCardsReply, error) {
	payload, err := token.ExtractPayload(ctx)
	if err != nil {
		return nil, err
	}
	if req.Owner != payload.Owner || req.Name != payload.Name {
		return nil, errors.New("invalid token")
	}
	creditCards, gErr := s.cs.SearchCreditCards(ctx, &biz.GetCreditCardsRequest{
		Owner:  payload.Owner,
		Name:   payload.Name,
		Number: req.Number,
	})
	if gErr != nil {
		return nil, gErr
	}
	creditCardList := make([]*pb.CreditCards, len(creditCards))
	for i, card := range creditCards {
		creditCardList[i] = &pb.CreditCards{
			Number:          card.Number,
			Cvv:             card.Cvv,
			ExpirationYear:  card.ExpirationYear,
			ExpirationMonth: card.ExpirationMonth,
		}
	}

	return &pb.SearchCreditCardsReply{
		CreditCards: creditCardList,
	}, nil
}
func (s *CreditCardsServiceService) ListCreditCards(ctx context.Context, req *pb.ListCreditCardsRequest) (*pb.ListCreditCardsReply, error) {
	payload, err := token.ExtractPayload(ctx)
	if err != nil {
		return nil, err
	}
	creditCards, lErr := s.cs.ListCreditCards(ctx, &biz.CreditCardsRequest{
		Owner: payload.Owner,
		Name:  payload.Name,
	})
	if lErr != nil {
		return nil, lErr
	}
	creditCardList := make([]*pb.CreditCards, len(creditCards))
	for i, card := range creditCards {
		creditCardList[i] = &pb.CreditCards{
			Number:          card.Number,
			Cvv:             card.Cvv,
			ExpirationYear:  card.ExpirationYear,
			ExpirationMonth: card.ExpirationMonth,
			Owner:           card.Owner,
			Name:            card.Name,
			Id:              card.Id,
		}
	}

	return &pb.ListCreditCardsReply{
		CreditCards: creditCardList,
	}, nil
}
