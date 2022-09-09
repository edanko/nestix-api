package http

import (
	"context"
	"errors"
	"net"
	"net/http"

	"github.com/edanko/nestix-api/pkg/logs"
)

// StdlibAdapter is http echo app adapter
type StdlibAdapter struct {
	httpServer *http.Server
	logger     *logs.Logger
}

// NewStdlibAdapter provides new primary HTTP adapter
func NewStdlibAdapter(httpServer *http.Server, logger *logs.Logger) *StdlibAdapter {
	return &StdlibAdapter{
		httpServer: httpServer,
		logger:     logger,
	}
}

// Start starts http application adapter
func (a *StdlibAdapter) Start(ctx context.Context) error {
	a.httpServer.BaseContext = func(_ net.Listener) context.Context { return ctx }

	a.logger.Info().Str("endpoint", a.httpServer.Addr).Msg("starting HTTP listener")
	if err := a.httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		a.logger.Panic().Err(err).Msg("failed to start HTTP echo")
		return err
	}
	return nil
}

// Stop stops http application adapter
func (a *StdlibAdapter) Stop(ctx context.Context) error {
	return a.httpServer.Shutdown(ctx)
}
