package auth

import (
	"context"
	"encoding/json"
	"mime"
	"net/http"

	"github.com/deepmap/oapi-codegen/v2/pkg/util"

	desc "github.com/knowdateapp/knowdateapp/backend/user/internal/api/http/v1/auth"
)

func (i *Implementation) Login(w http.ResponseWriter, r *http.Request) {
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

	request := desc.LoginJSONRequestBody{}
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

	token, err := i.service.CreateToken(ctx, request.Username, request.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		resp := desc.DefaultErrorResponse{
			// TODO: make it constant
			Code:    "unauthorized",
			Message: "user unauthorized",
		}
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	response := &desc.LoginResponse{
		Token: token,
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(response)
}
