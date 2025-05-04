package client

type Client struct {
	ClientID      string   `json:"client_id"`
	ClientName    string   `json:"client_name"`
	Email         string   `json:"email"`
	Company       string   `json:"company"`
	ContactNumber string   `json:"contact_number"`
	Address       string   `json:"address"`
	Projects      []string `json:"projects"`
}
