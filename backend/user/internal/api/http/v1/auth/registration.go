package auth

import "github.com/go-chi/chi/v5"

func RegisterAuthServerHandler(router chi.Router, handler ServerInterface, baseURL string) {
	HandlerFromMuxWithBaseURL(handler, router, baseURL)
}
