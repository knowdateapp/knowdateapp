package note

import (
	"context"
	"encoding/json"
	"mime"
	"net/http"

	"github.com/deepmap/oapi-codegen/v2/pkg/util"

	desc "github.com/knowdateapp/knowdateapp/backend/note/internal/api/http/v1/note"
	"github.com/knowdateapp/knowdateapp/backend/note/internal/domain/models"
)

func (i *Implementation) CreateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", mime.TypeByExtension(".json"))

	if !util.IsMediaTypeJson(r.Header.Get("Content-Type")) {
		w.WriteHeader(http.StatusBadRequest)
		resp := desc.DefaultErrorResponse{
			// TODO: make it constant
			Code:    "request-type-error",
			Message: "invalid content type",
		}
		_ = json.NewEncoder(w).Encode(resp)
	}

	request := desc.CreateNoteJSONRequestBody{}
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

	note, err := i.service.Create(ctx, &models.Note{
		Id:      request.Id.String(),
		Title:   request.Title,
		OwnerId: request.OwnerId.String(),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := desc.DefaultErrorResponse{
			// TODO: make it constant
			Code:    "not-created",
			Message: "note not created",
		}
		_ = json.NewEncoder(w).Encode(resp)
	}

	_ = note

	w.WriteHeader(http.StatusCreated)
}
