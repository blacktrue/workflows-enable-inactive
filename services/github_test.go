package services

import (
	"bytes"
	"github.com/blacktrue/workflows-enable-inactive/services/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestGithubService_GetWorkflows(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString("{\"total_count\": 1, \"workflows\": [{\"id\": 12345, \"state\": \"active\"}]}")),
		}

		ctrl := gomock.NewController(t)
		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil).Times(1)

		srv := NewGithubService(mockClient)
		workflows, err := srv.GetWorkflows("owner/fake-repository", "fake-token")
		assert.NoError(t, err)
		assert.Len(t, workflows, 1)
	})

	t.Run("API error", func(t *testing.T) {
		mockResponse := &http.Response{
			StatusCode: http.StatusServiceUnavailable,
			Body:       io.NopCloser(bytes.NewBufferString("{\"error\": \"error description\"}")),
		}

		ctrl := gomock.NewController(t)
		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil).Times(1)

		srv := NewGithubService(mockClient)
		workflows, err := srv.GetWorkflows("owner/fake-repository", "fake-token")
		assert.Error(t, err)
		assert.Len(t, workflows, 0)
	})
}

func TestGithubService_EnableWorkflow(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		mockResponse := &http.Response{
			StatusCode: http.StatusNoContent,
			Body:       nil,
		}

		ctrl := gomock.NewController(t)
		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil).Times(1)

		srv := NewGithubService(mockClient)
		updated, err := srv.EnableWorkflow(12345, "owner/fake-repository", "fake-token")
		assert.NoError(t, err)
		assert.True(t, updated)
	})

	t.Run("API error", func(t *testing.T) {
		mockResponse := &http.Response{
			StatusCode: http.StatusServiceUnavailable,
			Body:       io.NopCloser(bytes.NewBufferString("{\"error\": \"error description\"}")),
		}

		ctrl := gomock.NewController(t)
		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil).Times(1)

		srv := NewGithubService(mockClient)
		updated, err := srv.EnableWorkflow(12345, "owner/fake-repository", "fake-token")
		assert.Error(t, err)
		assert.False(t, updated)
	})
}
