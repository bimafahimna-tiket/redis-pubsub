package dto

type MessageRequest struct {
	Msg string `json:"message"`
}

type MessagePubSubRequest struct {
	Msg     string `json:"message"`
	Channel string `json:"channel"`
}
