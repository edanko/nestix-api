package tenant

import "context"

type contextKey struct {
	name string
}

func (k *contextKey) String() string { return "context value " + k.name }

// tenantContextKey is how tenant id value is stored/retrieved.
var tenantContextKey = &contextKey{"tenant_id"}

// ContextWithTenantID returns a new Context that carries tenant id.
func ContextWithTenantID(ctx context.Context, tenantID string) context.Context {
	if ctx == nil {
		return nil
	}
	if tenantID == "" {
		return ctx
	}

	return context.WithValue(ctx, tenantContextKey, tenantID)
}

// FromContext returns the tenant id value stored in ctx, if any.
func FromContext(ctx context.Context) (string, bool) {
	if ctx == nil {
		return "", false
	}

	t, ok := ctx.Value(tenantContextKey).(string)

	return t, ok
}

// MustFromContext returns the tenant id value stored in ctx, panics if not found or ctx is nil.
func MustFromContext(ctx context.Context) string {
	if ctx == nil {
		panic("context must not be nil")
	}

	t, ok := ctx.Value(tenantContextKey).(string)
	if !ok {
		panic("tenant id not found in context")
	}

	return t
}
