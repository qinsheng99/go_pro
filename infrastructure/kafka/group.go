package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

type group struct {
	option
	handler Handler
	subOpts SubscribeOptions

	notifyReady func()
}

func (g *group) Setup(sarama.ConsumerGroupSession) error {
	g.notifyReady()

	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (g *group) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a group loop of ConsumerGroupClaim's Messages().
func (g *group) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	handle := g.genHanler(session)

	for {
		select {
		case message := <-claim.Messages():
			handle(message)

			if g.subOpts.AutoAck {
				session.MarkMessage(message, "")
			}

		case <-session.Context().Done():
			return nil
		}
	}
}

func (g *group) genHanler(session sarama.ConsumerGroupSession) func(*sarama.ConsumerMessage) {
	handler := g.handler
	if handler == nil {
		handler = func(event MqEvent) error {
			return nil
		}
	}

	log := g.log

	eh := g.ErrorHandler
	if eh == nil {
		eh = func(e MqEvent) error {
			log.Error(e.Error())

			return nil
		}
	}

	unmarshal := g.codec.Unmarshal

	return func(msg *sarama.ConsumerMessage) {
		ke := &event{
			km:   msg,
			m:    new(Message),
			sess: session,
		}

		if err := unmarshal(msg.Value, ke.m); err != nil {
			ke.err = fmt.Errorf("unmarshal msg failed, err: %v", err)
			ke.m.Body = msg.Value

			if err := eh(ke); err != nil {
				log.Error(err)
			}

			return
		}

		if err := handler(ke); err != nil {
			ke.err = fmt.Errorf("handle event, err: %v", err)

			if err := eh(ke); err != nil {
				log.Error(err)
			}
		}
	}
}
