package handlers

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/edanko/nestix-api/pkg/version"
)

func BuildVersionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, version.GetVersion())
	}
}
