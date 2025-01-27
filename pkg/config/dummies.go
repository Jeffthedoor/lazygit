package config

import (
	yaml "github.com/Jeffthedoor/yaml"
)

// NewDummyAppConfig creates a new dummy AppConfig for testing
func NewDummyAppConfig() *AppConfig {
	appConfig := &AppConfig{
		Name:        "lazygit",
		Version:     "unversioned",
		Commit:      "",
		BuildDate:   "",
		Debug:       false,
		BuildSource: "",
		UserConfig:  GetDefaultConfig(),
		AppState:    &AppState{},
	}
	_ = yaml.Unmarshal([]byte{}, appConfig.AppState)
	return appConfig
}
