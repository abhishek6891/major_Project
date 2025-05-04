package route

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"majorProject/src/user/data/client"
)

// Utility function to save clients to file
func saveClientsToFile(filePath string, clientsList []client.Client) error {
	data, err := json.MarshalIndent(clientsList, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}

func RegisterClient(route *gin.Engine) {
	filePath := "Doc/project/client.json"

	// GET all clients
	route.GET("/clients",
		func(c *gin.Context) {
			clientsList, err := client.LoadClientsFromFile(filePath)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, clientsList)
		})

	// POST new client
	route.POST("/clients",
		func(c *gin.Context) {
			var newClient client.Client
			if err := c.ShouldBindJSON(&newClient); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			clientsList, err := client.LoadClientsFromFile(filePath)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			clientsList = append(clientsList, newClient)
			if err := saveClientsToFile(filePath, clientsList); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusCreated, gin.H{"message": "Client added successfully"})
		})

	// PUT update client by index
	route.PUT("/clients/:index", func(c *gin.Context) {
		index, err := strconv.Atoi(c.Param("index"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index"})
			return
		}

		var updatedClient client.Client
		if err := c.ShouldBindJSON(&updatedClient); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		clientsList, err := client.LoadClientsFromFile(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if index < 0 || index >= len(clientsList) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
			return
		}

		clientsList[index] = updatedClient
		if err := saveClientsToFile(filePath, clientsList); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Client updated successfully"})
	})

	// DELETE client by index
	route.DELETE("/clients/:index", func(c *gin.Context) {
		index, err := strconv.Atoi(c.Param("index"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index"})
			return
		}

		clientsList, err := client.LoadClientsFromFile(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if index < 0 || index >= len(clientsList) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
			return
		}

		clientsList = append(clientsList[:index], clientsList[index+1:]...)
		if err := saveClientsToFile(filePath, clientsList); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Client deleted successfully"})
	})
}
