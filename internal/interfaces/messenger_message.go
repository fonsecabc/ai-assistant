package interfaces

import "time"

type MessengerMessage struct {
	ID                    string    `json:"id"`
	ConversationsThreadID string    `json:"conversationsThreadId"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
	CreatedBy             string    `json:"createdBy"`
	Client                struct {
		ClientType string `json:"clientType"`
	} `json:"client"`
	Senders []struct {
		ActorID string `json:"actorId"`
	} `json:"senders"`
	Recipients []struct {
		ActorID            string `json:"actorId"`
		DeliveryIdentifier struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"deliveryIdentifier"`
	} `json:"recipients"`
	Archived         bool   `json:"archived"`
	Text             string `json:"text"`
	RichText         string `json:"richText"`
	Attachments      []any  `json:"attachments"`
	TruncationStatus string `json:"truncationStatus"`
	Status           struct {
		StatusType string `json:"statusType"`
	} `json:"status"`
	Direction        string `json:"direction"`
	ChannelID        string `json:"channelId"`
	ChannelAccountID string `json:"channelAccountId"`
	Type             string `json:"type"`
}
