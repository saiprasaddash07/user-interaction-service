package server

import (
	"github.com/saiprasaddash07/user-interaction-service/config"
	"github.com/saiprasaddash07/user-interaction-service/services/db"
	"github.com/saiprasaddash07/user-interaction-service/services/logger"
	"github.com/saiprasaddash07/user-interaction-service/services/redis"
)

func Init() {
	config := config.Get()
	logger.InitLogger()
	db.Init()
	redis.Init()
	r := NewRouter()
	r.Run(":" + config.ServerPort)
}
