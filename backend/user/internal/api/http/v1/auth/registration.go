package auth

import "github.com/go-chi/chi/v5"

func RegisterAuthServerHandler(router chi.Router, handler ServerInterface) {
	HandlerFromMuxWithBaseURL(handler, router, "/api/v1")
}
