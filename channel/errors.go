package channel

import "fmt"

var ErrChannelNotFound = func(message string) error {
	return fmt.Errorf("channel not found: %s", message)
}
