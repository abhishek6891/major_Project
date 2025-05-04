package payment

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func SavePaymentToFile(p Payment) error {
	var payments []Payment

	// Load existing payments
	file, _ := os.ReadFile(PaymentDataFile)
	if len(file) > 0 {
		json.Unmarshal(file, &payments)
	}

	// Append new one
	payments = append(payments, p)

	// Save back
	updated, err := json.MarshalIndent(payments, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(PaymentDataFile, updated, 0644)
}

func GeneratePaymentID() string {
	return fmt.Sprintf("pay_%d", time.Now().UnixNano())
}

func SimulatePaymentProcessing() string {
	statuses := []string{"success", "failed"}
	rand.Seed(time.Now().UnixNano())
	return statuses[rand.Intn(len(statuses))]
}
