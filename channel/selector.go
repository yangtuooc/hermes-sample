package channel

import (
	"context"
	"hermes/channel/message"
)

type SelectedVendor func(vendor Vendor) error

// VendorSelector 是一个选择 Vendor 的策略，一个通道可以有多个Vendor，VendorSelector 用于选择一个 Vendor来发送消息
type VendorSelector interface {
	Select(ctx context.Context, message *message.Message, vendors Vendors, selected SelectedVendor) error
}

// the default is a round-robin selector
var _ VendorSelector = (*roundRobinSelector)(nil)

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

func NewRoundRobinSelector() VendorSelector {
	return &roundRobinSelector{}
}
