package version

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersion(t *testing.T) {
	AddFeature("mysql")
	AddFeature("kafka")
	AddFeature("redis")

	assert.Equal(t, "Version: v0.0.1-dev", GetVersion())
	assert.Equal(t, "Features: [mysql, kafka, redis]", GetFeatures())

	t.Log(GetVersionInfo())
}
