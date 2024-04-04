// this file info.go was generated with go generate command

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
		CommitHash: "27f8696fd83d8420fdce788afd259c1ffcf8ed9e",
		Date:       "2024-04-03T15:05:54Z",
		Signature:  "2ZhrXZRLBymyvmDhbReXW6xtxi5Ynhm5WaurSEHk4GT7",
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
