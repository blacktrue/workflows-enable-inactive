package utils

import (
	"encoding/json"
	"io"
	"os"

	"github.com/blacktrue/workflows-enable-inactive/models"
)

func GetConfig(path string) (models.Config, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return models.Config{}, err
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	var cfg models.Config
	if err := json.Unmarshal(byteValue, &cfg); err != nil {
		return models.Config{}, err
	}

	return cfg, nil
}
