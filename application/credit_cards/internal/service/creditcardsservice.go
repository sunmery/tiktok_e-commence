package service

import (
	"context"
	"credit_cards/internal/biz"

	pb "credit_cards/api/credit_cards/v1"
)

type CreditCardsServiceService struct {
	pb.UnimplementedCreditCardsServiceServer

	cs *biz.CreditCardsUsecase
}

func NewCreditCardsServiceService(cs *biz.CreditCardsUsecase) *CreditCardsServiceService {
	return &CreditCardsServiceService{cs: cs}
}

func (s *CreditCardsServiceService) CreateCreditCard(ctx context.Context, req *pb.CreditCards) (*pb.CardsReply, error) {
	result, err := s.cs.CreateCreditCards(ctx, &biz.CreditCards{
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
	return &pb.CardsReply{
		Message: result.Message,
		Code:    result.Code,
	}, nil
}
func (s *CreditCardsServiceService) UpdateCreditCard(ctx context.Context, req *pb.UpdateCreditCardsRequest) (*pb.CardsReply, error) {
	result, err := s.cs.UpdateCreditCards(ctx, &biz.CreditCards{
		CreditCardNumber:          req.CreditCardNumber,
		CreditCardCvv:             req.CreditCardCvv,
		CreditCardExpirationYear:  req.CreditCardExpirationYear,
		CreditCardExpirationMonth: req.CreditCardExpirationMonth,
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
	result, err := s.cs.DeleteCreditCards(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.CardsReply{
		Message: result.Message,
		Code:    result.Code,
	}, nil
}
func (s *CreditCardsServiceService) GetCreditCards(ctx context.Context, req *pb.GetCreditCardsRequest) (*pb.GetCreditCardsReply, error) {
	creditCards, err := s.cs.GetCreditCards(ctx, req.CreditCardNumber)
	if err != nil {
		return nil, err
	}
	creditCardList := make([]*pb.CreditCards, len(creditCards))
	for i, card := range creditCards {
		creditCardList[i] = &pb.CreditCards{
			CreditCardNumber:          card.CreditCardNumber,
			CreditCardCvv:             card.CreditCardCvv,
			CreditCardExpirationYear:  card.CreditCardExpirationYear,
			CreditCardExpirationMonth: card.CreditCardExpirationMonth,
		}
	}

	return &pb.GetCreditCardsReply{
		CreditCards: creditCardList,
	}, nil
}
func (s *CreditCardsServiceService) ListCreditCards(ctx context.Context, req *pb.ListCreditCardsRequest) (*pb.ListCreditCardsReply, error) {
	creditCards, err := s.cs.GetCreditCards(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	creditCardList := make([]*pb.CreditCards, len(creditCards))
	for i, card := range creditCards {
		creditCardList[i] = &pb.CreditCards{
			CreditCardNumber:          card.CreditCardNumber,
			CreditCardCvv:             card.CreditCardCvv,
			CreditCardExpirationYear:  card.CreditCardExpirationYear,
			CreditCardExpirationMonth: card.CreditCardExpirationMonth,
		}
	}

	return &pb.ListCreditCardsReply{
		CreditCards: creditCardList,
	}, nil
}
