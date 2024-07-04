package card

import (
	"context"
	"encoding/json"
	"mime"
	"net/http"

	"github.com/knowdateapp/knowdateapp/backend/note/internal/api/http/code"
	desc "github.com/knowdateapp/knowdateapp/backend/note/internal/api/http/v1/card"
)

func (i *Implementation) GetCardById(w http.ResponseWriter, r *http.Request, workspace desc.Workspace, cardId desc.CardId) {
	w.Header().Add("Content-Type", mime.TypeByExtension(".json"))

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	card, err := i.service.GetByID(ctx, workspace, cardId)
	if err != nil {
		// TODO: process error
		w.WriteHeader(http.StatusInternalServerError)
		body := desc.NewDefaultErrorResponse(code.UnexpectedError, "internal server error: %s", err)
		_ = json.NewEncoder(w).Encode(&body)
		return
	}

	response := &desc.GetCardByIdResponse{
		Id:        card.ID,
		Question:  card.Question,
		Answer:    card.Answer,
		Workspace: card.Workspace,
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}
