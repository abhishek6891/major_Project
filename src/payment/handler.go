package payment

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const PaymentDataFile = "./Doc/project/payment_history.json"

// POST /api/payment/initiate
func InitiatePayment(c *gin.Context) {
	var req struct {
		ClientID    string  `json:"client_id"`
		DeveloperID string  `json:"developer_id"`
		Amount      float64 `json:"amount"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	payment := Payment{
		ID:          GeneratePaymentID(),
		ClientID:    req.ClientID,
		DeveloperID: req.DeveloperID,
		Amount:      req.Amount,
		Status:      "pending",
		Timestamp:   time.Now(),
	}

	// Simulate processing
	time.Sleep(2 * time.Second)
	payment.Status = SimulatePaymentProcessing()

	if err := SavePaymentToFile(payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save payment"})
		return
	}

	c.JSON(http.StatusOK, payment)
}

// GET /api/payment/all
func GetAllPayments(c *gin.Context) {
	data, err := os.ReadFile(PaymentDataFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read payment file"})
		return
	}

	var payments []Payment
	if err := json.Unmarshal(data, &payments); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid payment data"})
		return
	}

	c.JSON(http.StatusOK, payments)
}

// GET /api/payment/:id
func GetPaymentByID(c *gin.Context) {
	id := c.Param("id")

	data, err := os.ReadFile(PaymentDataFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read payment file"})
		return
	}

	var payments []Payment
	if err := json.Unmarshal(data, &payments); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid payment data"})
		return
	}

	for _, p := range payments {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
}
