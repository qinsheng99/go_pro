package kafka

import (
	"testing"
)

func TestInit(t *testing.T) {
	err := Init()
	if err != nil {
		t.Fatal(err)
	}

	err = Connect()
	if err != nil {
		t.Fatal(err)
	}
}

func TestBroker(t *testing.T) {
	if err := Init(); err != nil {
		t.Fatalf("mq init error: %v", err)
	}

	if err := Connect(); err != nil {
		t.Fatalf("mq connect error: %v", err)
	}

	msg := Message{
		Body: []byte(`{"message":"broker_test"}`),
	}
	done := make(chan bool)

	sub, err := Subscribe("mq-test", func(event MqEvent) error {
		m := event.Message()
		if string(m.Body) != string(msg.Body) {
			t.Fatalf("Unexpected msg %s, expected %s", string(m.Body), string(msg.Body))
		}

		t.Logf("body: %s , extra: %v", string(m.Body), event.Extra())

		close(done)

		return nil
	})
	if err != nil {
		t.Fatalf("Unexpected subscribe error: %v", err)
	}

	if err := PushData("mq-test", &msg); err != nil {
		t.Fatalf("Unexpected publish error: %v", err)
	}

	<-done
	_ = sub.Unsubscribe()

	if err := Disconnect(); err != nil {
		t.Fatalf("Unexpected disconnect error: %v", err)
	}
}
