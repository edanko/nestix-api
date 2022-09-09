package decorator

import (
	"context"
	"fmt"

	"github.com/edanko/nestix-api/pkg/logs"
)

type commandLoggingDecorator struct {
	base   CommandHandler
	logger *logs.Logger
}

func (d commandLoggingDecorator) HandlerName() string {
	return d.base.HandlerName()
}

func (d commandLoggingDecorator) NewCommand() any {
	return d.base.NewCommand()
}

func (d commandLoggingDecorator) Handle(ctx context.Context, cmd any) (err error) {
	logger := d.logger.With().
		Str("command", d.base.HandlerName()).
		Str("command_body", fmt.Sprintf("%#v", cmd)).
		Logger()

	logger.Debug().Msg("Executing command")
	defer func() {
		if err == nil {
			logger.Info().Msg("Command executed successfully")
		} else {
			logger.Error().Err(err).Msg("Failed to execute command")
		}
	}()

	return d.base.Handle(ctx, cmd)
}

type queryLoggingDecorator[C any, R any] struct {
	base   QueryHandler[C, R]
	logger *logs.Logger
}

func (d queryLoggingDecorator[C, R]) Handle(ctx context.Context, query C) (result R, err error) {
	logger := d.logger.With().
		Str("query", generateQueryName(query)).
		Str("query_body", fmt.Sprintf("%#v", query)).
		Logger()

	logger.Debug().Msg("Executing query")
	defer func() {
		if err == nil {
			logger.Info().Msg("Query executed successfully")
		} else {
			logger.Error().Err(err).Msg("Failed to execute query")
		}
	}()

	return d.base.Handle(ctx, query)
}
