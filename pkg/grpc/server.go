package grpc

import (
	"time"

	grpczerolog "github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2"
	grpcctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/edanko/nestix-api/pkg/logs"
)

// ServerConfig provides values for gRPC server configuration
type ServerConfig struct {
	ServerMinTime time.Duration
	ServerTime    time.Duration
	ServerTimeout time.Duration
}

// NewServer provides new grpc server
func NewServer(
	cfg ServerConfig,
	unaryInterceptors []grpc.UnaryServerInterceptor,
	streamInterceptors []grpc.StreamServerInterceptor,
	logger logs.Logger,
) *grpc.Server {
	server := grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             cfg.ServerMinTime, // If a client pings more than once every 5 minutes, terminate the connection
			PermitWithoutStream: true,              // Allow pings even when there are no active streams
		}),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    cfg.ServerTime,    // Ping the client if it is idle for 2 hours to ensure the connection is still active
			Timeout: cfg.ServerTimeout, // Wait 20 second for the ping ack before assuming the connection is dead
		}),
		grpcmiddleware.WithUnaryServerChain(
			append([]grpc.UnaryServerInterceptor{
				grpcctxtags.UnaryServerInterceptor(
					grpcctxtags.WithFieldExtractor(grpcctxtags.CodeGenRequestFieldExtractor),
				),
				logging.UnaryServerInterceptor(grpczerolog.InterceptorLogger(logger.Logger)),

				// middleware.SetMetadataFromUnaryRequest(),
				// firewall.SetIdentityFromUnaryRequest(),
				// middleware.LogUnaryRequest(),
				// middleware.TransformUnaryOutgoingError(),
			}, unaryInterceptors...)...,
		),
		grpcmiddleware.WithStreamServerChain(
			append([]grpc.StreamServerInterceptor{
				grpcctxtags.StreamServerInterceptor(
					grpcctxtags.WithFieldExtractor(grpcctxtags.CodeGenRequestFieldExtractor),
				),
				logging.StreamServerInterceptor(grpczerolog.InterceptorLogger(logger.Logger)),

				// middleware.SetMetadataFromStreamRequest(),
				// firewall.SetIdentityFromStreamRequest(),
				// middleware.LogStreamRequest(),
				// middleware.TransformStreamOutgoingError(),
			}, streamInterceptors...)...,
		),
	)

	return server
}
