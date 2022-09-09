package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

// Key to use when setting the request ID.
type ctxKeyRequestID int

// requestIDKey is the key that holds the unique request ID in a request context.
const requestIDKey ctxKeyRequestID = 0

// requestIDHeader is the name of the HTTP Header which contains the request id.
var requestIDHeader = "X-Request-Id"

// RequestID is a middleware that injects a request ID into the context of each
// request. A request ID is a string of the UUID.
func RequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := r.Header.Get(requestIDHeader)
		if requestID == "" {
			requestID = uuid.NewString()
		}
		ctx = context.WithValue(ctx, requestIDKey, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

// GetRequestID returns a request ID from the given context if one is present.
// Returns the empty string if a request ID cannot be found.
func GetRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if requestID, ok := ctx.Value(requestIDKey).(string); ok {
		return requestID
	}
	return ""
}
