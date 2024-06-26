package note

import (
	"context"
	"encoding/json"
	"mime"
	"net/http"

	"github.com/knowdateapp/knowdateapp/backend/note/internal/api/http/code"
	desc "github.com/knowdateapp/knowdateapp/backend/note/internal/api/http/v1/note"
)

func (i *Implementation) GetNoteById(w http.ResponseWriter, r *http.Request, workspace desc.Workspace, noteId desc.NoteId) {
	w.Header().Add("Content-Type", mime.TypeByExtension(".json"))

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	note, err := i.service.GetByID(ctx, workspace, noteId)
	if err != nil {
		// TODO: process error
		w.WriteHeader(http.StatusInternalServerError)
		body := desc.NewDefaultErrorResponse(code.UnexpectedError, "internal server error: %s", err)
		_ = json.NewEncoder(w).Encode(&body)
		return
	}

	response := &desc.GetNoteByIdResponse{
		Id:         note.ID,
		Title:      note.Title,
		Workspace:  note.Workspace,
		ContentUri: note.ContentUri,
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}
