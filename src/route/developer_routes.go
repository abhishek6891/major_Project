package route

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"majorProject/src/user/data/developer"
)

// Utility function to save developers to file
func saveDevelopersToFile(filePath string, developersList []developer.Developer) error {
	data, err := json.MarshalIndent(developersList, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}

func RegisterDeveloper(route *gin.Engine) {
	filePath := "Doc/project/developer.json"

	// GET all developers
	route.GET("/developers", func(c *gin.Context) {
		developersList, err := developer.LoadDevelopersFromFile(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, developersList)
	})

	// POST new developer
	route.POST("/developers", func(c *gin.Context) {
		var newDev developer.Developer
		if err := c.ShouldBindJSON(&newDev); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		developersList, err := developer.LoadDevelopersFromFile(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		developersList = append(developersList, newDev)
		if err := saveDevelopersToFile(filePath, developersList); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Developer added successfully"})
	})

	// PUT update developer by index
	route.PUT("/developers/:index", func(c *gin.Context) {
		index, err := strconv.Atoi(c.Param("index"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index"})
			return
		}

		var updatedDev developer.Developer
		if err := c.ShouldBindJSON(&updatedDev); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		developersList, err := developer.LoadDevelopersFromFile(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if index < 0 || index >= len(developersList) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Developer not found"})
			return
		}

		developersList[index] = updatedDev
		if err := saveDevelopersToFile(filePath, developersList); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Developer updated successfully"})
	})

	// DELETE developer by index
	route.DELETE("/developers/:index", func(c *gin.Context) {
		index, err := strconv.Atoi(c.Param("index"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index"})
			return
		}

		developersList, err := developer.LoadDevelopersFromFile(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if index < 0 || index >= len(developersList) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Developer not found"})
			return
		}

		developersList = append(developersList[:index], developersList[index+1:]...)
		if err := saveDevelopersToFile(filePath, developersList); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Developer deleted successfully"})
	})
}
