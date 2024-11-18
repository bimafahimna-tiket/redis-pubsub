package pubsub

import (
	"encoding/json"
	"poc-redis-pubsub/internal/pkg/logger"
	"time"
)

type Payload struct {
	UniqueID  int64  `json:"unique_id"`
	Type      string `json:"type"`
	Operation string `json:"operation"`
	Msg       string `json:"msg"`
}

func NewJsonPayload(typePayload, operation, cache string) string {
	var payload = Payload{UniqueID: time.Now().Unix(), Type: typePayload, Operation: operation, Msg: cache}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		logger.Log.Error("failed to marshal payload")
		return ""
	}
	return string(jsonData)
}
func ToJson(payload Payload) string {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		logger.Log.Error("failed to marshal payload")
		return ""
	}
	return string(jsonData)
}
