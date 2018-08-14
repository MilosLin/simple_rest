package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_NormalUsage : 一般使用案例
func Test_NormalUsage(t *testing.T) {
	assert.Equal(t, 6215, Forge().GetInt("API.ListenPort"))
	assert.Equal(t, "world", Forge().GetString("env.hello"))
}
