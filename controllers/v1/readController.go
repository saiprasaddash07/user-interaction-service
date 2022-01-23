package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saiprasaddash07/user-interaction-service/constants"
	"github.com/saiprasaddash07/user-interaction-service/controllers/v1/interactionServices"
	"github.com/saiprasaddash07/user-interaction-service/helpers/request"
	"github.com/saiprasaddash07/user-interaction-service/helpers/response"
	"github.com/saiprasaddash07/user-interaction-service/helpers/util"
)

func ReadHandler(c *gin.Context) {
	readFromContext, ok := c.Get("interactions")
	if !ok {
		c.JSON(http.StatusBadRequest, util.SendErrorResponse(errors.New(constants.INVALID_REQUEST)))
		return
	}
	readObj := readFromContext.(*request.Interaction)

	err := interactionServices.InsertRead(readObj)

	if err != nil {
		c.JSON(http.StatusInternalServerError, util.SendErrorResponse(err))
		return
	}

	res := response.Response{
		Status:  constants.API_SUCCESS_STATUS,
		Message: constants.READ_MESSAGE,
	}
	c.JSON(http.StatusOK, util.StructToJSON(res))
}
