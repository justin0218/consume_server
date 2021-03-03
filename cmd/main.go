package main

import (
	"consume_server/internal/services"
	"consume_server/store"
	"fmt"
	"time"
)

func init() {
	log := new(store.Log)
	log.Get().Debug("server started at %v", time.Now())
	fmt.Printf("server started at %v", time.Now())
}

func main() {
	mq := new(store.Rabbitmq)
	mq.Get()
	wxPay := new(services.WxPay)
	err := wxPay.ConsumeWxPayNotice()
	if err != nil {
		panic(err)
	}
}
