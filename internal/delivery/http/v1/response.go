package v1

import (
	"github.com/gin-gonic/gin"
	"log"
)

type dataResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type response struct {
	Message string `json:"message"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	log.Println(message)
	c.AbortWithStatusJSON(statusCode, response{message})
}

func newDataResponse(c *gin.Context, statusCode int, response dataResponse) {
	log.Println(response)
	c.AbortWithStatusJSON(statusCode, response)
}
