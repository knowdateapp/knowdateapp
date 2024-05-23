package note

import (
	"encoding/json"
	"mime"
	"net/http"

	desc "github.com/knowdateapp/knowdateapp/backend/note/internal/api/http/v1/note"
)

func (i *Implementation) UpdateNote(w http.ResponseWriter, r *http.Request, workspace desc.Workspace, noteId desc.NoteId) {
	w.Header().Add("Content-Type", mime.TypeByExtension(".json"))

	reader, err := r.MultipartReader()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := desc.DefaultErrorResponse{
			// TODO: make it constant
			Code:    "request-type-error",
			Message: "invalid body",
		}
		_ = json.NewEncoder(w).Encode(resp)
	}

	_ = reader

	w.WriteHeader(http.StatusOK)
	// TODO: send an answer
}
