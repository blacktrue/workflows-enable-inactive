package models

type WorkflowsResponse struct {
	TotalCount int32      `json:"total_count"`
	Workflows  []Workflow `json:"workflows"`
}

type Workflow struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
}

type ValidationResult struct {
	Repository string `json:"repository"`
	Error      string `json:"error"`
	Updated    bool   `json:"updated"`
}
