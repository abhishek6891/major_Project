package route

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"majorProject/src/projects"
)

// Utility function to write back to file
func saveProjectsToFile(filePath string, projectsList []projects.Project) error {
	data, err := json.MarshalIndent(projectsList, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}

func RegisterProject(route *gin.Engine) {
	filePath := "Doc/project/project.json"

	// GET all projects
	route.GET("/projects", func(c *gin.Context) {
		projectsList, err := projects.LoadProjectsFromFile(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, projectsList)
	})

	// POST: Create a new project
	route.POST("/projects", func(c *gin.Context) {
		var newProject projects.Project
		if err := c.ShouldBindJSON(&newProject); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		projectsList, err := projects.LoadProjectsFromFile(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		projectsList = append(projectsList, newProject)
		if err := saveProjectsToFile(filePath, projectsList); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Project added successfully!"})
	})

	// PUT: Update project by index
	route.PUT("/projects/:index", func(c *gin.Context) {
		index, err := strconv.Atoi(c.Param("index"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index"})
			return
		}

		var updatedProject projects.Project
		if err := c.ShouldBindJSON(&updatedProject); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		projectsList, err := projects.LoadProjectsFromFile(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if index < 0 || index >= len(projectsList) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}

		projectsList[index] = updatedProject
		if err := saveProjectsToFile(filePath, projectsList); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Project updated successfully"})
	})

	// DELETE: Remove project by index
	route.DELETE("/projects/:index", func(c *gin.Context) {
		index, err := strconv.Atoi(c.Param("index"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index"})
			return
		}

		projectsList, err := projects.LoadProjectsFromFile(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if index < 0 || index >= len(projectsList) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}

		projectsList = append(projectsList[:index], projectsList[index+1:]...)
		if err := saveProjectsToFile(filePath, projectsList); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
	})
}
