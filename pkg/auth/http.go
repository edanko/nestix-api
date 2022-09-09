package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	apperrors "github.com/edanko/nestix-api/pkg/errors"
	// "firebase.google.com/go/v4/auth"
)

type HttpMiddleware struct {
	JWSValidator
}

func (a HttpMiddleware) Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// ctx := r.Context()

			// bearerToken := a.tokenFromHeader(r)
			// if bearerToken == "" {
			// 	httperr.Unauthorised("empty-bearer-token", nil, w, r)
			// 	return nil
			// }

			// jws, err := GetJWSFromRequest(c.Request())
			// if err != nil {
			// 	return fmt.Errorf("getting jws: %w", err)
			// }
			//
			// // if the JWS is valid, we have a JWT, which will contain a bunch of claims.
			// token, err := a.ValidateJWS(jws)
			// if err != nil {
			// 	return fmt.Errorf("validating JWS: %w", err)
			// }

			// token, err := a.AuthClient.VerifyIDToken(ctx, bearerToken)
			// if err != nil {
			// 	httperr.Unauthorised("unable-to-verify-jwt", err, w, r)
			// 	return
			// }

			// it's always a good idea to use custom type as context value (in this case ctxKey)
			// because nobody from the outside of the package will be able to override/read this value
			// ctx := context.WithValue(c.Request().Context(), userContextKey, User{
			// 	UUID:        token.UID,
			// 	Email:       token.Claims["email"].(string),
			// 	Role:        token.Claims["role"].(string),
			// 	DisplayName: token.Claims["name"].(string),
			// })
			// c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}

func (a HttpMiddleware) tokenFromHeader(r *http.Request) string {
	headerValue := r.Header.Get("Authorization")

	if len(headerValue) > 7 && strings.ToLower(headerValue[0:6]) == "bearer" {
		return headerValue[7:]
	}

	return ""
}

type User struct {
	UUID  string
	Email string
	Role  string

	DisplayName string
}

type ctxKey int

const (
	userContextKey ctxKey = iota
)

var (
	NoUserInContextError = apperrors.NewAuthorizationError("no user in context", "no-user-found")
)

func UserFromCtx(ctx context.Context) (User, error) {
	u, ok := ctx.Value(userContextKey).(User)
	if ok {
		return u, nil
	}

	return User{}, NoUserInContextError
}
