package adapters

import (
	"ai-assistant/internal/interfaces"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MessengerAdapter struct {
	MessengerAccessToken string
	MessengerBaseURL     string
}

type MessengerSendMessageRequest struct {
	Type             string `json:"type"`
	Text             string `json:"text"`
	ChannelId        string `json:"channelId"`
	ChannelAccountId string `json:"channelAccountId"`
	SenderActorId    string `json:"senderActorId"`
}

func NewMessengerAdapter(messengerAccessToken, messengerAPIBaseURL string) *MessengerAdapter {
	return &MessengerAdapter{
		MessengerAccessToken: messengerAccessToken,
		MessengerBaseURL:     messengerAPIBaseURL,
	}
}

func (a *MessengerAdapter) MakeRequest(method, path string, body io.Reader) (*http.Response, error) {
	fmt.Println("MAKING REQUEST: ", method, a.MessengerBaseURL+path)
	req, err := http.NewRequest(method, a.MessengerBaseURL+path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+a.MessengerAccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *MessengerAdapter) GetMessage(threadId, messageid string) (interfaces.MessengerMessage, error) {
	resp, err := a.MakeRequest(http.MethodGet, "/conversations/v3/conversations/threads/"+threadId+"/messages/"+messageid, nil)
	if err != nil {
		fmt.Println("GET MESSAGE REQUEST ERROR: ", err)
		return interfaces.MessengerMessage{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR READING GET MESSAGE REQUEST BODY: ", err)
		return interfaces.MessengerMessage{}, err
	}

	m := interfaces.MessengerMessage{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		fmt.Println("ERROR UNMARSHALLING GET MESSAGE REQUEST BODY: ", err)
	}

	return m, err
}

func (a *MessengerAdapter) SendMessage(threadId, messageType, text, channelId, channelAccountId, senderActorId string) error {
	data, err := json.Marshal(MessengerSendMessageRequest{
		Type:             messageType,
		Text:             text,
		ChannelId:        channelId,
		ChannelAccountId: channelAccountId,
		SenderActorId:    senderActorId,
	})
	if err != nil {
		fmt.Println("ERROR MARSHALLING SEND MESSAGE REQUEST: ", err)
		return err
	}

	resp, err := a.MakeRequest(http.MethodPost, "/conversations/v3/conversations/threads/"+threadId+"/messages", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("SEND MESSAGE REQUEST ERROR: ", err)
		return err
	}
	// Implement error handler

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR READING SEND MESSAGE REQUEST BODY: ", err)
		return err
	}

	fmt.Println("RESPONSE BODY: ", string(body))

	m := interfaces.MessengerMessage{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		fmt.Println("ERROR UNMARSHALLING SEND MESSAGE REQUEST BODY: ", err)
	}

	return err
}
