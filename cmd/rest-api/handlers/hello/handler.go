package hello

import (
	"boilerplate/pkg/postgresql"

	"github.com/go-redis/cache/v8"
)

type Handler struct {
	Db         postgresql.Client
	RedisCache *cache.Cache
}

func NewHelloHandler(
	db postgresql.Client,
	redisCache *cache.Cache,
) *Handler {
	return &Handler{
		Db:         db,
		RedisCache: redisCache,
	}
}

type ErrAnswer struct {
	Err string `json:"error"`
}
