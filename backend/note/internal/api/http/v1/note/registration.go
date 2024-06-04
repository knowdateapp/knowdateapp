package note

import "github.com/go-chi/chi/v5"

func RegisterNoteServerHandler(router chi.Router, handler ServerInterface, baseURL string) {
	HandlerFromMuxWithBaseURL(handler, router, baseURL)
}
