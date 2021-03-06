package DAO

import (
	"errors"
	"log"

	"github.com/saiprasaddash07/user-interaction-service/constants"
	"github.com/saiprasaddash07/user-interaction-service/services/db"
)

func DoesContentExist(contentId int64) error {
	var count int64
	err := db.GetClient(constants.DB_READER).QueryRow("SELECT COUNT(*) AS count FROM content WHERE contentId=?;", contentId).Scan(&count)

	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New(constants.ERROR_CONTENT_NOT_FOUND)
	}
	log.Println("Hi")

	return nil
}

func DoesUserExist(userId int64) bool {
	var count int64
	err := db.GetClient(constants.DB_READER).QueryRow("SELECT COUNT(*) AS count FROM users WHERE userId=?;", userId).Scan(&count)
	if err != nil {
		return false
	}
	if count == 0 {
		return false
	}
	return true
}

func DoesUserAlreadyLiked(userId int64, contentId int64) bool {
	var count int64
	err := db.GetClient(constants.DB_READER).QueryRow("SELECT COUNT(*) AS count FROM likes WHERE userId=? AND contentId=?;", userId, contentId).Scan(&count)
	if err != nil {
		return false
	}
	if count == 0 {
		return false
	}
	return true
}

func DoesUserAlreadyRead(userId int64, contentId int64) bool {
	var count int64
	err := db.GetClient(constants.DB_READER).QueryRow("SELECT COUNT(*) AS count FROM readInteraction WHERE userId=? AND contentId=?;", userId, contentId).Scan(&count)
	if err != nil {
		return false
	}
	if count == 0 {
		return false
	}
	return true
}

func GetTopContents(limit int) ([]int64, error) {
	rows, err := db.GetClient(constants.DB_READER).Query(`
		SELECT
			contentId
		FROM
			(
				SELECT
					c.contentId AS contentId,
					COUNT(l.likeId) AS likes,
					COUNT(rI.readId) AS readsContent
				FROM
					content c
					JOIN likes l ON l.contentId = c.contentId
					JOIN readInteraction rI ON c.contentId = rI.contentId
					GROUP BY c.contentId
			) AS interaction
		ORDER BY
			likes DESC,
			readsContent DESC
		LIMIT ?;
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contentIds []int64
	for rows.Next() {
		var contentId int64
		err = rows.Scan(&contentId)
		if err != nil {
			return nil, err
		}
		contentIds = append(contentIds, contentId)
	}

	return contentIds, nil
}