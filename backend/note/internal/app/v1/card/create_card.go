package card

import (
	"context"
	"encoding/json"
	"mime"
	"net/http"

	"github.com/knowdateapp/knowdateapp/backend/note/internal/api/http/code"
	"github.com/knowdateapp/knowdateapp/backend/note/internal/api/http/common"
	desc "github.com/knowdateapp/knowdateapp/backend/note/internal/api/http/v1/card"
	"github.com/knowdateapp/knowdateapp/backend/note/internal/domain/models"
)

func (i *Implementation) CreateCard(w http.ResponseWriter, r *http.Request, workspace desc.Workspace) {
	w.Header().Add("Content-Type", mime.TypeByExtension(".json"))

	request := desc.CreateCardJSONRequestBody{}
	err := common.DecodeJsonRequest(r, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		body := desc.NewDefaultErrorResponse(code.RequestTypeError, "failed to decode request: %s", err)
		_ = json.NewEncoder(w).Encode(&body)
		return
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	card, err := i.service.Create(ctx, &models.Card{
		Question:  request.Question,
		Answer:    request.Answer,
		Workspace: workspace,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		body := desc.NewDefaultErrorResponse(code.NotCreatedError, "card was not created: %s", err)
		_ = json.NewEncoder(w).Encode(&body)
		return
	}

	response := &desc.CreateCardResponse{
		Id:        card.ID,
		Question:  card.Question,
		Answer:    card.Answer,
		Workspace: card.Workspace,
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(response)
}
