package message

import (
	"github.com/google/uuid"
	"time"
)

const (
	messageId = "messageId"
	timestamp = "timestamp"
)

type headers map[string]any

func (h headers) MessageId() string {
	return h[messageId].(string)
}

func (h headers) Timestamp() int64 {
	return h[timestamp].(int64)
}

func newHeaders() headers {
	return headers{
		messageId: uuid.NewString(),
		timestamp: time.Now().Unix(),
	}
}
