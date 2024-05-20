package generate

//
//import (
//	"github.com/spf13/afero"
//	"github.com/stretchr/testify/assert"
//	"os/exec"
//	"testing"
//)
//
//func TestNewVersion(t *testing.T) {
//	fs = afero.NewOsFs()
//
//	tPath, err := afero.TempDir(fs, "", "test")
//	assert.NoError(t, err)
//
//	err = fs.MkdirAll(tPath, 0755)
//	assert.NoError(t, err)
//
//	_, err = fs.Stat(tPath)
//	assert.NoError(t, err)
//
//	// create git init in temp dir
//	err = exec.Command("git", "init", tPath).Run()
//	assert.NoError(t, err)
//
//	v, err := NewVersion("test")
//	assert.NoError(t, err)
//
//	err = v.Generate()
//	assert.NoError(t, err)
//}
