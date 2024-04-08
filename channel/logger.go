package channel

import (
	"context"
	"encoding/json"
	"fmt"
	"hermes/channel/message"
)

var LoggerVendorListener = &loggerListener{}

var _ VendorListener = (*loggerListener)(nil)

type loggerListener struct {
}

func (l *loggerListener) OnRequest(ctx context.Context, message *message.Message, vendor Vendor, request any) {
	messageBytes, _ := json.Marshal(message)
	requestBytes, _ := json.Marshal(request)

	formatted := fmt.Sprintf("on request: vendor: [%s: %s], message: %s, request: %s\n", vendor.Id(), vendor.Name(), string(messageBytes), string(requestBytes))
	fmt.Print(formatted)
}

func (l *loggerListener) OnResponse(ctx context.Context, message *message.Message, vendor Vendor, response any, err error) {
	messageBytes, _ := json.Marshal(message)
	responseBytes, _ := json.Marshal(response)

	formatted := fmt.Sprintf("on response: vendor: [%s: %s], message: %s, response: %s, error: %v\n", vendor.Id(), vendor.Name(), string(messageBytes), string(responseBytes), err)
	fmt.Print(formatted)
}
