package tests

import (
	"github.com/google/uuid"
	"hermes/channel/message"
	"testing"
)

func TestMessage(t *testing.T) {
	msg := message.New("")
	t.Run("no channel id is set", func(t *testing.T) {
		if msg.GetChannelId() != "" {
			t.Errorf("Expected channel id to be empty, got %s", msg.GetChannelId())
		}
	})
	t.Run("channel id is set", func(t *testing.T) {
		msg.SetChannelId("test")
		if msg.GetChannelId() != "test" {
			t.Errorf("Expected channel id to be test, got %s", msg.GetChannelId())
		}
	})

	t.Run("no request id is set", func(t *testing.T) {
		if msg.GetRequestId() != "" {
			t.Errorf("Expected request id to be empty, got %s", msg.GetRequestId())
		}
	})
	t.Run("request id is set", func(t *testing.T) {
		requestId := uuid.NewString()
		msg.SetRequestId(requestId)
		if msg.GetRequestId() != requestId {
			t.Errorf("Expected request id to be %s, got %s", requestId, msg.GetRequestId())
		}
	})
}
