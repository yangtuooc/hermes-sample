package channel

type MessageChannelFactory interface {
	GetChannel(id string) MessageChannel
	Channels() MessageChannels
	Registry
}

type Registry interface {
	Register(id string, channel MessageChannel)
}

var _ MessageChannelFactory = (*factory)(nil)

type factory struct {
	registry map[string]MessageChannel
}

func (f *factory) Channels() MessageChannels {
	channels := make(MessageChannels, 0, len(f.registry))
	for _, channel := range f.registry {
		channels = append(channels, channel)
	}
	return channels
}

func (f *factory) GetChannel(id string) MessageChannel {
	return f.registry[id]
}

func (f *factory) Register(id string, channel MessageChannel) {
	f.registry[id] = channel
}

func NewFactory(pairs ...map[string]MessageChannel) MessageChannelFactory {
	f := &factory{
		registry: make(map[string]MessageChannel),
	}
	for _, pair := range pairs {
		for id, channel := range pair {
			f.Register(id, channel)
		}
	}
	return f
}
