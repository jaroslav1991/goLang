package config

import (
	"os"
	"testing"
)

func TestGetDbConfig(t *testing.T) {
	// test success case
	t.Run("local env", func(t *testing.T) {
		actualConfig, actualErr := GetDbConfig("local")

		if actualErr != nil {
			t.Error("actualErr is not nil for local env: actual is", actualErr)
		}

		if actualConfig != localDbConfig {
			t.Error("actualConfig is invalid: actual is", actualConfig)
		}
	})

	// test unknown env case
	t.Run("unknown env", func(t *testing.T) {
		_, actualErr := GetDbConfig("production")
		if actualErr == nil {
			t.Error("actualErr is nil")
		}
	})
}

func TestGetAppEnvironment(t *testing.T) {
	t.Run("default env is local", func(t *testing.T) {
		actualEnv := GetAppEnvironment()
		if actualEnv != "local" {
			t.Error("default env is not local: actual is", actualEnv)
		}
	})

	t.Run("env is correct", func(t *testing.T) {
		expectedEnv := "production"

		if err := os.Setenv("APP_ENV", expectedEnv); err != nil {
			t.Error("os.SetEnv failed with error", err)
		}

		if actualEnv := GetAppEnvironment(); actualEnv != expectedEnv {
			t.Errorf("actualEnv != expectedEnv: '%s' != '%s'", actualEnv, expectedEnv)
		}
	})
}
