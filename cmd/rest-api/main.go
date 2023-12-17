package main

import (
	"boilerplate/internal/config"
	"boilerplate/pkg/gin-middleware"
	"boilerplate/pkg/gorm/postgresql"
	"boilerplate/pkg/jwt"
	"boilerplate/pkg/redis"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/cache/v8"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/gorm"
	"log"
	"os"

	_ "boilerplate/cmd/rest-api/docs"
)

var (
	postgres  *gorm.DB
	redisInst *cache.Cache
	appConfig config.TomlConfig
)

func init() {

	var err error

	// Get config path from arg
	configPath := flag.String("config", os.Getenv("PWD")+"/configs/local.toml", "a string")
	flag.Parse()

	// Read config
	appConfig = config.Init(*configPath)

	// Save logs to file and set rotate
	log.SetOutput(&lumberjack.Logger{
		Filename:   appConfig.Logger.LogsPath + "restapi.log",
		MaxSize:    appConfig.Logger.RestAPI.MaxSize,
		MaxBackups: appConfig.Logger.RestAPI.MaxBackups,
		MaxAge:     appConfig.Logger.RestAPI.MaxAge,
		Compress:   appConfig.Logger.RestAPI.Compress,
	})

	if postgres, err = postgresql.NewPostgresClient(postgresql.DbConfig{
		Host:     appConfig.Postgres.Host,
		Port:     appConfig.Postgres.Port,
		User:     appConfig.Postgres.User,
		Password: appConfig.Postgres.Password,
		DbName:   appConfig.Postgres.Database,
		Logger:   nil,
	}); err != nil {
		log.Panic(err)
	}

	// redis
	redisInst = redis.NewRedisClient(redis.ConfigRedis{
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
	authService := jwt.NewJwt(appConfig.JWT.LifeTime, appConfig.JWT.Secret)

	// Add the JWT middleware
	router.Use(middleware.CheckAuth(authService))

	// Define your routes and handlers

	if err := router.Run(":8080"); err != nil {
		log.Panic(err)
	}
}
