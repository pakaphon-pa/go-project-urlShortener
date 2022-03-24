package main

import (
	"log"
	"net/http"
	"urlShortener/controller"
	"urlShortener/service"
	"urlShortener/storage"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	viper.AutomaticEnv() // for global env

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}

	dsn := viper.GetString(`redis.REDIS_DSN`)

	if len(dsn) == 0 {
		dsn = "redis:6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	_, errRedis := client.Ping().Result()

	if errRedis != nil {
		panic(errRedis)
	}

	storage := storage.NewRedisService(client)
	service := service.NewGeneratorService(storage)
	controller := controller.NewController(service)

	server := gin.Default()
	server.POST("/", controller.Save)
	server.GET("/", controller.Get)
	server.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusCreated, "Hello world")
	})

	server.Run(viper.GetString("server.address"))
}
