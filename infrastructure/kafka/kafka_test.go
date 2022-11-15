package kafka

import (
	"encoding/json"
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

type GameFields struct {
	UserResult string `json:"user_result"`
	PredPath   string `json:"pred_path"`
	Cls, Pos   int
}

// GameType
// 文本分类 text  图像分类 image  风格迁移style
type GameType struct {
	Type   string `json:"type"`
	UserId int    `json:"user_id"`
}

type Game struct {
	GameType
	GameFields
}

func TestMqGame(t *testing.T) {
	err := Init()
	if err != nil {
		t.Fatal(err)
	}

	err = Connect()
	if err != nil {
		t.Fatal(err)
	}

	data1 := Game{
		GameType: GameType{Type: "text", UserId: 1},
		GameFields: GameFields{
			PredPath: "xihe-obj/competitions/昇思AI挑战赛-多类别图像分类/submit_result/s9qfqri3zpc8j2x7_1/result_example_5120-2022-8-8-15-3-16.txt",
			Cls:      256,
			Pos:      1,
		},
	}
	//
	data2 := Game{
		GameType: GameType{Type: "image", UserId: 2},
		GameFields: GameFields{
			PredPath: "xihe-obj/competitions/昇思AI挑战赛-多类别图像分类/submit_result/s9qfqri3zpc8j2x7_1/result_example_5120-2022-8-8-15-3-16.txt",
			Cls:      256,
			Pos:      1,
		},
	}
	//
	//data3 := Game{
	//	GameType: GameType{Type: "style1", UserId: 3},
	//	GameFields: GameFields{
	//		UserResult: "xihe-obj/competitions/昇思AI挑战赛-艺术家画作风格迁移/submit_result/victor_1/result",
	//	},
	//}
	//
	//data4 := Game{
	//	GameType: GameType{Type: "dd", UserId: 4},
	//	GameFields: GameFields{
	//		UserResult: "xihe-obj/competitions/昇思AI挑战赛-艺术家画作风格迁移/submit_result/victor_1/result",
	//	},
	//}

	bys1, err := json.Marshal(data1)
	if err != nil {
		t.Fatal(err)
	}
	bys2, err := json.Marshal(data2)
	if err != nil {
		t.Fatal(err)
	}
	//bys3, err := json.Marshal(data3)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//bys4, err := json.Marshal(data4)
	//if err != nil {
	//	t.Fatal(err)
	//}

	msg1 := Message{Body: bys1}
	msg2 := Message{Body: bys2}
	//msg3 := Message{Body: bys3}
	//msg4 := Message{Body: bys4}

	err = PushData("game", &msg1)
	if err != nil {
		t.Fatal(err)
	}
	//
	err = PushData("game", &msg2)
	if err != nil {
		t.Fatal(err)
	}
	//
	//err = PushData("xihe-game", &msg3)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//err = PushData("xihe-game", &msg4)
	//if err != nil {
	//	t.Fatal(err)
	//}

	err = Disconnect()
	if err != nil {
		t.Fatal(err)
	}
}
