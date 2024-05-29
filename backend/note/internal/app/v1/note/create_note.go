package note

import (
	"context"
	"encoding/json"
	"mime"
	"net/http"

	"github.com/knowdateapp/knowdateapp/backend/note/internal/api/http/code"
	"github.com/knowdateapp/knowdateapp/backend/note/internal/api/http/common"
	desc "github.com/knowdateapp/knowdateapp/backend/note/internal/api/http/v1/note"
	"github.com/knowdateapp/knowdateapp/backend/note/internal/domain/models"
)

func (i *Implementation) CreateNote(w http.ResponseWriter, r *http.Request, workspace desc.Workspace) {
	w.Header().Add("Content-Type", mime.TypeByExtension(".json"))

	request := desc.CreateNoteJSONRequestBody{}
	err := common.DecodeJsonRequest(r, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		body := desc.NewDefaultErrorResponse(code.RequestTypeError, "failed to decode request: %s", err)
		_ = json.NewEncoder(w).Encode(&body)
		return
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	note, err := i.service.Create(ctx, &models.Note{
		Title:     request.Title,
		Workspace: workspace,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		body := desc.NewDefaultErrorResponse(code.NotCreatedError, "note was not created: %s", err)
		_ = json.NewEncoder(w).Encode(&body)
		return
	}

	response := &desc.CreateNoteResponse{
		Id:         note.ID,
		Title:      note.Title,
		Workspace:  note.Workspace,
		ContentUri: note.ContentUri,
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(response)
}
