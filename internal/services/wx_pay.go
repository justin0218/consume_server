package services

import (
	"consume_server/internal/models/wx"
	"encoding/json"
	"fmt"
)

type WxPay struct {
	baseService
}

func (s *WxPay) ConsumeWxPayNotice() (err error) {
	mq := s.Rabbitmq.Get()
	mqChan, e := mq.Channel()
	if e != nil {
		err = e
		return
	}
	defer mqChan.Close()
	q, e := mqChan.QueueDeclare(
		"wxPayNotice", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if e != nil {
		err = e
		return
	}
	msgs, err := mqChan.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			wxPayReq := new(wx.WxpayReq)
			err := json.Unmarshal(d.Body, wxPayReq)
			if err != nil {
				continue
			}
			fmt.Printf("%+v", wxPayReq)
			_ = d.Ack(true)
		}
	}()
	<-forever
	return
}
