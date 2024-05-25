package note

import (
	"context"
	"encoding/json"
	"mime"
	"net/http"

	desc "github.com/knowdateapp/knowdateapp/backend/note/internal/api/http/v1/note"
)

func (i *Implementation) ListNotes(w http.ResponseWriter, r *http.Request, workspace desc.Workspace) {
	w.Header().Add("Content-Type", mime.TypeByExtension(".json"))

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	notes, err := i.service.GetByWorkspace(ctx, workspace)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := desc.DefaultErrorResponse{
			// TODO: make it constant
			Code:    "unexpected-error",
			Message: "can not get notes by workspace",
		}
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	response := desc.ListNotesResponse{
		Notes: make([]desc.Note, 0, len(notes)),
	}

	for _, note := range notes {
		response.Notes = append(response.Notes, desc.Note{
			Id:         note.ID,
			Title:      note.Title,
			Workspace:  note.Workspace,
			ContentUri: note.ContentUri,
		})
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}
