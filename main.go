package main

import (
	"context"
	"log"

	"github.com/saiprasaddash07/user-interaction-service/config"
	"github.com/saiprasaddash07/user-interaction-service/pkg/consumer"
	"github.com/saiprasaddash07/user-interaction-service/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	config := config.Get()
	log.Println(config.KafkaBrokers)
	go consumer.Consume(ctx, config.KafkaViewTopic)
	server.Init()
}