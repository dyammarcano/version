// this file version.go was generated with go generate command

package version

import "fmt"

var ver *Version

type Version struct {
	Version    string   `json:"version"`
	CommitHash string   `json:"commitHash"`
	Date       string   `json:"date"`
	Signature  string   `json:"signature"`
	Features   []string `json:"features"`
}

func init() {
	ver = &Version{
		Version:    "v0.0.1-dev",
		CommitHash: "27f8696fd83d8420fdce788afd259c1ffcf8ed9e",
		Date:       "2024-04-03T15:05:54Z",
		Signature:  "2ZhrXZRLBymyvmDhbReXW6xtxi5Ynhm5WaurSEHk4GT7",
		Features:   []string{},
	}
}

// AddFeature adds a feature description
func AddFeature(feature string) {
	ver.Features = append(ver.Features, feature)
}

// Get returns the Info struct
func Get() Version {
	return *ver
}

// GetFeatures returns the version
func GetFeatures() string {
	return fmt.Sprintf("Features: %s", ver.Features)
}

// GetCommitHash returns the commit hash
func GetCommitHash() string {
	return fmt.Sprintf("CommitHash: %s", ver.CommitHash)
}

// GetDate returns the date
func GetDate() string {
	return fmt.Sprintf("Date: %s", ver.Date)
}

// GetVersion returns the version
func GetVersion() string {
	return fmt.Sprintf("Version: %s", ver.Version)
}
