package version

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewVersion(t *testing.T) {
	v, err := NewVersion()
	assert.NoError(t, err)

	err = v.Generate(v.ProjectPath)
	assert.NoError(t, err)
}
