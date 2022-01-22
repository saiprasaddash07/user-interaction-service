package batchupdater

import (
	"log"

	"github.com/saiprasaddash07/user-interaction-service/helpers/DAO"
	"github.com/saiprasaddash07/user-interaction-service/helpers/request"
)

func BatchUpdateLikes(likesChannel chan request.Interaction) {
	query := "INSERT INTO likes (userId, contentId) VALUES "
	valuesToWrite := []interface{}{}
	for i:=0;i<len(likesChannel);i++ {
		like := <-likesChannel
		query += "(?, ?), "
		valuesToWrite = append(valuesToWrite, like.UserId, like.ContentId)
	}
	 
	err := DAO.WriteBatch(valuesToWrite, query)

	if err != nil {
		log.Println(err.Error())
	}
}
