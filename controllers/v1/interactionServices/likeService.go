package interactionServices

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/saiprasaddash07/user-interaction-service/constants"
	"github.com/saiprasaddash07/user-interaction-service/helpers/DAO"
	"github.com/saiprasaddash07/user-interaction-service/helpers/request"
	"github.com/saiprasaddash07/user-interaction-service/pkg/producer"
	"github.com/saiprasaddash07/user-interaction-service/services/redis"
)

func InsertLike(likeObj *request.Interaction) error  {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	likeObjKey := fmt.Sprintf("contentId:%d:userId:%d:likes", likeObj.ContentId, likeObj.UserId)
	alreadyLiked, _ := redis.Get(likeObjKey)

	if alreadyLiked == "1" {
		return errors.New(constants.ALREADY_LIKED)
	}

	if ok := DAO.DoesUserAlreadyLiked(likeObj.UserId, likeObj.ContentId); ok {
		return errors.New(constants.ALREADY_LIKED)
	}

	log.Println(likeObjKey)
	go redis.Set(likeObjKey, "1", constants.CACHE_TTL_VERY_LONG)

	go producer.Produce(ctx, "sample-topic", *likeObj)
	return nil
}
