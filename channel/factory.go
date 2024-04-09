package channel

import "sync"

var (
	factoryInstance MessageChannelFactory
	factoryOnce     sync.Once
)

// MessageChannelFactory 消息通道工厂，用于获取消息通道
type MessageChannelFactory interface {
	GetChannel(id string) NamedChannel
	Channels() NamedChannels
	Registry
}

// Registry 注册中心，用于注册消息通道
type Registry interface {
	Register(channel NamedChannel)
}

var _ MessageChannelFactory = (*factory)(nil)

type factory struct {
	registry map[string]NamedChannel
}

func (f *factory) Channels() NamedChannels {
	channels := make(NamedChannels, 0, len(f.registry))
	for _, channel := range f.registry {
		channels = append(channels, channel)
	}
	return channels
}

func (f *factory) GetChannel(id string) NamedChannel {
	return f.registry[id]
}

func (f *factory) Register(channel NamedChannel) {
	f.registry[channel.Id()] = channel
}

func NewFactory(channels ...NamedChannel) MessageChannelFactory {
	factoryOnce.Do(func() {
		factoryInstance = &factory{
			registry: make(map[string]NamedChannel, len(channels)),
		}
		for _, channel := range channels {
			factoryInstance.Register(channel)
		}
	})
	return factoryInstance
}
