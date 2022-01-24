package interactionServices

import (
	"github.com/saiprasaddash07/user-interaction-service/helpers/DAO"
)

func GetTopContents(size int) ([]int64, error) {
	contentIds, err := DAO.GetTopContents(size)

	if err != nil {
		return nil, err
	}

	return contentIds, nil
}
