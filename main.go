package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

	server := gin.Default()
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusCreated, "Hello world")
	})

	server.Run(viper.GetString("server.address"))
}
