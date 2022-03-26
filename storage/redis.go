package storage

import (
	"fmt"
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
	err := r.redis.Set(shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

func (r *RedisService) Read(shortUrl string) string {
	result, err := r.redis.Get(shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}

func NewRedisService(redis *redis.Client) RedisInterface {
	return &RedisService{
		redis: redis,
	}
}
