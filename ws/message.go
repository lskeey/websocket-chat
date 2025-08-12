package ws

type WsMessage struct {
	RecipientID uint   `json:"recipient_id"`
	Content     string `json:"content"`
}
