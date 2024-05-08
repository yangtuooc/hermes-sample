package message

import (
	"hermes/app/domain"
	"hermes/channel/message"
)

func buildSimpleMessage(template *domain.Template, simpleMessage *domain.SimpleMessage) *message.Message {
	payload := template.Render(simpleMessage.Args)
	msg := message.New(payload)
	msg.SetChannel(template.Channel)
	msg.SetTo(simpleMessage.To)
	msg.SetRequestId(simpleMessage.RequestId)
	msg.SetHeader("template", template.ShallowCopy())
	msg.SetHeader("args", simpleMessage.Args)
	msg.SetHeader("extra", simpleMessage.Extra)
	return msg
}
