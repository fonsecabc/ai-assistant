package interfaces

type MessengerWebhookEvent struct {
	EventID          int64  `json:"eventId"`
	SubscriptionID   int    `json:"subscriptionId"`
	PortalID         int    `json:"portalId"`
	AppID            int    `json:"appId"`
	OccurredAt       int64  `json:"occurredAt"`
	SubscriptionType string `json:"subscriptionType"`
	AttemptNumber    int    `json:"attemptNumber"`
	ObjectID         int64  `json:"objectId"`
	MessageID        string `json:"messageId"`
	MessageType      string `json:"messageType"`
	ChangeFlag       string `json:"changeFlag"`
}
