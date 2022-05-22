package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	apiserverv1 "github.com/tanmingqian/go-tam/api/apiserver/service/v1"
	v1 "github.com/tanmingqian/go-tam/api/helloworld/v1"
	"github.com/tanmingqian/go-tam/app/apiserver/service/internal/conf"
	"github.com/tanmingqian/go-tam/app/apiserver/service/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.GreeterService, userer *service.UserService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterGreeterServer(srv, greeter)
	apiserverv1.RegisterUserServiceServer(srv, userer)
	return srv
}
