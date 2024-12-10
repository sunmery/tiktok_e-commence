package server

import (
	"context"
	v1 "credit_cards/api/credit_cards/v1"
	"credit_cards/internal/conf"
	"credit_cards/internal/service"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	jwtV5 "github.com/golang-jwt/jwt/v5"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.27.0"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, ac *conf.Auth, tr *conf.Trace, creditCards *service.CreditCardsServiceService, logger log.Logger) *grpc.Server {
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
	publicKey, err := parseRSAPublicKeyFromPEM([]byte(ac.Jwt.ApiKey))
	if err != nil {
		panic("failed to parse public key")
	}
	var opts = []grpc.ServerOption{
		grpc.Middleware(
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
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterCreditCardsServiceServer(srv, creditCards)
	return srv
}
