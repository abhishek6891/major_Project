package video

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// RegisterVideoCallRoutes sets up the video call scheduling endpoint.
func RegisterVideoCallRoutes(router *gin.Engine) {
	router.POST("/schedule-call", func(c *gin.Context) {
		var newCall VideoCall // VideoCall is from model.go
		if err := c.ShouldBindJSON(&newCall); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newCall.StartTime = time.Now()
		newCall.RoomID = newCall.CallerID + "_" + newCall.CalleeID // Create a unique room ID

		if err := SaveCallSession(newCall, "Doc/project/video_calls.json"); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save call"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Call scheduled",
			"roomId":  newCall.RoomID,
		})
	})
}
