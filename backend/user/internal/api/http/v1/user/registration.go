package user

import "github.com/go-chi/chi/v5"

func RegisterUserServerHandler(router chi.Router, handler ServerInterface) {
	HandlerFromMuxWithBaseURL(handler, router, "/api/v1")
}
