package version

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

type (
	Version struct {
		Version    string   `json:"version"`
		CommitHash string   `json:"commitHash"`
		Date       string   `json:"date"`
		Features   []string `json:"features"`
	}

	VersionGerator struct {
		ProjectPath string
		Repo        *git.Repository
		Commit      *object.Commit
	}
)

// NewVersion creates a new VersionGerator
func NewVersion() (*VersionGerator, error) {
	projectPath, err := findGitRoot()
	if err != nil {
		return nil, err
	}

	repo, err := git.PlainOpen(projectPath) // open project dir
	if err != nil {
		return nil, err
	}

	ref, err := repo.Head()
	if err != nil {
		return nil, err
	}

	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return nil, err
	}

	p := &VersionGerator{
		ProjectPath: projectPath,
		Repo:        repo,
		Commit:      commit,
	}

	return p, nil
}

// Generate creates a version.go file in the destination path
func (v *VersionGerator) Generate(destPath string) error {
	destPath = filepath.Join(destPath, "version")

	// create folder if not exists
	if _, err := os.Stat(destPath); os.IsNotExist(err) {
		if err = os.MkdirAll(destPath, os.ModePerm); err != nil {
			return err
		}
	}

	versionFile := filepath.Join(destPath, "version.go")

	file, err := os.Create(versionFile)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		if err = file.Close(); err != nil {
			log.Fatalf("error closing file: %v", err)
		}
	}(file)

	date := v.Commit.Author.When.Format("2006-01-02T15:04:05Z")

	tag, err := v.getTag()
	if err != nil {
		return fmt.Errorf("error getting tag: %w", err)
	}

	data := &Version{
		Version:    tag,
		CommitHash: v.Commit.Hash.String(),
		Date:       date,
	}

	templateVersionFile := filepath.Join("version.tmpl")

	tmpl, err := template.New(templateVersionFile).ParseFiles(templateVersionFile)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	if err = tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	return nil
}

// getTag returns the tag of the repository
func (v *VersionGerator) getTag() (string, error) {
	tags, err := v.Repo.Tags()
	if err != nil {
		return "", err
	}

	tag := "v0.0.0"

	callback := func(ref *plumbing.Reference) error {
		if ref.Name().IsTag() {
			tag = ref.Name().Short()
		}

		return nil
	}

	if err = tags.ForEach(callback); err != nil {
		return "", err
	}

	return tag, nil
}

// findGitRoot returns the root path of the git repository
func findGitRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	projectRoot := strings.TrimSpace(string(output))

	return filepath.Clean(projectRoot), nil
}
