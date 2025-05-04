package route

import (
	"github.com/gin-gonic/gin"
	"majorProject/src/projects"
	"majorProject/src/user/data/client"
	"majorProject/src/user/data/developer"

	/*"majorProject/src/projects"
	"majorProject/src/user/data/client"
	"majorProject/src/user/data/developer"*/
)

func RegisterProjectRoutes(r *gin.Engine) {
	projects.RegisterProjectRoutes(r)
}

func RegisterClientRoutes(r *gin.Engine) {
	client.RegisterClientRoutes(r)
}

func RegisterDeveloperRoutes(r *gin.Engine) {
	developer.RegisterDeveloperRoutes(r)
}
