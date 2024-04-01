package message

const (
	channelIdKey = "channelId"
	requestIdKey = "requestId"
)

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

func (m *Message) SetChannelId(channelId string) {
	m.SetHeader(channelIdKey, channelId)
}

func (m *Message) GetChannelId() string {
	id, ok := m.GetHeader(channelIdKey).(string)
	if !ok {
		return ""
	}
	return id
}

func (m *Message) SetRequestId(requestId string) {
	m.SetHeader(requestIdKey, requestId)
}

func (m *Message) GetRequestId() string {
	requestId, ok := m.GetHeader(requestIdKey).(string)
	if !ok {
		return ""
	}
	return requestId
}

func New(payload any) *Message {
	return &Message{
		Headers: newHeaders(),
		Payload: payload,
	}
}
