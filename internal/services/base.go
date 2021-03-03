package services

import "consume_server/store"

type baseService struct {
	Config   store.Config
	Rabbitmq store.Rabbitmq
}
