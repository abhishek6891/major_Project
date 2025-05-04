package signup

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"majorProject/src/user/userLocalDb"
	"net/http"
	"net/url"
)

func SignUpRequestWithGet(c *gin.Context) {
	query := c.Request.URL.RawQuery
	params, _ := url.ParseQuery(query)

	email := params.Get("email")
	password := params.Get("password")
	phone := params.Get("phone")
	role := params.Get("role")

	// Validating the role
	if role != userLocalDb.DEVELOPER && role != userLocalDb.CLIENT {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role. Must be 'developer' or 'client'."})
		return
	}

	// Checking if the email is already exists or not
	if _, exists := userLocalDb.ValidUsers[email]; exists {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists with this email"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
		return
	}

	// Storing user data
	userLocalDb.ValidUsers[email] = string(hashedPassword)
	userLocalDb.RegisteredUser[email] = phone

	c.JSON(http.StatusCreated, gin.H{
		"message": "User signed up successfully",
		"email":   email,
		"role":    role,
	})
}
