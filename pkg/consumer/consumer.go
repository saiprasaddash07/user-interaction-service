package consumer

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/saiprasaddash07/user-interaction-service/config"
	"github.com/saiprasaddash07/user-interaction-service/helpers/request"
	batchupdater "github.com/saiprasaddash07/user-interaction-service/pkg/batchUpdater"
	"github.com/segmentio/kafka-go"
)

func Consume(ctx context.Context, topic string) {
	likesChannel := make(chan request.Interaction, 1)

	config := config.Get()
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     config.KafkaBrokers,
		Topic:       topic,
		MinBytes:    10e3,
		MaxBytes:    10e6,
		MaxWait:     time.Duration(10) * time.Second,
		StartOffset: kafka.FirstOffset,
		GroupID:     "likes-group",
	})

	for {
		m, err := reader.ReadMessage(context.TODO())
		if err != nil {
			fmt.Errorf("Consumer: %w", err)
		}

		contentId, _ := strconv.ParseInt(string(m.Key), 10, 64)
		userId, _ := strconv.ParseInt(string(m.Value), 10, 64)

		likesChannel <- request.Interaction{
			ContentId: contentId,
			UserId:    userId,
		}

		if len(likesChannel) == 1 {
			batchupdater.BatchUpdateLikes(likesChannel)
		}

		fmt.Printf("Message on %s: %s = %s\n", topic, string(m.Key), string(m.Value))
	}
}
