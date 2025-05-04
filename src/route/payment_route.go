package route

import (
	"github.com/gin-gonic/gin"
	"majorProject/src/payment"
)

func RegisterPaymentRoutes(r *gin.Engine) {
	pay := r.Group("/api/payment")
	{
		pay.POST("/initiate", payment.InitiatePayment)
		pay.GET("/all", payment.GetAllPayments)
		pay.GET("/:id", payment.GetPaymentByID)
	}
}
