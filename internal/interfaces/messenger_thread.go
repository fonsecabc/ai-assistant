package interfaces

import "time"

type MessengerThread struct {
	ID                             string    `json:"id"`
	CreatedAt                      time.Time `json:"createdAt"`
	Status                         string    `json:"status"`
	OriginalChannelID              string    `json:"originalChannelId"`
	OriginalChannelAccountID       string    `json:"originalChannelAccountId"`
	LatestMessageTimestamp         time.Time `json:"latestMessageTimestamp"`
	LatestMessageSentTimestamp     time.Time `json:"latestMessageSentTimestamp"`
	LatestMessageReceivedTimestamp time.Time `json:"latestMessageReceivedTimestamp"`
	AssignedTo                     string    `json:"assignedTo"`
	Spam                           bool      `json:"spam"`
	InboxID                        string    `json:"inboxId"`
	AssociatedContactID            string    `json:"associatedContactId"`
	Archived                       bool      `json:"archived"`
}
