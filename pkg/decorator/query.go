package decorator

import (
	"context"
	"fmt"
	"strings"
)

func ApplyQueryDecorators[H any, R any](
	handler QueryHandler[H, R],
	// logger *logs.Logger,
	metricsClient MetricsClient,
) QueryHandler[H, R] {
	// return queryLoggingDecorator[H, R]{
	// 	base: queryMetricsDecorator[H, R]{
	// 		base:   handler,
	// 		client: metricsClient,
	// 	},
	// 	logger: logger,
	// }
	return queryMetricsDecorator[H, R]{
		base:   handler,
		client: metricsClient,
	}
}

type QueryHandler[Q any, R any] interface {
	Handle(ctx context.Context, query Q) (R, error)
}

func generateQueryName(handler any) string {
	return strings.Split(fmt.Sprintf("%T", handler), ".")[1]
}
