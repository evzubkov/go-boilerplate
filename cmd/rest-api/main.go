package main

import (
	"boilerplate/internal/config"
	"boilerplate/pkg/gin-middleware"
	"boilerplate/pkg/jwt"
	"boilerplate/pkg/redis"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/cache/v8"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"

	"flag"
	"log"
	"os"

	_ "boilerplate/cmd/rest-api/docs"
)

var (
	redisCache *cache.Cache
	appConfig  config.TomlConfig
)

func init() {

	// Получаем путь до конфига из командной строки
	configPath := flag.String("config", os.Getenv("PWD")+"/configs/local.toml", "a string")
	flag.Parse()

	// Загружаем конфигурацию
	appConfig = config.Init(*configPath)

	// Сохраняем логи в файл
	log.SetOutput(&lumberjack.Logger{
		Filename:   appConfig.Logger.LogsPath + "restapi.log",
		MaxSize:    appConfig.Logger.RestAPI.MaxSize,
		MaxBackups: appConfig.Logger.RestAPI.MaxBackups,
		MaxAge:     appConfig.Logger.RestAPI.MaxAge,
		Compress:   appConfig.Logger.RestAPI.Compress,
	})

	// Кэш
	redisCache = redis.NewRedisClient(redis.ConfigRedis{
		Host:     appConfig.Redis.Host,
		Port:     appConfig.Redis.Port,
		Database: appConfig.Redis.Database,
		Username: appConfig.Redis.Username,
		Password: appConfig.Redis.Password,
	})
}

// @title boilerplate
// @version 1.0
// @description API
// @contact.name Ev. Zubkov
// @contact.url
// @contact.email evzubkov@inbox.ru
// @license.name Commercial
// @BasePath /v1
func main() {

	router := gin.Default()

	// Create a new JWT instance
	authService := jwt.NewJwt(time.Minute, "secret")

	// Add the JWT middleware
	router.Use(middleware.CheckAuth(authService))

	// Define your routes and handlers

	if err := router.Run(":8080"); err != nil {
		log.Panic(err)
	}
}
