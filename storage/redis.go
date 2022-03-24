package storage

import (
	"time"

	"github.com/go-redis/redis"
)

const CacheDuration = 6 * time.Hour

type RedisService struct {
	redis *redis.Client
}

type RedisInterface interface {
	Save(shortUrl string, originalUrl string)
	Read(shortUrl string) string
}

func (r *RedisService) Save(shortUrl string, originalUrl string) {

}

func (r *RedisService) Read(shortUrl string) string {
	return "test"
}

func NewRedisService(redis *redis.Client) RedisInterface {
	return &RedisService{
		redis: redis,
	}
}
