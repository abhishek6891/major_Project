package projects

type Project struct {
	ProjectTitle       string   `json:"project_title"`
	ClientName         string   `json:"client_name"`
	ProjectSummary     string   `json:"project_summary"`
	ProjectQuotation   int      `json:"project_quotation_price"`
	CreatedDate        string   `json:"created_date"`
	WireframeGallery   []string `json:"wireframe_gallery"`
	RequiredTechSkills []string `json:"required_tech_skills"`
	Deadline           string   `json:"deadline"`
	PriorityLevel      string   `json:"priority_level"`
	ContactEmail       string   `json:"contact_email"`
	ProjectStatus      string   `json:"project_status"`
	RepositoryLink     string   `json:"repository_link"`
	DesignToolUsed     string   `json:"design_tool_used"`
	Notes              string   `json:"notes"`
}
