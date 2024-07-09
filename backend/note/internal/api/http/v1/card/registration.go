package card

import "github.com/go-chi/chi/v5"

func RegisterCardServerHandler(router chi.Router, handler ServerInterface, baseURL string) {
	HandlerFromMuxWithBaseURL(handler, router, baseURL)
}
