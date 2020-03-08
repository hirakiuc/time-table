package api

import (
	"github.com/gin-gonic/gin"
)

type ApiServer struct {
}

func NewApiServer() *ApiServer {
	return &ApiServer{}
}

func (s *ApiServer) Start() {
	r := gin.Default()
	r.Get("/schedules", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// listen and serve on 0.0.0.0:8080
	r.Run()
}
