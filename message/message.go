package message

type Message struct {
	Headers headers `json:"headers"`
	Payload any     `json:"payload"`
}

func (m *Message) SetHeader(key string, value any) {
	m.Headers[key] = value
}

func (m *Message) GetHeader(key string) any {
	return m.Headers[key]
}

func (m *Message) SetRequestId(requestId string) {
	m.SetHeader("requestId", requestId)
}

func (m *Message) GetRequestId() string {
	return m.GetHeader("requestId").(string)
}

func New(payload any) *Message {
	return &Message{
		Headers: newHeaders(),
		Payload: payload,
	}
}
