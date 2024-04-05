package version

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestNewVersion(t *testing.T) {
	v, err := NewVersion()
	assert.NoError(t, err)

	err = v.Generate(v.ProjectPath)
	assert.NoError(t, err)

	destPath := filepath.Join(v.ProjectPath, "version", "version.go")

	if _, err = os.Stat(destPath); os.IsNotExist(err) {
		if err = os.MkdirAll(destPath, os.ModePerm); err != nil {
			assert.Error(t, err)
		}
	}

	err = os.RemoveAll(destPath)
	assert.NoError(t, err)
}
