package dto

type MessageRequest struct {
	Msg string `json:"message"`
}

type MessagePubSubRequest struct {
	Msg     string `json:"message"`
	Channel string `json:"channel"`
}

type GetAllCacheResponse struct {
	Cache []string `json:"cache"`
}

type UpdateCacheRequest struct {
	Cache     string `json:"cache"`
	Channel   string `json:"channel"`
	Operation string `json:"operation"`
}
