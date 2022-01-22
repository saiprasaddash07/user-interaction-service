package DAO

import (
	"log"

	"github.com/saiprasaddash07/user-interaction-service/constants"
	"github.com/saiprasaddash07/user-interaction-service/services/db"
)

func WriteBatch(valuesToWrite []interface{}, writeSql string) error {
	writeSql = writeSql[0 : len(writeSql)-2]
	stmt, err := db.GetClient(constants.DB_WRITER).Prepare(writeSql)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	writeSql += " ON DUPLICATE KEY UPDATE likeId = likeId"
	_, err = stmt.Exec(valuesToWrite...)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
