package server

import (
	"github.com/gorilla/handlers"
	v1 "user/api/user/v1"
	"user/internal/conf"
	"user/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	c *conf.Server,
	cc *conf.Casdoor,
	user *service.UserServiceService,
	logger log.Logger,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
		http.Filter(handlers.CORS( // 浏览器跨域
			handlers.AllowedOrigins([]string{"http://localhost:3000", "http://127.0.0.1:3000"}),
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
