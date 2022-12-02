package kafka

import (
	"context"
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

type subscriber struct {
	cli sarama.Client
	cg  sarama.ConsumerGroup

	t  string
	gc group

	once  sync.Once
	ready chan struct{}
	stop  chan struct{}
	done  chan struct{}
}

type SubscribeOptions struct {
	// AutoAck defaults to true. When a handler returns
	// with a nil error the message is receipt already.
	AutoAck bool
	// Subscribers with the same queue name
	// will create a shared subscription where each
	// receives a subset of messages.
	Queue string

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

func newSubscriber(
	topic string,
	cli sarama.Client, cg sarama.ConsumerGroup,
	gc group,

) (s *subscriber) {
	s = &subscriber{
		t:   topic,
		cli: cli,
		cg:  cg,
		gc:  gc,

		ready: make(chan struct{}),
		stop:  make(chan struct{}),
		done:  make(chan struct{}),
	}

	s.gc.notifyReady = s.notifyReady

	return
}

func (s *subscriber) Topic() string {
	return s.t
}

func (s *subscriber) Unsubscribe() error {
	var mErr []error

	s.once.Do(func() {
		close(s.stop)

		// wait
		<-s.done

		mErr = append(mErr, s.cg.Close())
		mErr = append(mErr, s.cli.Close())
	})

	return fmt.Errorf("%v", mErr)
}

func (s *subscriber) start() {
	log := s.gc.log
	ctx := s.gc.subOpts.Context
	topic := []string{s.t}

	go func() {
		defer close(s.done)

		for {
			select {
			case err := <-s.cg.Errors():
				if err != nil {
					log.Errorf("consumer error: %v", err)
				}

			case <-s.stop:
				log.Errorf("consumer stopped")
				return

			default:
				err := s.cg.Consume(ctx, topic, &s.gc)
				switch err {
				case nil:
					continue
				case sarama.ErrClosedConsumerGroup:
					return
				default:
					log.Error(err)
				}
			}
		}
	}()

	<-s.ready
}

func (s *subscriber) notifyReady() {
	close(s.ready)
}
