package note

import (
	"context"
	"encoding/json"
	"mime"
	"net/http"

	desc "github.com/knowdateapp/knowdateapp/backend/note/internal/api/http/v1/note"
	"github.com/knowdateapp/knowdateapp/backend/note/internal/domain/models"
)

func (i *Implementation) CreateNote(w http.ResponseWriter, r *http.Request, workspace desc.Workspace) {
	w.Header().Add("Content-Type", mime.TypeByExtension(".json"))

	request := desc.CreateNoteJSONBody{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := desc.DefaultErrorResponse{
			// TODO: make it constant
			Code:    "request-type-error",
			Message: "invalid body",
		}
		_ = json.NewEncoder(w).Encode(resp)
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	id, err := i.service.Create(ctx, &models.Note{
		Title:     request.Title,
		Workspace: workspace,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := desc.DefaultErrorResponse{
			// TODO: make it constant
			Code:    "not-created",
			Message: "note was not created",
		}
		_ = json.NewEncoder(w).Encode(resp)
	}

	response := &desc.CreateNoteResponse{
		Id: id,
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(response)
}
