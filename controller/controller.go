package controller

import (
	"net/http"
	"urlShortener/service"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service service.GeneratorServiceInterface
}

type ControllerInterface interface {
	Save(ctx *gin.Context)
	Get(ctx *gin.Context)
}

type Request struct {
	LongUrl string `json:"long_url" binding:"required"`
}

func (c *Controller) Save(ctx *gin.Context) {
	var creationRequest Request
	if err := ctx.ShouldBindJSON(&creationRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := c.service.GenerateShortLink(creationRequest.LongUrl)

	host := "http://localhost:3004/"
	ctx.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func (c *Controller) Get(ctx *gin.Context) {
	shortUrl := ctx.Param("shortUrl")
	initialUrl := c.service.Get(shortUrl)
	ctx.Redirect(302, initialUrl)
}

func NewController(service service.GeneratorServiceInterface) ControllerInterface {
	return &Controller{
		service: service,
	}
}
