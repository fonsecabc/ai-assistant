package handlers

import (
	"ai-assistant/domain/usecases"
	"ai-assistant/internal/interfaces"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MessengerWebhookHandlerReq []interfaces.MessengerWebhookEvent

type MessengerWebhookHandler struct {
	AnswerMessageUseCase *usecases.AnswerMessageUseCase
}

func NewMessengerWebhookHandler(auc *usecases.AnswerMessageUseCase) *MessengerWebhookHandler {
	return &MessengerWebhookHandler{
		AnswerMessageUseCase: auc,
	}
}

func (h *MessengerWebhookHandler) Handle(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	var req MessengerWebhookHandlerReq
	err = json.Unmarshal(b, &req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for _, event := range req {
		switch event.SubscriptionType {
		case "conversation.newMessage":
			// Remove this if clause after tests
			if event.PortalID == 7577254786 {
				err := h.AnswerMessageUseCase.Perform(fmt.Sprint(event.MessageID), fmt.Sprint(event.PortalID))
				if err != nil {
					fmt.Println("ERROR ANSWERING MESSAGE: ", err)
					http.Error(w, "Unable to answer message", http.StatusBadRequest)
					return
				}
			}
		}
	}
}
