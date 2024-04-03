////go:build generate

package main

import (
	"crypto/sha256"
	"github.com/caarlos0/log"
	"github.com/dyammarcano/base58"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"os"
	"path/filepath"
	"text/template"
)

//go:generate go run gen.go

type Version struct {
	Version    string   `json:"version"`
	CommitHash string   `json:"commitHash"`
	Date       string   `json:"date"`
	Signature  string   `json:"signature"`
	Features   []string `json:"features"`
}

var repo *git.Repository

func init() {
	projectPath, err := filepath.Abs("../../../")
	checkErr(err)

	log.Infof("open git project: %s", projectPath)
	repo, err = git.PlainOpen(projectPath) // open project dir
	checkErr(err)
}

func getCommit() (*object.Commit, error) {
	ref, err := repo.Head()
	if err != nil {
		return nil, err
	}

	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return nil, err
	}

	return commit, nil
}

func getCommitHash() (string, error) {
	commit, err := getCommit()
	if err != nil {
		return "", err
	}

	return commit.Hash.String(), nil
}

func getDate() (string, error) {
	commit, err := getCommit()
	if err != nil {
		return "", err
	}

	return commit.Author.When.Format("2006-01-02T15:04:05Z"), nil
}

func getTag() (string, error) {
	tags, err := repo.Tags()
	if err != nil {
		return "", err
	}

	var tag string
	callback := func(ref *plumbing.Reference) error {
		if ref.Name().IsTag() {
			tag = ref.Name().Short()
		}
		return nil
	}

	if err = tags.ForEach(callback); err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}

	return tag, nil
}

func sign(tag, commitHash, date string) string {
	h := sha256.New()
	h.Write([]byte(tag))
	h.Write([]byte(commitHash))
	h.Write([]byte(date))

	signData := base58.StdEncoding.EncodeToString(h.Sum(nil))
	log.Infof("generating signature: %s", signData)

	return signData
}

func checkErr(msg any) {
	if msg != nil {
		log.Errorf("error: %s", msg)
		os.Exit(1)
	}
}

func main() {
	versionFile := "version.go"
	templateVersionFile := filepath.Join("version.tmpl") //internal/version/generate/version.tmpl

	tmpl, err := template.New(templateVersionFile).ParseFiles(templateVersionFile)
	checkErr(err)

	err = os.Chdir("../..")
	checkErr(err)

	file, err := os.Create(versionFile)
	checkErr(err)

	commitHash, err := getCommitHash()
	checkErr(err)

	date, err := getDate()
	checkErr(err)

	tag, err := getTag()
	checkErr(err)

	if tag == "" {
		tag = "v0.0.1-dev"
	}

	data := Version{
		Version:    tag,
		CommitHash: commitHash,
		Date:       date,
		Signature:  sign(tag, commitHash, date),
	}

	log.Infof("generating version file: %s", versionFile)

	err = tmpl.Execute(file, data)
	checkErr(err)
}
