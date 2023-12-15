package redis

import (
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type ConfigRedis struct {
	Host     string
	Port     string
	Database int
	Username string
	Password string
}

func NewRedisClient(config ConfigRedis) *cache.Cache {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs:    map[string]string{"redis": config.Host + ":" + config.Port},
		DB:       config.Database,
		Username: config.Username,
		Password: config.Password,
	})

	redisCache := cache.New(&cache.Options{Redis: ring, LocalCache: cache.NewTinyLFU(1000, time.Minute)})

	return redisCache
}
