package card

import (
	"context"
	"encoding/json"
	"mime"
	"net/http"

	"github.com/knowdateapp/knowdateapp/backend/note/internal/api/http/code"
	desc "github.com/knowdateapp/knowdateapp/backend/note/internal/api/http/v1/card"
)

func (i *Implementation) ListCards(w http.ResponseWriter, r *http.Request, workspace desc.Workspace) {
	w.Header().Add("Content-Type", mime.TypeByExtension(".json"))

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	cards, err := i.service.GetByWorkspace(ctx, workspace)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		body := desc.NewDefaultErrorResponse(code.UnexpectedError, "can not get cards by workspace")
		_ = json.NewEncoder(w).Encode(&body)
		return
	}

	response := desc.ListCardsResponse{
		Cards: make([]desc.Card, 0, len(cards)),
	}

	for _, card := range cards {
		response.Cards = append(response.Cards, desc.Card{
			Id:        card.ID,
			Question:  card.Question,
			Answer:    card.Answer,
			Workspace: card.Workspace,
		})
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}
