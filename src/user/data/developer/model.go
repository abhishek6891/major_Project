package developer

type Developer struct {
	DeveloperID  string   `json:"developer_id"`
	Name         string   `json:"name" validate:"required"`
	Email        string   `json:"email" validate:"required,email"`
	Skills       []string `json:"skills" validate:"required,min=1,dive,required"`
	Experience   int      `json:"experience" validate:"required,min=0"`
	LinkedIn     string   `json:"linkedin"`
	Available    string   `json:"available"`
	ContactEmail string   `json:"contact_email" validate:"required,email"`
}
