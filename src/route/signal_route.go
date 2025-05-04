package route

import (
	"github.com/gin-gonic/gin"

	"majorProject/src/video"
)

func RegisterSignalRoutes(router *gin.Engine) {
	router.GET("/ws/signal", func(c *gin.Context) {
		video.SignalingHandler(c.Writer, c.Request)
	})
}
