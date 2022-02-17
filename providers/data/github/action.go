package action

type Workflows struct {
	TotalCount int `json:"total_count"`
	Workflows  *[]Workflow
}

type Workflow struct {
	ID        int    `json:"id"`
	NodeID    string `json:"node_id"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	State     string `json:"state"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	URL       string `json:"url"`
	HTMLURL   string `json:"html_url"`
	BadgeURL  string `json:"badge_url"`
}
