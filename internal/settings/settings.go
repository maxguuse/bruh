package settings

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

type Settings struct {
	Project *ProjectDetails `yaml:"project"`
}

type ProjectDetails struct {
	Name  string `yaml:"name"`
	Owner string `yaml:"owner"`
}

const (
	SettingsFile = "bruh.yaml"
)

var (
	ErrFileNotFound    = errors.New("bruh.yaml file not found")
	ErrInvalidFile     = errors.New("Invalid bruh.yaml file format")
	ErrInvalidSettings = errors.New("Invalid settings")
)

func TryParse() (*Settings, error) {
	stg := &Settings{}

	blob, err := os.ReadFile(SettingsFile)
	if err != nil {
		return nil, ErrFileNotFound
	}

	err = yaml.Unmarshal(blob, stg)
	if err != nil {
		return nil, ErrInvalidFile
	}

	return stg, nil
}

func Write(stg *Settings) error {
	blob, err := yaml.Marshal(stg)
	if err != nil {
		return ErrInvalidSettings
	}

	err = os.WriteFile(SettingsFile, blob, 0644)
	if err != nil {
		return err
	}

	return nil
}
