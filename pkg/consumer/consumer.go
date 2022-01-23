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

func ConsumeLikes(ctx context.Context, topic string) {
	config := config.Get()
	likesChannel := make(chan request.Interaction, config.LikeChannelSize)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     config.KafkaBrokers,
		Topic:       topic,
		MinBytes:    10e3,
		MaxBytes:    10e6,
		MaxWait:     time.Duration(config.WaitTimeForConsumer) * time.Minute,
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

		if len(likesChannel) == config.LikeChannelSize {
			batchupdater.BatchUpdateLikes(likesChannel)
		}

		fmt.Printf("Message on %s: %s = %s\n", topic, string(m.Key), string(m.Value))
	}
}

func ConsumeReads(ctx context.Context, topic string) {
	config := config.Get()
	readsChannel := make(chan request.Interaction, config.ReadChannelSize)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     config.KafkaBrokers,
		Topic:       topic,
		MinBytes:    10e3,
		MaxBytes:    10e6,
		MaxWait:     time.Duration(config.WaitTimeForConsumer) * time.Minute,
		StartOffset: kafka.FirstOffset,
		GroupID:     "reads-group",
	})

	for {
		m, err := reader.ReadMessage(context.TODO())
		if err != nil {
			fmt.Errorf("Consumer: %w", err)
		}

		contentId, _ := strconv.ParseInt(string(m.Key), 10, 64)
		userId, _ := strconv.ParseInt(string(m.Value), 10, 64)

		readsChannel <- request.Interaction{
			ContentId: contentId,
			UserId:    userId,
		}

		if len(readsChannel) == config.ReadChannelSize {
			batchupdater.BatchUpdateReads(readsChannel)
		}

		fmt.Printf("Message on %s: %s = %s\n", topic, string(m.Key), string(m.Value))
	}
}
