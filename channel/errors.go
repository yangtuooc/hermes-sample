package channel

import "fmt"

var ErrChannelNotFound = func(channelId string) error {
	if channelId == "" {
		return fmt.Errorf("not specified channel id")
	}
	return fmt.Errorf("channel not found: %s", channelId)
}
