package http

import (
	"context"
	"errors"
	"net"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/edanko/nestix-api/pkg/logs"
)

// EchoAdapter is http echo app adapter
type EchoAdapter struct {
	address string
	echo    *echo.Echo
	logger  logs.Logger
}

// NewEchoAdapter provides new primary HTTP adapter
func NewEchoAdapter(server *echo.Echo, addr string, logger logs.Logger) *EchoAdapter {
	return &EchoAdapter{
		address: addr,
		echo:    server,
		logger:  logger,
	}
}

// Start starts http application adapter
func (a *EchoAdapter) Start(ctx context.Context) error {
	a.echo.Server.BaseContext = func(_ net.Listener) context.Context { return ctx }

	a.echo.HideBanner = true
	a.echo.HidePort = true

	a.logger.Info(
		"starting HTTP listener",
		map[string]any{
			"address": a.address,
		},
	)
	// a.echo.StartTLS(a.address, "cert.pem", "key.pem")
	// if err := a.echo.Start(a.address); !errors.Is(err, http.ErrServerClosed) {
	if err := a.echo.StartTLS(a.address, "cert.pem", "key.pem"); !errors.Is(err, http.ErrServerClosed) {
		a.logger.Fatal("failed to start HTTPs server", err, nil)
		return err
	}
	return nil
}

// Stop stops http application adapter
func (a *EchoAdapter) Stop(ctx context.Context) error {
	return a.echo.Shutdown(ctx)
}
