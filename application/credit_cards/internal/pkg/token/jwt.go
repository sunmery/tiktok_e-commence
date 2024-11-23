package token

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtV5 "github.com/golang-jwt/jwt/v5"
)

type Payload struct {
	ID    string `json:"id"`
	Name  string `json:"username"`
	Owner string `json:"owner"`
	Type  string `json:"type"`

	jwtV5.RegisteredClaims
}

func ParseJWTFromContext(ctx context.Context) (*Payload, error) {
	payload, ok := jwt.FromContext(ctx)
	if !ok {
		return nil, errors.New("token not found")
	}
	user := payload.(*Payload)
	return user, nil
}
