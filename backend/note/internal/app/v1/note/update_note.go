package note

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime"
	"mime/multipart"
	"net/http"

	"github.com/knowdateapp/knowdateapp/backend/note/internal/api/http/code"
	"github.com/knowdateapp/knowdateapp/backend/note/internal/api/http/common"
	desc "github.com/knowdateapp/knowdateapp/backend/note/internal/api/http/v1/note"
	"github.com/knowdateapp/knowdateapp/backend/note/internal/domain/models"
)

func (i *Implementation) UpdateNote(w http.ResponseWriter, r *http.Request, workspace desc.Workspace, noteId desc.NoteId) {
	w.Header().Add("Content-Type", mime.TypeByExtension(".json"))

	reader, err := r.MultipartReader()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		body := desc.NewDefaultErrorResponse(code.RequestTypeError, "failed to decode request: %s", err)
		_ = json.NewEncoder(w).Encode(&body)
		return
	}

	const (
		noteField = "note"
		fileField = "file"
	)

	var (
		notePart desc.UpdateNotePart
		file     *bytes.Buffer
	)

	for {
		var part *multipart.Part
		part, err = reader.NextPart()
		if err != nil {
			if err == io.EOF {
				break
			}
			w.WriteHeader(http.StatusBadRequest)
			body := desc.NewDefaultErrorResponse(code.RequestTypeError, "failed to decode multipart request: %s", err)
			_ = json.NewEncoder(w).Encode(&body)
			return
		}

		switch part.FormName() {
		case noteField:
			err = common.DecodeJsonMultipart(part, &notePart)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				body := desc.NewDefaultErrorResponse(code.RequestTypeError, "failed to decode note field: %s", err)
				_ = json.NewEncoder(w).Encode(&body)
				_ = part.Close()
				return
			}
		case fileField:
			file = bytes.NewBuffer(make([]byte, 0, 4048))
			n := int64(0)

			n, err = file.ReadFrom(part)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				body := desc.NewDefaultErrorResponse(code.UnexpectedError, "failed to decode file field: %s", err)
				_ = json.NewEncoder(w).Encode(&body)
				_ = part.Close()
				return
			}
			if n == 0 {
				file = nil
			}
		}

		_ = part.Close()
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	note, err := i.service.Update(ctx, &models.Note{
		ID:        noteId,
		Title:     notePart.Title,
		Workspace: workspace,
	}, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		body := desc.NewDefaultErrorResponse(code.NotUpdatedError, "note was not updated: %s", err)
		_ = json.NewEncoder(w).Encode(&body)
		return
	}

	response := &desc.UpdateNoteResponse{
		Id:         note.ID,
		Title:      note.Title,
		Workspace:  note.Workspace,
		ContentUri: note.ContentUri,
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}
