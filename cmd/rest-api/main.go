package main

import (
	helloHandler "boilerplate/cmd/rest-api/handlers/hello"
	"boilerplate/internal/config"
	"boilerplate/pkg/postgresql"
	"boilerplate/pkg/redis"
	"context"

	"github.com/go-redis/cache/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"

	"flag"
	"log"
	"os"

	_ "boilerplate/cmd/rest-api/docs"
)

var (
	db         postgresql.Client
	redisCache *cache.Cache
	appConfig  config.TomlConfig
)

func init() {

	var err error

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

	// БД
	db, err = postgresql.NewClient(context.TODO(), 3, postgresql.ClientConfig{
		Host:     appConfig.Postgres.Host,
		Port:     appConfig.Postgres.Port,
		Database: appConfig.Postgres.Database,
		Username: appConfig.Postgres.User,
		Password: appConfig.Postgres.Password,
	})
	if err != nil {
		log.Panic(err)
	}

	// Кэш
	redisCache = redis.NewClientCache(redis.RedisConfig{
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

	log.Printf("\n\n\nStart app on %s:%s\n", appConfig.HTTP.Host, appConfig.HTTP.Port)
	log.Println("Param: ")
	log.Print(appConfig)

	// Echo instance
	e := echo.New()

	// docs
	e.GET("/v1/swagger/*", echoSwagger.WrapHandler)

	// router
	userGroup := e.Group("/v1")
	prepareUserRoutes(userGroup)

	// Middleware
	e.Use(middleware.Logger())
	e.Logger.SetOutput(log.Writer())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	// Start server
	e.Logger.Fatal(e.Start(appConfig.HTTP.Host + ":" + appConfig.HTTP.Port))
}

func prepareUserRoutes(grp *echo.Group) {
	// @securityDefinitions.apiKey JWTKeyAuth
	// @in header
	// @name Authorization
	/*
		requireClaims := middleware.JWTWithConfig(middleware.JWTConfig{
			Claims:     &utils.JwtCustomClaims{},
			SigningKey: []byte(appConfig.JWT.Secret),
		})
	*/

	// hello
	hHandler := helloHandler.NewHelloHandler(db, redisCache)
	helloGroup := grp.Group("/hello")
	helloGroup.GET("/hello", hHandler.Hello)

}
