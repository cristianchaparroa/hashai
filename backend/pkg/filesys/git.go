package filesys

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

type GitFileSystem interface {
	GetGitRootPath() (string, error)
	GetGitFilePath(file string) (string, error)
}

type gitFileSystem struct {
}

func NewGitFileSystem() GitFileSystem {
	return &gitFileSystem{}
}

func (s *gitFileSystem) GetGitFilePath(file string) (string, error) {
	basePath, err := s.GetGitRootPath()
	if err != nil {
		return "", err
	}
	fmt.Println("--> base path: ", basePath)
	return filepath.Join(basePath, file), nil
}

func (s *gitFileSystem) GetGitRootPath() (string, error) {
	// Execute git rev-parse --show-toplevel to get repository root
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	// Get repository root path and trim newline
	return strings.TrimSpace(out.String()), nil
}
