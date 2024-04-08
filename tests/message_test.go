package tests

import (
	"github.com/google/uuid"
	"hermes/channel/message"
	"testing"
)

func TestMessage(t *testing.T) {
	msg := message.New("")
	t.Run("no channel id is set", func(t *testing.T) {
		if msg.ChannelId() != "" {
			t.Errorf("Expected channel id to be empty, got %s", msg.ChannelId())
		}
	})
	t.Run("channel id is set", func(t *testing.T) {
		msg.SetChannelId("test")
		if msg.ChannelId() != "test" {
			t.Errorf("Expected channel id to be test, got %s", msg.ChannelId())
		}
	})

	t.Run("no request id is set", func(t *testing.T) {
		if msg.RequestId() != "" {
			t.Errorf("Expected request id to be empty, got %s", msg.RequestId())
		}
	})
	t.Run("request id is set", func(t *testing.T) {
		requestId := uuid.NewString()
		msg.SetRequestId(requestId)
		if msg.RequestId() != requestId {
			t.Errorf("Expected request id to be %s, got %s", requestId, msg.RequestId())
		}
	})
}
