package main

import (
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
)

var (
	topicName   = "step_data"
	channelName = "falcon"
)

func main() {
	//pro()
	go consu(10)
	go consu(99)
	time.Sleep(10 * 1e9)
}

type mHan struct {
	q  *nsq.Consumer
	id int
}

func (m *mHan) HandleMessage(msg *nsq.Message) error {
	fmt.Println("client:", m.id, "recv-----", string(msg.Body))
	return nil
}

func consu(id int) {
	cfg := nsq.NewConfig()
	//cfg.LocalAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	cfg.DefaultRequeueDelay = 0
	cfg.MaxBackoffDuration = time.Millisecond * 50
	q, _ := nsq.NewConsumer(topicName, channelName, cfg)
	m := &mHan{
		q:  q,
		id: id,
	}
	q.AddHandler(m)
	err := q.ConnectToNSQLookupd("192.168.29.154:4161")
	if err != nil {
		fmt.Println(err)
		return
	}
	//time.Sleep(10 * 1e9)
	//q.Stop()
}

func pro() {
	cfg := nsq.NewConfig()
	w, _ := nsq.NewProducer("192.168.29.154:4150", cfg)
	var err error
	var data string
	for i := 0; i < 10000; i++ {
		data = fmt.Sprintf("hello nsq :%d", i)
		err = w.Publish("step_data", []byte(data))
		if err != nil {
			fmt.Println(err)
			break
		}

	}
	w.Stop()
}
