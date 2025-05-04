package client

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterClientRoutes(r *gin.Engine) {
	group := r.Group("/clients")
	group.GET("/", GetClients)
	group.POST("/", CreateClient)
	group.PUT("/:email", UpdateClient)
	group.DELETE("/:email", DeleteClient)
}

func GetClients(c *gin.Context) {
	clients, err := LoadClients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, clients)
}

func CreateClient(c *gin.Context) {
	var newClient Client
	if err := c.BindJSON(&newClient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	clients, _ := LoadClients()
	clients = append(clients, newClient)
	SaveClients(clients)
	c.JSON(http.StatusOK, gin.H{"message": "Client created successfully"})
}

func UpdateClient(c *gin.Context) {
	email := c.Param("email")
	var updatedClient Client
	if err := c.BindJSON(&updatedClient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	clients, _ := LoadClients()
	for i, cl := range clients {
		if cl.Email == email {
			clients[i] = updatedClient
			SaveClients(clients)
			c.JSON(http.StatusOK, gin.H{"message": "Client updated"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
}

func DeleteClient(c *gin.Context) {
	email := c.Param("email")
	clients, _ := LoadClients()
	for i, cl := range clients {
		if cl.Email == email {
			clients = append(clients[:i], clients[i+1:]...)
			SaveClients(clients)
			c.JSON(http.StatusOK, gin.H{"message": "Client deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
}
