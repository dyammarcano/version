package version

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersion(t *testing.T) {
	AddFeature("feature1")
	AddFeature("feature2")
	AddFeature("feature3")

	g := Get()

	assert.Equal(t, "v0.0.1-dev", g.Version)
	assert.Equal(t, "feature1", g.Features[0])
}
