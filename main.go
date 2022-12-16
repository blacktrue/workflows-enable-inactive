package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/blacktrue/workflows-enable-inactive/models"
	"github.com/blacktrue/workflows-enable-inactive/services"
	"github.com/blacktrue/workflows-enable-inactive/utils"
)

func hasError(repositories []models.ValidationResult) bool {
	for _, repo := range repositories {
		if repo.Error != "none" {
			return true
		}
	}

	return false
}

func main() {
	httpClient := &http.Client{}
	githubService := services.NewGithub(httpClient)
	workflowsService := services.NewWorkflowService(githubService)
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("The argument config path is required")
		os.Exit(1)
		return
	}
	cfg, err := utils.GetConfig(args[0])
	if err != nil {
		fmt.Printf("Error when get repositories: %s", err.Error())
		os.Exit(1)
		return
	}

	repositories, err := workflowsService.CheckAndEnableWorkflows(cfg)
	if err != nil {
		fmt.Printf("Error when check repositories: %s", err.Error())
		os.Exit(1)
		return
	}

	serialized, err := json.Marshal(repositories)
	if err != nil {
		fmt.Printf("Error when serialized ressult: %s", err.Error())
		os.Exit(1)
		return
	}

	code := 0
	if hasError(repositories) {
		code = 1
	}
	fmt.Printf("%s", string(serialized))
	os.Exit(code)
}
