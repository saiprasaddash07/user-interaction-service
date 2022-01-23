package utils

import (
	"errors"
	"fmt"

	"github.com/saiprasaddash07/user-interaction-service/constants"
	"github.com/saiprasaddash07/user-interaction-service/helpers/DAO"
	"github.com/saiprasaddash07/user-interaction-service/helpers/request"
	"github.com/saiprasaddash07/user-interaction-service/helpers/util"
	"github.com/saiprasaddash07/user-interaction-service/services/redis"
)

func ValidateAndParseInteractionFields(interactionJSON map[string]interface{}, requiredFields []string, optionalFields []string) (*request.Interaction, bool) {
	lengthDiffRequiredFieldsAndinteractionJSON := len(interactionJSON) - len(requiredFields)
	if lengthDiffRequiredFieldsAndinteractionJSON < 0 || len(optionalFields) < lengthDiffRequiredFieldsAndinteractionJSON {
		return nil, false
	}

	countOfReqFields := len(requiredFields)
	var likeReq request.Interaction
	for k, v := range interactionJSON {
		if util.Contains(requiredFields, k) {
			countOfReqFields--
		} else if !util.Contains(optionalFields, k) {
			return nil, false
		}

		valueType := fmt.Sprintf("%T", v)
		switch k {
		case "userId":
			if valueType == "float64" && util.IsInteger(v.(float64)) {
				likeReq.UserId = int64(v.(float64))
			} else {
				return &likeReq, false
			}
		case "contentId":
			if valueType == "float64" && util.IsInteger(v.(float64)) {
				likeReq.ContentId = int64(v.(float64))
			} else {
				return &likeReq, false
			}
		default:
			return nil, false
		}
	}
	if countOfReqFields == 0 {
		return &likeReq, true
	}
	return nil, false
}

func ValidateInteractionDetails(interaction *request.Interaction) error {
	if interaction.UserId < 0 || interaction.ContentId < 0 {
		return errors.New(constants.ERROR_INVALID_INTERACTION_DETAILS)
	}

	// contentKey := fmt.Sprintf("content:%d:exist", interaction.ContentId)
	// contentExist, _ := redis.Get(contentKey)
	// if contentExist != "1" {
	// 	if err := DAO.DoesContentExist(interaction.ContentId); err != nil {
	// 		return err
	// 	}
	// 	go redis.Set(contentKey, "1", constants.CACHE_TTL_VERY_LONG)
	// }
	
	userKey := fmt.Sprintf("user:%d:exist", interaction.UserId)
	userExist, _ := redis.Get(userKey)
	if userExist != "1" {
		if ok := DAO.DoesUserExist(interaction.UserId); !ok {
			return errors.New(constants.ERROR_NO_USER_EXIST)
		}
		go redis.Set(userKey, "1", constants.CACHE_TTL_VERY_LONG)
	}

	return nil
}
