package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (k *kafka) Init(o ...Options) error {
	if k.connect {
		return fmt.Errorf("kafka is connect")
	}

	for _, op := range o {
		op(&k.option)
	}

	if len(k.option.address) == 0 {
		k.option.address = []string{"localhost:9092"}
	}

	if k.option.log == nil {
		k.option.log = logrus.WithField("kafka", "service")
	}

	if k.option.ctx == nil {
		k.option.ctx = context.Background()
	}

	if k.option.codec == nil {
		k.option.codec = CodeJson{}
	}

	return nil
}

func (k *kafka) IsConnect() (flag bool) {
	k.mutex.RLock()
	defer k.mutex.RUnlock()
	flag = k.connect
	return
}

func (k *kafka) clusterConfig() *sarama.Config {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true          //成功交付的消息将在success channel返回
	cfg.Producer.Return.Errors = true             //失败交付的消息将在failed channel返回
	cfg.Producer.RequiredAcks = sarama.WaitForAll //发送完数据需要leader和follow都确认
	cfg.Producer.Retry.Max = 3

	//if kMQ.opts.TLSConfig != nil {
	//	cfg.Net.TLS.Config = kMQ.opts.TLSConfig
	//	cfg.Net.TLS.Enable = true
	//}

	if !cfg.Version.IsAtLeast(sarama.MaxVersion) {
		cfg.Version = sarama.MaxVersion
	}

	cfg.Consumer.Return.Errors = true
	cfg.Consumer.Offsets.Initial = sarama.OffsetNewest

	return cfg
}

func (k *kafka) Connect() error {
	if k.IsConnect() {
		return fmt.Errorf("kafka is connect")
	}
	k.mutex.Lock()
	defer k.mutex.Unlock()

	if k.connect {
		return nil
	}

	producer, err := sarama.NewSyncProducer(k.address, k.clusterConfig())
	if err != nil {
		return err
	}

	k.pro = producer
	k.connect = true

	return nil
}

func (k *kafka) Disconnect() error {
	if !k.IsConnect() {
		return nil
	}

	k.mutex.Lock()
	defer k.mutex.Unlock()

	if !k.connect {
		return nil
	}

	k.connect = false

	return k.pro.Close()
}

func (k *kafka) String() string {
	return "kafka"
}

func (k *kafka) PushData(topic string, mq *Message) error {
	bys, err := k.codec.Marshal(mq)
	if err != nil {
		return err
	}

	pm := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(bys),
	}

	if key := mq.Getkey(); key != "" {
		pm.Key = sarama.StringEncoder(key)
	}

	_, _, err = k.pro.SendMessage(pm)
	return err
}
func (k *kafka) saramaClusterClient() (sarama.Client, error) {
	return sarama.NewClient(k.address, k.clusterConfig())
}
func (k *kafka) Subscribe(topics string, h Handler, opts ...SubscribeOption) (Subscriber, error) {
	var g sarama.ConsumerGroup
	opt := SubscribeOptions{
		AutoAck: true,
		Queue:   uuid.New().String(),
	}
	for _, o := range opts {
		o(&opt)
	}
	if opt.Context == nil {
		opt.Context = context.Background()
	}

	c, err := k.saramaClusterClient()
	if err != nil {
		return nil, err
	}

	g, err = sarama.NewConsumerGroupFromClient(opt.Queue, c)
	if err != nil {
		c.Close()
		return nil, err
	}

	gc := group{
		handler: h,
		subOpts: opt,
		option:  k.option,
	}

	s := newSubscriber(topics, c, g, gc)
	s.start()

	return s, nil
}

func (CodeJson) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (CodeJson) Unmarshal(v []byte, data interface{}) error {
	return json.Unmarshal(v, data)
}

func (CodeJson) String() string {
	return "json"
}
