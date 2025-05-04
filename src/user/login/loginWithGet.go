package login

import (
	"github.com/gin-gonic/gin"
	"majorProject/src/user/userLocalDb"
	"net/http"
	"net/url"
)

func LoginRequestWithGet(c *gin.Context) {
	query := c.Request.URL.RawQuery
	params, _ := url.ParseQuery(query)
	email := params.Get("email")
	phone := params.Get("phone")
	password := params.Get("password")
	role := params.Get("role") // New role parameter

	if role == userLocalDb.DEVELOPER {
		// TODO, Do developer login work
		userLocalDb.CURRENT_USER_ROLE = role
	} else if role == userLocalDb.CLIENT {
		// TODO, Do client login work
		userLocalDb.CURRENT_USER_ROLE = role
	} else {
		// TODO, Handle not valid role
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role. Please specify 'developer' or 'client'."})
		//return
	}

	var loginUser userLocalDb.LoginUser

	passwd, isExist := userLocalDb.ValidUsers[email]
	if isExist {
		if passwd == password {
			loginUser = userLocalDb.LoginUser{Email: email, Password: password, Phone: phone, Status: true, Message: "Logged in successfully."}
			userLocalDb.LoggedInUserList = append(userLocalDb.LoggedInUserList, loginUser)
			c.JSON(http.StatusOK, gin.H{"data": loginUser})
		} else {
			loginUser = userLocalDb.LoginUser{Email: email, Status: false, Message: "Password is incorrect."}
			c.JSON(http.StatusOK, gin.H{"data": loginUser})
		}
		return
	} else {
		loginUser = userLocalDb.LoginUser{Email: email, Status: false, Message: "Login is not authorized."}
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"data": loginUser})
	}
}
