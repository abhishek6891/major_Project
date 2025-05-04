package sample

import (
	"fmt"
	"os"
)

func main() {
	// Define your JSON string
	jsonString := `{
		"project_title": "E-Commerce Website",
		"client_name": "Alice Johnson",
		"project_summary": "An online store to sell handmade crafts with payment integration and order tracking.",
		"project_quotation_price": 4500,
		"created_date": "2025-03-21",
		"wireframe_gallery": [
			"https://example.com/wireframes/ecommerce1.png",
			"https://example.com/wireframes/ecommerce2.png"
		],
		"required_tech_skills": ["React", "Node.js", "MongoDB", "Stripe API"],
		"deadline": "2025-05-01",
		"priority_level": "High",
		"contact_email": "alice.j@example.com",
		"project_status": "In Progress",
		"repository_link": "https://github.com/client/ecommerce-store",
		"design_tool_used": "Figma",
		"notes": "Focus on mobile responsiveness and fast load times."
	}`

	// Convert the JSON string to a byte slice
	data := []byte(jsonString)

	// Write the byte slice to a file
	err := os.WriteFile("project.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	fmt.Println("JSON data successfully written to project.json")
}
