package tests

import (
	"context"
	"hermes/channel"
	"hermes/channel/message"
)

var _ channel.VendorSelector = (*retrySelector)(nil)

type retrySelector struct {
	retries int
}

func (r *retrySelector) Select(ctx context.Context, message *message.Message, vendors channel.Vendors, selected channel.SelectedVendor) error {
	var vendor channel.Vendor
	if len(vendors) >= 1 {
		vendor = vendors.First()
	}
	var err error
	for i := 0; i < r.retries; i++ {
		if err = selected(vendor); err == nil {
			return nil
		}
	}
	return err
}

func NewRetrySelector(retries int) channel.VendorSelector {
	return &retrySelector{retries: retries}
}
