package server

import (
	"tiktok/api/user"
	"tiktok/internal/conf"
	"tiktok/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.UserService, logger log.Logger) *grpc.Server {
	// logger
	// exporter, err := stdouttrace.New(stdouttrace.WithWriter(ioutil.Discard))
	// if err != nil {
	// 	fmt.Printf("creating stdout exporter: %v", err)
	// 	panic(err)
	// }
	// tp := tracesdk.NewTracerProvider(
	// 	tracesdk.WithBatcher(exporter),
	// 	tracesdk.WithResource(resource.NewSchemaless(
	// 		semconv.ServiceNameKey.String(Name)),
	// 	))

	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			// tracing.Server(tracing.WithTracerProvider(tp)),

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
	user.RegisterUserServer(srv, greeter)
	return srv
}
