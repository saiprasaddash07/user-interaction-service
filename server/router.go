package server

import (
	"github.com/gin-gonic/gin"
	"github.com/saiprasaddash07/user-interaction-service/constants"
	v1 "github.com/saiprasaddash07/user-interaction-service/controllers/v1"
	"github.com/saiprasaddash07/user-interaction-service/controllers/v1/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	version1 := router.Group("api/v1")
	{
		interactionGroupV1 := version1.Group("interaction")
		{
			interactionGroupV1.POST("/like", middlewares.GetRequestBody(constants.LIKE_POST_REQUIRED_FIELDS, constants.LIKE_POST_OPTIONAL_FIELDS), v1.LikeHandler)
		}
	}

	return router
}
