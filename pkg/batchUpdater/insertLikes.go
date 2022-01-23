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
	 
	err := DAO.WriteBatch(valuesToWrite, query, "likeId")

	if err != nil {
		log.Println(err.Error())
	}
}

func BatchUpdateReads(readsChannel chan request.Interaction) {
	query := "INSERT INTO readInteraction (userId, contentId) VALUES "
	valuesToWrite := []interface{}{}
	for i:=0;i<len(readsChannel);i++ {
		read := <-readsChannel
		query += "(?, ?), "
		valuesToWrite = append(valuesToWrite, read.UserId, read.ContentId)
	}
	 
	err := DAO.WriteBatch(valuesToWrite, query, "readId")

	if err != nil {
		log.Println(err.Error())
	}
}
