package interactionServices

import (
	"context"
	"errors"
	"fmt"

	"github.com/saiprasaddash07/user-interaction-service/constants"
	"github.com/saiprasaddash07/user-interaction-service/helpers/DAO"
	"github.com/saiprasaddash07/user-interaction-service/helpers/request"
	"github.com/saiprasaddash07/user-interaction-service/pkg/producer"
	"github.com/saiprasaddash07/user-interaction-service/services/redis"
)

func InsertRead(readObj *request.Interaction) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	readObjKey := fmt.Sprintf("contentId:%d:userId:%d:reads", readObj.ContentId, readObj.UserId)
	alreadyRead, _ := redis.Get(readObjKey)

	if alreadyRead == "1" {
		return errors.New(constants.ALREADY_READ)
	}

	if ok := DAO.DoesUserAlreadyRead(readObj.UserId, readObj.ContentId); ok {
		return errors.New(constants.ALREADY_READ)
	}

	go redis.Set(readObjKey, "1", constants.CACHE_TTL_VERY_LONG)

	go producer.ProduceReads(ctx, "read-topic", *readObj)
	return nil
}
