package producer

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/saiprasaddash07/user-interaction-service/config"
	"github.com/saiprasaddash07/user-interaction-service/helpers/request"
	"github.com/segmentio/kafka-go"
)

func Produce(ctx context.Context, topic string, data request.Interaction) {
	config := config.Get()
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      config.KafkaBrokers,
		Topic:        topic,
		BatchSize:    1,
		BatchBytes:   10e4,
		BatchTimeout: time.Duration(time.Minute * 1),
	})

	err := writer.WriteMessages(context.TODO(), kafka.Message{
		Key:   []byte(strconv.FormatInt(data.ContentId, 10)),
		Value: []byte(strconv.FormatInt(data.UserId, 10)),
	})
	if err != nil {
		fmt.Errorf("Produce: %w", err)
	}

	log.Println("Producing")
}
