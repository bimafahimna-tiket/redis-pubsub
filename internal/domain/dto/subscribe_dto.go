package dto

type SubscribeToChannelRequest struct {
	Channel string `json:"channel"`
}

type UnsubscribeToChannelRequest struct {
	Channel string `json:"channel"`
}
