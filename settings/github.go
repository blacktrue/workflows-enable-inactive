package settings

type Config struct {
	WorkflowsURL      string
	EnableWorkflowURL string
}

var Cfg Config

func init() {
	Cfg.WorkflowsURL = "https://api.github.com/repos/%s/actions/workflows"
	Cfg.EnableWorkflowURL = "https://api.github.com/repos/%s/actions/workflows/%d/enable"
}
