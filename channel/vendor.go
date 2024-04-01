package channel

type Vendor interface {
	Id() string
	Type() string
	Name() string
	AddListener(listener Listener)
	MessageChannel
}

type Vendors []Vendor

type InterceptableChannel interface {
	AddInterceptor(interceptor Interceptor)
}

type Registry interface {
	Register(vendor Vendor)
}

type SwitchableChannel interface {
	Vendor
	InterceptableChannel
	Registry
	AddSelector(selector Selector)
}
