package config

import (
	"io/ioutil"

	v1 "github.com/jrbeverly/github-app-for-code-change/pkg/config/v1"
	"gopkg.in/yaml.v2"
)

type AppConfiguration struct {
	Version string `yaml:"version"`
}

// Read the configuration file from the provided path
func ReadV1Configuration(path string) (*v1.GitHubConfiguration, error) {
	cfg := &v1.GitHubConfiguration{}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
