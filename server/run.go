package server

import (
	"demo/utils/grpchelper"
	"demo/utils/grpchelper/interceptor"
	"net"

	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"

	pb "demo/proto"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type demoService struct {
	pb.UnimplementedDemoServiceServer
}

func newDemoService() *demoService {
	return &demoService{}
}

type server struct {
	grpcServer   *grpc.Server
	healthServer *health.Server
	errorHandler *grpchelper.ErrorHandler
}

func New() *server {
	s := &server{}
	// set error return
	s.errorHandler = grpchelper.NewGRPCErrorHandler()

	unaryInterceptor := []grpc.UnaryServerInterceptor{
		interceptor.ServerRecovery(),
		interceptor.ServerLogging(false),
		interceptor.ServerError(s.errorHandler),
	}

	s.grpcServer = grpc.NewServer(
		grpc.ChainUnaryInterceptor(unaryInterceptor...))

	// grpc health check server
	s.healthServer = health.NewServer()
	healthpb.RegisterHealthServer(s.grpcServer, s.healthServer)
	s.healthServer.SetServingStatus("demo_server", healthpb.HealthCheckResponse_SERVING)

	// demo service
	pb.RegisterDemoServiceServer(s.grpcServer, newDemoService())

	return s
}

func (s *server) Start() {
	log.Info().Msg("GrpcServer starting...")
	s.errorHandler.SetErrorMap(errorMap)

	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}

	err = s.grpcServer.Serve(lis)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}
}

func (s *server) Shutdown() {
	log.Info().Msg("GrpcServer closing")
	s.healthServer.Shutdown()
	s.grpcServer.GracefulStop()
}
