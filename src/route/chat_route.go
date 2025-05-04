package route

import (
	"github.com/gin-gonic/gin"
	"majorProject/src/chat"
)

func RegisterChat(r *gin.Engine) {
	//func RegisterChatRoutes(router *gin.Engine) {
	//	router.GET("/ws/chat", ChatHandler)
	r.GET("/ws/chat", func(c *gin.Context) {
		chat.ChatHandler(c.Writer, c.Request)
	})
}
