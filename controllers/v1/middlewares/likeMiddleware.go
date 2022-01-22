package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saiprasaddash07/user-interaction-service/constants"
	"github.com/saiprasaddash07/user-interaction-service/controllers/v1/utils"
)

func GetRequestBody(interactionRequiredFields []string, interactionOptionalFields []string) gin.HandlerFunc {
	return func(context *gin.Context) {
		var requestObj interface{}

		if err := context.ShouldBind(&requestObj); err != nil {
			context.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"status":  constants.API_FAILED_STATUS,
				"message": constants.INVALID_REQUEST,
			})
			return
		}

		interactionJSON := requestObj.(map[string]interface{})

		interactions, ok := utils.ValidateAndParseInteractionFields(interactionJSON, interactionRequiredFields, interactionOptionalFields)
		if !ok {
			context.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"status":  constants.API_FAILED_STATUS,
				"message": constants.INVALID_REQUEST,
			})
			return
		}

		if err := utils.ValidateInteractionDetails(interactions); err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  constants.API_FAILED_STATUS,
				"message": err.Error(),
			})
			return
		}

		context.Set("interactions", interactions)
		context.Next()
	}
}
