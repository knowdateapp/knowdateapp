package note

import (
	"bytes"
	"encoding/json"
	"io"
	"mime"
	"net/http"
	"os"
	"path"

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
		return
	}

	const (
		noteField = "note"
		fileField = "file"
	)

	for {
		part, err := reader.NextPart()
		if err != nil {
			if err == io.EOF {
				break
			}
			w.WriteHeader(http.StatusBadRequest)
			resp := desc.DefaultErrorResponse{
				// TODO: make it constant
				Code:    "request-type-error",
				Message: "invalid body",
			}
			_ = json.NewEncoder(w).Encode(resp)
			return
		}

		switch part.FormName() {
		case noteField:
			note := desc.UpdateNotePart{}
			err = json.NewDecoder(part).Decode(&note)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				resp := desc.DefaultErrorResponse{
					// TODO: make it constant
					Code:    "request-type-error",
					Message: "invalid body",
				}
				_ = json.NewEncoder(w).Encode(resp)
				return
			}
		case fileField:
			file := bytes.NewBuffer(nil)
			file.Grow(4048)
			_, err = file.ReadFrom(part)
			if err != nil {
				// TODO: use internal server error
				w.WriteHeader(http.StatusBadRequest)
				resp := desc.DefaultErrorResponse{
					// TODO: make it constant
					Code:    "request-type-error",
					Message: "invalid body",
				}
				_ = json.NewEncoder(w).Encode(resp)
				return
			}
			name := path.Join(".", part.FileName())
			f, err := os.OpenFile(name, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0766)
			if err != nil {
				// TODO: use internal server error
				w.WriteHeader(http.StatusBadRequest)
				resp := desc.DefaultErrorResponse{
					// TODO: make it constant
					Code:    "request-type-error",
					Message: "invalid body",
				}
				_ = json.NewEncoder(w).Encode(resp)
				return
			}
			_, err = file.WriteTo(f)
			if err != nil {
				// TODO: use internal server error
				w.WriteHeader(http.StatusBadRequest)
				resp := desc.DefaultErrorResponse{
					// TODO: make it constant
					Code:    "request-type-error",
					Message: "invalid body",
				}
				_ = json.NewEncoder(w).Encode(resp)
				return
			}
			_ = f.Close()
		}

		_ = part.Close()
	}

	w.WriteHeader(http.StatusOK)
	// TODO: send an answer
}
