package version

const templateFile = `// this file version.go was generated with go generate command

package version

import (
	"fmt"
	"strings"
)

var i *info

type info struct {
	Version    string   ` + "`json:\"version\"`" + `
	CommitHash string   ` + "`json:\"commitHash\"`" + `
	Date       string   ` + "`json:\"date\"`" + `
	Features   []string ` + "`json:\"features\"`" + `
}

func init() {
	i = &info{
		Version:    "{{.Version}}",
		CommitHash: "{{.CommitHash}}",
		Date:       "{{.Date}}",
		Features:   []string{},
	}
}

// AddFeature adds a feature description
func AddFeature(feature string) {
	i.Features = append(i.Features, fmt.Sprintf("+%s", feature))
}

// GetVersionInfo returns the info
func GetVersionInfo() string {
	var sb strings.Builder
	sb.WriteString(i.Version)

	if i.CommitHash != "" {
		sb.WriteString("-")
		sb.WriteString(i.CommitHash)
	}

	if i.Date != "" {
		sb.WriteString("-")
		sb.WriteString(i.Date[:10]) // format date to yyyy-mm-dd
	}

	if len(i.Features) > 0 {
		sb.WriteString(" ")
		sb.WriteString(strings.Join(i.Features, " "))
	}

	return sb.String()
}
`
