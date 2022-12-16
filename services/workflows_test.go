package services

import (
	"errors"
	"testing"

	"github.com/blacktrue/workflows-enable-inactive/models"
	"github.com/blacktrue/workflows-enable-inactive/services/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGithubService_CheckAndEnableWorkflows(t *testing.T) {
	cfgMock := models.Config{
		Repositories: []string{"owner/fake-repository"},
	}

	t.Run("Happy path", func(t *testing.T) {
		workflows := []models.Workflow{
			{
				State: "active",
				ID:    12345,
			},
		}
		ctrl := gomock.NewController(t)
		mockGithubSrv := mocks.NewMockGithubService(ctrl)
		mockGithubSrv.EXPECT().GetWorkflows(gomock.Any(), gomock.Any()).Return(workflows, nil).Times(1)

		srv := NewWorkflow(mockGithubSrv)
		validations, err := srv.CheckAndEnableWorkflows(cfgMock)
		assert.NoError(t, err)
		assert.Len(t, validations, 1)
	})

	t.Run("Workflows is disabled_inactivity", func(t *testing.T) {
		workflows := []models.Workflow{
			{
				State: "disabled_inactivity",
				ID:    12345,
			},
		}

		ctrl := gomock.NewController(t)
		mockGithubSrv := mocks.NewMockGithubService(ctrl)
		mockGithubSrv.EXPECT().GetWorkflows(gomock.Any(), gomock.Any()).Return(workflows, nil).Times(1)
		mockGithubSrv.EXPECT().EnableWorkflow(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil).Times(1)

		srv := NewWorkflow(mockGithubSrv)
		validations, err := srv.CheckAndEnableWorkflows(cfgMock)
		assert.NoError(t, err)
		assert.Len(t, validations, 1)
	})

	t.Run("API error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockGithubSrv := mocks.NewMockGithubService(ctrl)
		mockGithubSrv.EXPECT().GetWorkflows(gomock.Any(), gomock.Any()).Return([]models.Workflow{}, errors.New("fake error")).Times(1)
		mockGithubSrv.EXPECT().EnableWorkflow(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil).Times(0)

		srv := NewWorkflow(mockGithubSrv)
		validations, err := srv.CheckAndEnableWorkflows(cfgMock)
		assert.NoError(t, err)
		assert.Len(t, validations, 1)
		assert.NotEmpty(t, validations[0].Error)
	})
}
