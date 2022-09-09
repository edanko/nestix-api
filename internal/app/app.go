package app

import "context"

type mediator interface {
	Send(ctx context.Context, cmd any) error
}

type Application struct {
	CommandBus mediator
	Queries    Queries
}

type Queries struct {
	// HourAvailability      query.HourAvailabilityHandler
}
