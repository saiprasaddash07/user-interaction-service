package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saiprasaddash07/user-interaction-service/config"
	"github.com/saiprasaddash07/user-interaction-service/controllers/v1/interactionServices"
	"github.com/saiprasaddash07/user-interaction-service/helpers/util"
)

func TopContentsHandler(c *gin.Context) {
	topContents, err := interactionServices.GetTopContents(config.Get().MaxNumberOfTopContents)

	if err != nil {
		c.JSON(http.StatusInternalServerError, util.SendErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, util.StructToJSON(topContents))
}
