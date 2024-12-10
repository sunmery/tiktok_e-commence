package server

import (
	v1 "checkout/api/checkout/v1"
	"checkout/internal/conf"
	"checkout/internal/service"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	jwtV5 "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/handlers"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	c *conf.Server,
	ac *conf.Auth,
	checkout *service.CheckoutService,
	logger log.Logger,
) *http.Server {
	publicKey, err := parseRSAPublicKeyFromPEM([]byte(ac.Jwt.ApiKey))
	if err != nil {
		panic("failed to parse public key")
	}
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			selector.Server(
				jwt.Server(
					func(token *jwtV5.Token) (interface{}, error) {
						// 检查是否使用了正确的签名方法
						if _, ok := token.Method.(*jwtV5.SigningMethodRSA); !ok {
							return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
						}
						return publicKey, nil
					},
					jwt.WithSigningMethod(jwtV5.SigningMethodRS256),
				),
			).
				Match(NewWhiteListMatcher()).Build(),
		),
		http.Filter(handlers.CORS( // 浏览器跨域
			handlers.AllowedOrigins([]string{"http://localhost:3000", "http://127.0.0.1:3000"}),
			handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "DELETE"}),
			handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
			handlers.AllowCredentials(),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterCheckoutServiceHTTPServer(srv, checkout)
	return srv
}
