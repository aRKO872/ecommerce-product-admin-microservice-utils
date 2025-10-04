package routers

import (
	"log"
	"net"
	"strconv"

	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/literals"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"

	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func (s *ServiceRouter) ServeGRPC(register func(s *grpc.Server), interceptorList ...grpc.UnaryServerInterceptor) {
	var cfg grpcConfig
	getEnv(&cfg)

	addr := cfg.Host + ":" + cfg.Port
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen to addr : %s", err.Error())
	}

	reqSize, err := strconv.Atoi(cfg.ReqRespSize)
	if err != nil {
		reqSize = literals.DefaultGRPCRequestRespSize
	}

	interceptors := append([]grpc.UnaryServerInterceptor{
		grpc_recovery.UnaryServerInterceptor(),
	}, interceptorList...)

	gserver := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(interceptors...)),
		grpc.MaxRecvMsgSize(reqSize*literals.MegaByte),
		grpc.MaxSendMsgSize(reqSize*literals.MegaByte),
	)

	register(gserver)

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(gserver, healthServer)

	s.wg.Add(1)
	go func ()  {
		log.Printf("gRPC server listening at %v\n", listener.Addr().String())
		if err := gserver.Serve(listener); err != nil {
			log.Fatalf("failed to serve gRPC server: %v", err.Error())
		}
	}()
}