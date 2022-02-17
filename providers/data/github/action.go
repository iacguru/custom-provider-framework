package action

import "time"

type Workflows struct {
	TotalCount int `json:"total_count"`
	Workflows  []Workflow
}

type Workflow struct {
	ID        int       `json:"id"`
	NodeID    string    `json:"node_id"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	State     string    `json:"state"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	URL       string    `json:"url"`
	HTMLURL   string    `json:"html_url"`
	BadgeURL  string    `json:"badge_url"`
}
