package server

import (
	"context"
	"crypto/rsa"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	jwtV5 "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/handlers"
	v1 "user/api/users/v1"
	"user/internal/conf"
	"user/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewWhiteListMatcher 创建jwt白名单
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	// whiteList["/admin.v1.AdminService/Login"] = struct{}{}
	whiteList["/api.user.v1.UserService/Signin"] = struct{}{}
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
func NewHTTPServer(
	c *conf.Server,
	ac *conf.Auth,
	user *service.UserServiceService,
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
			handlers.AllowedOrigins([]string{"http://localhost:3000", "http://127.0.0.1:3000", "http://127.0.0.1:443", "https://node1.apikv.com"}),
			handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "DELETE"}),
			handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
			handlers.AllowCredentials(),
		)),
		http.RequestDecoder(MultipartFormDataDecoder),
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
	v1.RegisterUserServiceHTTPServer(srv, user)
	return srv
}

func MultipartFormDataDecoder(r *http.Request, v interface{}) error {
	// 从Request Header的Content-Type中提取出对应的解码器
	_, ok := http.CodecForRequest(r, "Content-Type")
	// 如果找不到对应的解码器此时会报错
	if !ok {
		r.Header.Set("Content-Type", "application/json")
		// return errors.BadRequest("CODEC", r.Header.Get("Content-Type"))
	}
	// fmt.Printf("method:%s\n", r.Method)
	// if r.Method == "POST" {
	// 	data, err := ioutil.ReadAll(r.Body)
	// 	if err != nil {
	// 		return errors.BadRequest("CODEC", err.Error())
	// 	}
	// 	if err = codec.Unmarshal(data, v); err != nil {
	// 		return errors.BadRequest("CODEC", err.Error())
	// 	}
	// }

	return nil
}
