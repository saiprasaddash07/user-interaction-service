package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saiprasaddash07/user-interaction-service/config"
	"github.com/saiprasaddash07/user-interaction-service/constants"
	"github.com/saiprasaddash07/user-interaction-service/controllers/v1/interactionServices"
	"github.com/saiprasaddash07/user-interaction-service/helpers/util"
)

func TopContentsHandler(c *gin.Context) {
	header := c.Request.Header["User-Interaction-Header"]
	if len(header) == 0 || header[0] != config.Get().InteractionHeader {
		c.JSON(http.StatusUnauthorized, util.SendErrorResponse(errors.New(constants.INVALID_INTERNAL_REQUEST)))
		return
	}

	topContents, err := interactionServices.GetTopContents(config.Get().MaxNumberOfTopContents)

	if err != nil {
		c.JSON(http.StatusInternalServerError, util.SendErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, util.StructToJSON(topContents))
}
