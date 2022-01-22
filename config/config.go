package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName              string
	AppEnv               string
	DBUserName           string
	DBHostReader         string
	DBHostWriter         string
	DBPort               string
	DBPassword           string
	DBName               string
	DBMaxIdleConnections int
	DBMaxOpenConnections int
	ServerPort           string
	RedisAddress         string
	RedisPassword        string
	CacheEnabled         bool
	KafkaBrokers         []string
	KafkaViewTopic       string
}

var config Config

// init() Should run at the very beginning, before any other package or code.
func init() {
	appEnv := os.Getenv("APP_ENV")
	if len(appEnv) == 0 {
		appEnv = "dev"
	}

	configFilePath := "./config/.env.dev"

	switch appEnv {
	case "production":
		configFilePath = "./config/.env.prod"
		break
	case "stage":
		configFilePath = "./config/.env.stage"
		break
	}
	log.Println("reading env from: " + configFilePath)

	e := godotenv.Load(configFilePath)

	if e != nil {
		log.Println("error loading .env: ", e)
		panic(e.Error())
	}

	config.AppName = os.Getenv("SERVICE_NAME")
	config.AppEnv = appEnv
	config.DBUserName = os.Getenv("DB_USERNAME")
	config.DBHostReader = os.Getenv("DB_HOST_READER")
	config.DBHostWriter = os.Getenv("DB_HOST_WRITER")
	config.DBPort = os.Getenv("DB_PORT")
	config.DBPassword = os.Getenv("DB_PASSWORD")
	config.DBName = os.Getenv("DB_NAME")
	config.ServerPort = os.Getenv("SERVER_PORT")
	config.DBMaxIdleConnections, _ = strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONENCTION"))
	config.DBMaxOpenConnections, _ = strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS"))
	config.RedisAddress = os.Getenv("REDIS_ADDRESS")
	config.RedisPassword = os.Getenv("REDIS_PASSWORD")
	config.CacheEnabled, _ = strconv.ParseBool(os.Getenv("CACHE_ENABLED"))
	config.KafkaBrokers = strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
	config.KafkaViewTopic = os.Getenv("KAFKA_VIEW_TOPIC")
}

func Get() Config {
	return config
}

func IsProduction() bool {
	return config.AppEnv == "production"
}
