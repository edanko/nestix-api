package decorator

import (
	"context"
)

func ApplyCommandDecorators(
	handler CommandHandler,
	// logger *logs.Logger,
	metricsClient MetricsClient,
) CommandHandler {
	// return commandLoggingDecorator{
	// 	base: commandMetricsDecorator{
	// 		base:   handler,
	// 		client: metricsClient,
	// 	},
	// 	logger: logger,
	// }
	return commandMetricsDecorator{
		base:   handler,
		client: metricsClient,
	}
}

type CommandHandler interface {
	HandlerName() string
	NewCommand() any
	Handle(ctx context.Context, cmd any) error
}
