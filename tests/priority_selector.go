package tests

import (
	"context"
	"hermes/channel"
	"hermes/channel/message"
)

var _ channel.VendorSelector = (*prioritySelector)(nil)

type prioritySelector struct {
	priority []string
}

func (w *prioritySelector) Select(ctx context.Context, message *message.Message, vendors channel.Vendors, selected channel.SelectedVendor) error {
	orderedVendors := w.orderByPriority(vendors)
	var err error
	for _, vendor := range orderedVendors {
		if err = selected(vendor); err == nil {
			return nil
		}
	}
	return err
}

func (w *prioritySelector) orderByPriority(vendors channel.Vendors) channel.Vendors {
	vm := vendors.MapById()
	orderedVendors := make(channel.Vendors, 0, len(vendors))
	for _, id := range w.priority {
		if v, ok := vm[id]; ok {
			orderedVendors = append(orderedVendors, v)
		}
	}
	return orderedVendors
}

func newPrioritySelector(priority []string) channel.VendorSelector {
	return &prioritySelector{priority: priority}
}
