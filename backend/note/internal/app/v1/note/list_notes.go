package note

import (
	"encoding/json"
	"mime"
	"net/http"

	desc "github.com/knowdateapp/knowdateapp/backend/note/internal/api/http/v1/note"
)

func (i *Implementation) ListNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", mime.TypeByExtension(".json"))

	response := desc.ListNotesResponse{
		Notes: make([]desc.Note, 0),
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}
