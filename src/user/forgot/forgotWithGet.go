package forgot

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"majorProject/src/user/userLocalDb"
	"net/http"
	"net/url"
	"time"
)

func ForgotPasswordWithGet(c *gin.Context) {
	query := c.Request.URL.RawQuery
	params, _ := url.ParseQuery(query)

	email := params.Get("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	// Check if user exists
	if _, exists := userLocalDb.ValidUsers[email]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "No user found with this email"})
		return
	}

	// Generate reset token
	token := uuid.New().String()
	expiry := time.Now().Add(15 * time.Minute)

	// Save token
	userLocalDb.ResetTokens[token] = email
	userLocalDb.TokenExpiry[token] = expiry

	// Mock "send email"
	resetLink := "http://localhost:8080/reset-password?token=" + token

	c.JSON(http.StatusOK, gin.H{
		"message":     "Password reset link generated.",
		"reset_token": token,
		"expires_in":  "15 minutes",
		"reset_link":  resetLink, // mock link
	})
}
