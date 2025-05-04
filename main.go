package main

import (
	"github.com/gin-gonic/gin"
	routes "majorProject/src/route"
	"majorProject/src/user/forgot"
	"majorProject/src/user/login"
	"majorProject/src/user/signup"
	"majorProject/src/video"
)

func main() {
	route := gin.Default()

	route.GET("/login", login.LoginRequestWithGet)
	route.GET("/signup", signup.SignUpRequestWithGet)
	route.GET("/forgotpassword", forgot.ForgotPasswordWithGet)
	route.GET("/resetpassword", forgot.ResetPasswordWithGet)

	routes.RegisterProject(route)
	routes.RegisterClient(route)
	routes.RegisterDeveloper(route)

	routes.RegisterChat(route)
	video.RegisterVideoCallRoutes(route)
	routes.RegisterSignalRoutes(route)
	routes.RegisterPaymentRoutes(route)
	/*
		routes.RegisterProjectRoutes(route)
		routes.RegisterClientRoutes(route)
		routes.RegisterDeveloperRoutes(route)
	*/
	route.Run() // listen and serve on 0.0.0.0:8080
}
