package tests

import (
	"context"
	"errors"
	"hermes/channel"
	"hermes/channel/message"
)

var _ channel.VendorSelector = (*specificSelector)(nil)

type specificSelector struct {
	parser *specificSelectorParser
}

func (s *specificSelector) Select(ctx context.Context, message *message.Message, vendors channel.Vendors, selected channel.SelectedVendor) error {
	vendor, err := s.parser.Parse(message, vendors)
	if err != nil {
		return err
	}
	return selected(vendor)
}

type specificSelectorParser struct {
}

func (s *specificSelectorParser) Parse(message *message.Message, vendors channel.Vendors) (channel.Vendor, error) {
	headerValue := message.GetHeader("vendorId")
	if headerValue == nil {
		return nil, errors.New("specific selector: vendorId not found in message headers using key 'vendorId'")
	}
	vendorId, ok := headerValue.(string)
	if !ok {
		return nil, errors.New("specific selector: vendorId is not a string")
	}
	vendor := vendors.Find(vendorId)
	if vendor == nil {
		return nil, errors.New("specific selector: vendor not found in vendors using vendorId")
	}
	return vendor, nil
}

func newSpecificSelector() channel.VendorSelector {
	return &specificSelector{
		parser: &specificSelectorParser{},
	}
}
