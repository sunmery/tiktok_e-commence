package server

import (
	"context"
	v1 "credit_cards/api/credit_cards/v1"
	"credit_cards/internal/conf"
	"credit_cards/internal/service"
	"crypto/rsa"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtV5 "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/handlers"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.27.0"
)

// NewWhiteListMatcher 创建jwt白名单
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	// whiteList["/admin.v1.AdminService/Login"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// parseRSAPublicKeyFromPEM 解析 RSA 公钥
func parseRSAPublicKeyFromPEM(pemBytes []byte) (*rsa.PublicKey, error) {
	publicKey, err := jwtV5.ParseRSAPublicKeyFromPEM(pemBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSA public key: %w", err)
	}
	return publicKey, nil
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, ac *conf.Auth, tr *conf.Trace, creditCards *service.CreditCardsServiceService, logger log.Logger) *http.Server {
	publicKey, err := parseRSAPublicKeyFromPEM([]byte(ac.Jwt.ApiKey))
	if err != nil {
		panic("failed to parse public key")
	}

	// trace start
	ctx := context.Background()

	res, err := resource.New(ctx,
		resource.WithAttributes(
			// The service name used to display traces in backends
			// serviceName,
			semconv.ServiceNameKey.String(tr.Jaeger.ServiceName),
			attribute.String("exporter", "otlptracehttp"),
			attribute.String("environment", "dev"),
			attribute.Float64("float", 312.23),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	// shutdownTracerProvider, err := initTracerProvider(ctx, res, tr.Jaeger.Http.Endpoint)
	_, err2 := initTracerProvider(ctx, res, tr.Jaeger.Http.Endpoint)
	if err2 != nil {
		log.Fatal(err)
	}
	// trace end

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(), // trace 链路追踪
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
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Cookie"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"http://localhost:3000"}),
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
	v1.RegisterCreditCardsServiceHTTPServer(srv, creditCards)
	return srv
}
