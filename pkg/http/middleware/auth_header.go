package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/edanko/nestix-api/pkg/identity"
)

const (
	HeaderEmail  = "Remote-Email"
	HeaderGroups = "Remote-Groups"
	HeaderName   = "Remote-Name"
	HeaderUser   = "Remote-User"
)

type IdentityConfig struct {
	Skipper            func(c echo.Context) bool
	TargetHeaderEmail  string
	TargetHeaderGroups string
	TargetHeaderName   string
	TargetHeaderUser   string
}

var DefaultIdentityConfig = IdentityConfig{
	Skipper:            middleware.DefaultSkipper,
	TargetHeaderEmail:  HeaderEmail,
	TargetHeaderGroups: HeaderGroups,
	TargetHeaderName:   HeaderName,
	TargetHeaderUser:   HeaderUser,
}

func Identity() echo.MiddlewareFunc {
	return IdentityWithConfig(DefaultIdentityConfig)
}

func IdentityWithConfig(config IdentityConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultIdentityConfig.Skipper
	}
	if config.TargetHeaderEmail == "" {
		config.TargetHeaderEmail = HeaderEmail
	}
	if config.TargetHeaderGroups == "" {
		config.TargetHeaderGroups = HeaderGroups
	}
	if config.TargetHeaderName == "" {
		config.TargetHeaderName = HeaderName
	}
	if config.TargetHeaderUser == "" {
		config.TargetHeaderUser = HeaderUser
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()

			i := identity.Identity{
				Email:  req.Header.Get(config.TargetHeaderEmail),
				Groups: strings.Split(req.Header.Get(config.TargetHeaderGroups), ","),
				Name:   req.Header.Get(config.TargetHeaderName),
				User:   req.Header.Get(config.TargetHeaderUser),
			}

			c.SetRequest(req.WithContext(identity.ContextWithIdentity(req.Context(), &i)))

			return next(c)
		}
	}
}
