package kafka

type Impl interface {
	Init(o ...Options) error
	IsConnect() (flag bool)
	Disconnect() error
	Connect() error
	String() string
	Subscribe(topics string, h Handler, opts ...SubscribeOption) (Subscriber, error)
	PushData(topic string, mq *Message) error
}

type MqEvent interface {
	Topic() string
	Message() *Message
	Error() error
	Extra() map[string]interface{}
}

type Handler func(mqEvent MqEvent) error
type SubscribeOption func(*SubscribeOptions)
type Options func(*option)

type Subscriber interface {
	Topic() string
	Unsubscribe() error
}
