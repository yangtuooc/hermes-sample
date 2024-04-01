package channel

import (
	"context"
	"hermes/message"
)

type SelectedVendor func(vendor Vendor) error

// Selector is a strategy to select a vendor to send a message
type Selector interface {
	Select(ctx context.Context, message *message.Message, vendors Vendors, selected SelectedVendor) error
}

// the default is a round-robin selector
var _ Selector = (*roundRobinSelector)(nil)

type roundRobinSelector struct {
}

func (r *roundRobinSelector) Select(ctx context.Context, message *message.Message, vendors Vendors, selected SelectedVendor) error {
	var err error
	for _, vendor := range vendors {
		if err = selected(vendor); err == nil {
			return nil
		}
	}
	return err
}

func NewRoundRobinSelector() Selector {
	return &roundRobinSelector{}
}
