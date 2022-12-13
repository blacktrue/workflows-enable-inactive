package services

import (
	"github.com/blacktrue/workflows-enable-inactive/models"
)

//go:generate mockgen -destination=./mocks/github_service_mock.go --build_flags=--mod=mod -package=mocks . GithubSrv
type GithubSrv interface {
	GetWorkflows(repository string, token string) ([]models.Workflow, error)
	EnableWorkflow(workflowID int32, repository string, token string) (bool, error)
}

type WorkflowsService struct {
	github GithubSrv
}

func NewWorkflowService(service GithubSrv) WorkflowsService {
	return WorkflowsService{
		github: service,
	}
}

func (s WorkflowsService) CheckAndEnableWorkflows(cfg models.Config) ([]models.ValidationResult, error) {
	results := make([]models.ValidationResult, 0)
	for _, repository := range cfg.Repositories {
		workflows, err := s.github.GetWorkflows(repository, cfg.Token)
		if err != nil {
			errorMsg := err.Error()
			results = append(results, models.ValidationResult{
				Repository: repository,
				Updated:    false,
				Error:      errorMsg,
			})
			continue
		}
		hasDisableWorkflows, workflow := hasDisableWorkflow(workflows)
		var updated bool
		errorMsg := "none"
		if hasDisableWorkflows {
			updated, err = s.github.EnableWorkflow(workflow.ID, repository, cfg.Token)
			if err != nil {
				errorMsg = err.Error()
			}
		}

		results = append(results, models.ValidationResult{
			Repository: repository,
			Updated:    updated,
			Error:      errorMsg,
		})
	}

	return results, nil
}

func hasDisableWorkflow(workflows []models.Workflow) (bool, models.Workflow) {
	for _, workflow := range workflows {
		if workflowIsDisabled(workflow) {
			return true, workflow
		}
	}

	return false, models.Workflow{}
}

func workflowIsDisabled(workflow models.Workflow) bool {
	return workflow.State == "disabled_inactivity"
}
