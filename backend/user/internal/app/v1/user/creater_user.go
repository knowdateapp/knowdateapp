package user

import (
	"context"
	"encoding/json"
	"mime"
	"net/http"

	"github.com/deepmap/oapi-codegen/v2/pkg/util"

	desc "github.com/knowdateapp/knowdateapp/backend/user/internal/api/http/v1/user"
	"github.com/knowdateapp/knowdateapp/backend/user/internal/domain/models"
)

func (i *Implementation) CreateUser(w http.ResponseWriter, r *http.Request) {
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

	request := desc.CreateUserJSONRequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := desc.DefaultErrorResponse{
			// TODO: make it constant
			Code:    "request-type-error",
			Message: "invalid body",
		}
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	err := i.service.Create(ctx, &models.User{
		Username: request.Username,
		Password: request.Password,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := desc.DefaultErrorResponse{
			// TODO: make it constant
			Code:    "not-created",
			Message: "user was not created",
		}
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
