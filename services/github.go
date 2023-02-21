package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/blacktrue/workflows-enable-inactive/models"
	"github.com/blacktrue/workflows-enable-inactive/settings"
)

//go:generate mockgen -destination=./mocks/http_client_mock.go --build_flags=--mod=mod -package=mocks . HTTPClient
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Github struct {
	client HTTPClient
}

func NewGithub(client HTTPClient) Github {
	return Github{
		client: client,
	}
}

func (s Github) GetWorkflows(repository string, token string) ([]models.Workflow, error) {
	url := fmt.Sprintf(settings.Cfg.WorkflowsURL, repository)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []models.Workflow{}, fmt.Errorf("[services:github][method:get_workflows][new_request][error:%w]", err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	response, err := s.client.Do(req)
	if err != nil {
		return []models.Workflow{}, fmt.Errorf("[services:github][method:get_workflows][http_get][error:%w]", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK {
		return []models.Workflow{}, fmt.Errorf("[services:github][method:get_workflows][http_get][error:%s]", string(body))
	}
	if err != nil {
		return []models.Workflow{}, fmt.Errorf("[services:github][method:get_workflows][io_read_all][error:%w]", err)
	}
	res := models.WorkflowsResponse{}
	if err := json.Unmarshal(body, &res); err != nil {
		return []models.Workflow{}, fmt.Errorf("[services:github][method:get_workflows][json_unmarshal][error:%w]", err)
	}

	return res.Workflows, nil
}

func (s Github) EnableWorkflow(workflowID int32, repository string, token string) (bool, error) {
	url := fmt.Sprintf(settings.Cfg.EnableWorkflowURL, repository, workflowID)
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		return false, fmt.Errorf("[services:github][method:enable_workflow][new_request][error:%w]", err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	response, err := s.client.Do(req)
	if err != nil {
		return false, fmt.Errorf("[services:github][method:enable_workflow][http_put][error:%w]", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if response.StatusCode != http.StatusNoContent {
		return false, fmt.Errorf("[services:github][method:enable_workflow][http_put][error:%s]", string(body))
	}

	return true, nil
}
