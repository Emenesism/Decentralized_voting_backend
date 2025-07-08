package config

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	os.Setenv("TEST", "test_value")

	value := os.Getenv("TEST")

	if value != "test_value" {
		t.Errorf("Expected 'test_value', got '%s'", value)
	}

	os.Unsetenv("TEST")

	value = getEnv("TEST", "def_value")

	if value != "def_value" {
		t.Errorf("Expected 'def_value', got '%s'", value)
	}

}

func TestLoadConfig(t *testing.T) {
	os.Setenv("PORT", "9090")

	LoadConfig()

	if AppConfig.Port != "9090" {
		t.Errorf("Expected AppConfig.Port to be '9090', got '%s'", AppConfig.Port)
	}

	os.Unsetenv("PORT")
}
