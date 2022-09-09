package decorator

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type MetricsClient interface {
	Inc(key string, value int)
}

type commandMetricsDecorator struct {
	base   CommandHandler
	client MetricsClient
}

func (d commandMetricsDecorator) HandlerName() string {
	return d.base.HandlerName()
}

func (d commandMetricsDecorator) NewCommand() any {
	return d.base.NewCommand()
}

func (d commandMetricsDecorator) Handle(ctx context.Context, cmd any) (err error) {
	start := time.Now()

	actionName := d.base.HandlerName()

	defer func() {
		end := time.Since(start)

		d.client.Inc(fmt.Sprintf("commands.%s.duration", actionName), int(end.Seconds()))

		if err == nil {
			d.client.Inc(fmt.Sprintf("commands.%s.success", actionName), 1)
		} else {
			d.client.Inc(fmt.Sprintf("commands.%s.failure", actionName), 1)
		}
	}()

	return d.base.Handle(ctx, cmd)
}

type queryMetricsDecorator[C any, R any] struct {
	base   QueryHandler[C, R]
	client MetricsClient
}

func (d queryMetricsDecorator[C, R]) Handle(ctx context.Context, query C) (result R, err error) {
	start := time.Now()

	actionName := strings.ToLower(generateQueryName(query))

	defer func() {
		end := time.Since(start)

		d.client.Inc(fmt.Sprintf("querys.%s.duration", actionName), int(end.Seconds()))

		if err == nil {
			d.client.Inc(fmt.Sprintf("querys.%s.success", actionName), 1)
		} else {
			d.client.Inc(fmt.Sprintf("querys.%s.failure", actionName), 1)
		}
	}()

	return d.base.Handle(ctx, query)
}
