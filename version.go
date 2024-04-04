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
		Version:    "v0.0.1-dev",
		CommitHash: "f619bdc00a3461d6f0c16a83f40b08c0a0bf2496",
		Date:       "2024-04-04T07:58:55Z",
		Signature:  "8BGNzX2MvdncYtLdjgWUkkJ7uUjuFjAQ6FvKdtftsadF",
		Features:   []string{},
	}
}

// AddFeature adds a feature description
func AddFeature(feature string) {
	i.Features = append(i.Features, feature)
}

// GetFeatures returns the info
func GetFeatures() string {
	return fmt.Sprintf("Features: [%s]", strings.Join(i.Features, ", "))
}

// GetCommitHash returns the commit hash
func GetCommitHash() string {
	return fmt.Sprintf("CommitHash: %s", i.CommitHash)
}

// GetDate returns the date
func GetDate() string {
	return fmt.Sprintf("Date: %s", i.Date)
}

// GetVersion returns the info
func GetVersion() string {
	return fmt.Sprintf("Version: %s", i.Version)
}

// GetSignature returns the signature
func GetSignature() string {
	return fmt.Sprintf("Signature: %s", i.Signature)
}

func GetVersionInfo() string {
	return fmt.Sprintf("%s, %s, %s, %s, %s", GetVersion(), GetCommitHash(), GetDate(), GetSignature(), GetFeatures())
}
