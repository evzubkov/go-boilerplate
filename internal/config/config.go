package config

import (
	"log"
	"time"

	"github.com/BurntSushi/toml"
)

// Init - Загружает конфигурацию из файла `configFile`. Файл должен находиться в директории `$GOPATH`
func Init(configFile string) (config TomlConfig) {
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal(err)
	}
	log.Println("Init configuration into file:", configFile)
	log.Println("Debug: ", config.Debug)

	return
}

// TomlConfig - Основной конфиг
type TomlConfig struct {
	Debug    bool
	HTTP     ConfigHTTP
	Postgres ConfigPostgres
	Redis    ConfigRedis
	JWT      ConfigJWT
	Logger   ConfigLogger
}

// ConfigHTTP - Конфиг HTTP-сервера
type ConfigHTTP struct {
	Host      string
	Port      string
	EnableApi bool
}

// ConfigPostgres - Конфиг подключения к PostgreSQL
type ConfigPostgres struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

// JWT
type ConfigJWT struct {
	Secret   string
	LifeTime time.Duration
}

// ConfigRedis - Конфиг подключения к Redis
type ConfigRedis struct {
	Host     string
	Port     string
	Username string
	Password string
	Database int
	LifeTime time.Duration // В конфиге указывается в секундах
}

// ConfigLogger - конфигурация логеров
type ConfigLogger struct {
	LogsPath string
	RestAPI  ConfigLoggerRestAPI
}

// ConfigLoggerRestAPI - конфигурация логера сервиса RestAPI
type ConfigLoggerRestAPI struct {
	Compress   bool
	MaxSize    int
	MaxAge     int
	MaxBackups int
}
