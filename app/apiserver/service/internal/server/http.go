package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	apiserverv1 "github.com/tanmingqian/go-tam/api/apiserver/service/v1"
	v1 "github.com/tanmingqian/go-tam/api/helloworld/v1"
	"github.com/tanmingqian/go-tam/app/apiserver/service/internal/conf"
	"github.com/tanmingqian/go-tam/app/apiserver/service/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, userer *service.UserService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	apiserverv1.RegisterUserServiceHTTPServer(srv, userer)
	return srv
}
