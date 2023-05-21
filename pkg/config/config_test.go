package config

import (
	"path"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	configPath := path.Join("testdata", "config/config.yml")
	config, err := New(configPath)
	if err != nil {
		t.Errorf("error load config: %s", err)
	}

	if config.Mode != "debug" {
		t.Errorf("expected debug but got: %s", config.Mode)
	}

	if config.DBConfig.MaximumPoolSize != 50 {
		t.Errorf("expected MaximumPoolSize 50 but got: %d", config.DBConfig.MaximumPoolSize)
	}
}
