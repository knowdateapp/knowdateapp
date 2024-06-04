package user

import "github.com/go-chi/chi/v5"

func RegisterUserServerHandler(router chi.Router, handler ServerInterface, baseURL string) {
	HandlerFromMuxWithBaseURL(handler, router, baseURL)
}
