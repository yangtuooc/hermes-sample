package channel

type MessageChannelFactory interface {
	GetChannel(id string) MessageChannel
	Channels() MessageChannels
	Registry
}

type Registry interface {
	Register(channel NamedChannel)
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

func (f *factory) Register(channel NamedChannel) {
	f.registry[channel.Id()] = channel
}

func NewFactory(channels ...NamedChannel) MessageChannelFactory {
	f := &factory{
		registry: make(map[string]MessageChannel),
	}
	for _, channel := range channels {
		f.Register(channel)
	}
	return f
}
