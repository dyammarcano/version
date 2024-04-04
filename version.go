// this file version.go was generated with go generate command

package version

import (
	"fmt"
	"strings"
)

var i *info

type info struct {
	Version    string   `json:"version"`
	CommitHash string   `json:"commitHash"`
	Date       string   `json:"date"`
	Signature  string   `json:"signature"`
	Features   []string `json:"features"`
}

func init() {
	i = &info{
		Version:    "v0.0.0",
		CommitHash: "f0f717bc06750d9870e1bf825bdd0f9e842e27c7",
		Date:       "2024-04-04T08:20:22Z",
		Signature:  "FXiDqB3tbiotUAYmp8hbcVABv3431XDT4evaTPy5RExq",
		Features:   []string{},
	}
}

// AddFeature adds a feature description
func AddFeature(feature string) {
	i.Features = append(i.Features, fmt.Sprintf("+%s", feature))
}

// GetSignature returns the signature
func GetSignature() string {
	return i.Signature
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
		sb.WriteString(i.Date)
	}

	if len(i.Features) > 0 {
		sb.WriteString(" ")
		sb.WriteString(strings.Join(i.Features, " "))
	}

	return sb.String()
}
