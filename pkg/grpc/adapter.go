package grpc

import (
	"context"
	"net"

	grpcmain "google.golang.org/grpc"
	grpchealth "google.golang.org/grpc/health"
	grpchealthproto "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/edanko/nestix-api/pkg/logs"
)

// Adapter is grpc server app adapter
type Adapter struct {
	name         string
	address      string
	server       *grpcmain.Server
	healthServer *grpchealth.Server
	logger       logs.Logger
}

// NewAdapter provides new primary adapter
func NewAdapter(name, address string, server *grpcmain.Server) *Adapter {
	return &Adapter{
		name:         name,
		address:      address,
		server:       server,
		healthServer: grpchealth.NewServer(),
	}
}

// Start starts grpc application adapter
func (a *Adapter) Start(_ context.Context) error {
	grpchealthproto.RegisterHealthServer(a.server, a.healthServer)

	listener, err := net.Listen("tcp", a.address)
	if err != nil {
		a.logger.Warn().Err(err).Msg("failed to start tcp listener")
		return err
	}

	a.logger.Info().Str("endpoint", a.address).Msg("starting gRPC listener")

	a.healthServer.SetServingStatus(a.name, grpchealthproto.HealthCheckResponse_SERVING)

	return a.server.Serve(listener)
}

// Stop stops grpc application adapter
func (a *Adapter) Stop(_ context.Context) error {
	a.healthServer.SetServingStatus(
		a.name,
		grpchealthproto.HealthCheckResponse_NOT_SERVING,
	)

	a.server.GracefulStop()

	return nil
}
