package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"sync"
)

type kafka struct {
	option
	pro sarama.SyncProducer

	mutex sync.RWMutex

	connect bool
}

type Codeimpl interface {
	Marshal(v any) ([]byte, error)
	Unmarshal(v []byte, data interface{}) error
	String() string
}

type CodeJson struct{}

type option struct {
	address []string

	log *logrus.Entry

	ctx context.Context

	codec Codeimpl

	ErrorHandler Handler
}

func New() Impl {
	k := kafka{
		mutex: sync.RWMutex{},
	}
	return &k
}

func Address(addr ...string) Options {
	return func(o *option) {
		o.address = addr
	}
}

func Log(log *logrus.Entry) Options {
	return func(o *option) {
		o.log = log
	}
}

func Ctx(ctx context.Context) Options {
	return func(o *option) {
		o.ctx = ctx
	}
}

func Codec(code Codeimpl) Options {
	return func(o *option) {
		o.codec = code
	}
}

type Message struct {
	Key  string
	Body []byte
}

func (m *Message) Getkey() string {
	return m.Key
}
