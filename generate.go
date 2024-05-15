package version

import (
	"encoding/json"
	"fmt"
	"github.com/caarlos0/log"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	defaultTag      = "v0.0.0"
	txtName         = "VERSION"
	versionFileName = "version.go"
)

type (
	Version struct {
		Version    string   `json:"version"`
		CommitHash string   `json:"commitHash"`
		Date       string   `json:"date"`
		Features   []string `json:"features,omitempty"`
	}

	Generator struct {
		projectPath string
		repo        *git.Repository
		commit      *object.Commit
	}
)

// NewVersion creates a new VersionGerator
func NewVersion() (*Generator, error) {
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

	p := &Generator{
		projectPath: projectPath,
		repo:        repo,
		commit:      commit,
	}

	return p, nil
}

// Generate creates a version.go file in the destination path
func (g *Generator) Generate() error {
	destPath := filepath.Join(g.projectPath, "internal", "version")

	// create folder if not exists
	if _, err := os.Stat(destPath); os.IsNotExist(err) {
		if err = os.MkdirAll(destPath, os.ModePerm); err != nil {
			return err
		}
	}

	versionFile := filepath.Join(destPath, versionFileName)

	log.Infof("generating version file: %s", versionFile)

	file, err := os.Create(versionFile)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		if err = file.Close(); err != nil {
			log.Fatalf("error closing file: %g", err)
		}
	}(file)

	date := g.commit.Author.When.Format("2006-01-02T15:04:05Z")

	tag, err := g.getTag()
	if err != nil {
		return fmt.Errorf("error getting tag: %w", err)
	}

	data := &Version{
		Version:    tag,
		CommitHash: g.commit.Hash.String(),
		Date:       date,
	}

	tmpl, err := template.New(versionFile).Parse(templateFile)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	if err = tmpl.ExecuteTemplate(file, versionFile, data); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	return g.genTxt(data)
}

// getTag returns the tag of the repository
func (g *Generator) getTag() (string, error) {
	tags, err := g.repo.Tags()
	if err != nil {
		return "", fmt.Errorf("error getting tags: %w", err)
	}

	tag := defaultTag

	callback := func(ref *plumbing.Reference) error {
		if ref.Name().IsTag() {
			tag = ref.Name().Short()
		}

		return nil
	}

	if err = tags.ForEach(callback); err != nil {
		return "", fmt.Errorf("error getting tags: %w", err)
	}

	return tag, nil
}

// genTxt creates a VERSION file with the version information
func (g *Generator) genTxt(ver *Version) error {
	txtFile := filepath.Join(g.projectPath, txtName)
	file, err := os.Create(txtFile)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer func(file *os.File) {
		if err = file.Close(); err != nil {
			log.Fatalf("error closing file: %g", err)
		}
	}(file)

	log.Infof("generating version file: %s", txtFile)

	if err = json.NewEncoder(file).Encode(ver); err != nil {
		return fmt.Errorf("error encoding json: %w", err)
	}

	return nil
}

// findGitRoot returns the root path of the git repository
func findGitRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error getting git root: %w", err)
	}

	projectRoot := strings.TrimSpace(string(output))

	return filepath.Clean(projectRoot), nil
}
