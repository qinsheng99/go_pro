package kafka

var Default = New()

func Init(o ...Options) error {
	return Default.Init(o...)
}

func Connect() error {
	return Default.Connect()
}

func Disconnect() error {
	return Default.Disconnect()
}

func PushData(topic string, mq *Message) error {
	return Default.PushData(topic, mq)
}

func Subscribe(topics string, h Handler, opts ...SubscribeOption) (Subscriber, error) {
	return Default.Subscribe(topics, h, opts...)
}
