package mq

import "github.com/hibiken/asynq"

type MessagePayload struct {
	Msg  string         `json:"message"`
	Opts []asynq.Option `json:"-"`
}
