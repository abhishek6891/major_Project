package forgot

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"majorProject/src/user/userLocalDb"
	"net/http"
	"net/url"
	"time"
)

func ResetPasswordWithGet(c *gin.Context) {
	query := c.Request.URL.RawQuery
	params, _ := url.ParseQuery(query)

	token := params.Get("token")
	newPassword := params.Get("newPassword")

	if token == "" || newPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token and new password are required"})
		return
	}

	email, tokenExists := userLocalDb.ResetTokens[token]
	expiry, timeExists := userLocalDb.TokenExpiry[token]

	if !tokenExists || !timeExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired token"})
		return
	}

	if time.Now().After(expiry) {
		delete(userLocalDb.ResetTokens, token)
		delete(userLocalDb.TokenExpiry, token)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token has expired"})
		return
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Update password
	userLocalDb.ValidUsers[email] = string(hashedPassword)

	// Invalidate token
	delete(userLocalDb.ResetTokens, token)
	delete(userLocalDb.TokenExpiry, token)

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successful"})
}
