package kafka

import (
	"github.com/saiprasaddash07/user-interaction-service/config"
)

type KafkaConfig struct {
	BrokerAddresses []string
	Topic           string
}

func NewKafkaClient(topic string) KafkaConfig {
	config := config.Get()
	kConfig := KafkaConfig{
		BrokerAddresses: config.KafkaBrokers,
		Topic:           topic,
	}
	return kConfig
}
