package identity

import "context"

type contextKey struct {
	name string
}

func (k *contextKey) String() string { return "context value " + k.name }

var identityContextKey = &contextKey{"identity"}

// ContextWithIdentity returns a new Context that carries value i.
func ContextWithIdentity(ctx context.Context, i *Identity) context.Context {
	if ctx == nil {
		return nil
	}

	if i == nil {
		return ctx
	}

	return context.WithValue(ctx, identityContextKey, i)
}

// FromContext returns the Identity value stored in ctx, if any.
func FromContext(ctx context.Context) (*Identity, bool) {
	if ctx == nil {
		return nil, false
	}

	i, ok := ctx.Value(identityContextKey).(*Identity)

	return i, ok
}
