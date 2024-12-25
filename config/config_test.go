package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	t.Logf("Config loaded: %+v", cfg)

	t.Run("check env", func(t *testing.T) {
		assert.Equal(t, cfg.Env, "dev")
	})
}
