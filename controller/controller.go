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

func (c *Controller) Save(ctx *gin.Context) {

	ctx.JSON(http.StatusAccepted, "test")
}

func (c *Controller) Get(ctx *gin.Context) {

	ctx.JSON(http.StatusAccepted, "test")
}

func NewController(service service.GeneratorServiceInterface) ControllerInterface {
	return &Controller{
		service: service,
	}
}
