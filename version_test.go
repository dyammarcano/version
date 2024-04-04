package version

import (
	"testing"
)

func TestVersion(t *testing.T) {
	AddFeature("mysql")
	AddFeature("kafka")
	AddFeature("redis")

	t.Log(GetVersionInfo())
}
