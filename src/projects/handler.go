package projects

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterProjectRoutes(r *gin.Engine) {
	group := r.Group("/projects")
	group.GET("/getProjects", GetProjects)
	group.POST("/createProject", CreateProject)
	group.PUT("/updateProject/:title", UpdateProject)
	group.DELETE("/deleteProject/:title", DeleteProject)
}

func GetProjects(c *gin.Context) {
	projects, err := LoadProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, projects)
}

func CreateProject(c *gin.Context) {
	var newProject Project
	if err := c.BindJSON(&newProject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	projects, _ := LoadProjects()
	projects = append(projects, newProject)
	SaveProjects(projects)
	c.JSON(http.StatusOK, gin.H{"message": "Project created successfully"})
}

func UpdateProject(c *gin.Context) {
	title := c.Param("title")
	var updatedProject Project
	if err := c.BindJSON(&updatedProject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	projects, _ := LoadProjects()
	for i, proj := range projects {
		if proj.ProjectTitle == title {
			projects[i] = updatedProject
			SaveProjects(projects)
			c.JSON(http.StatusOK, gin.H{"message": "Project updated"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
}

func DeleteProject(c *gin.Context) {
	title := c.Param("title")
	projects, _ := LoadProjects()
	for i, proj := range projects {
		if proj.ProjectTitle == title {
			projects = append(projects[:i], projects[i+1:]...)
			SaveProjects(projects)
			c.JSON(http.StatusOK, gin.H{"message": "Project deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
}
