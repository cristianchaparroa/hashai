package main

import (
	"bytes"
	"context"
	"fmt"
	"hashtracker/config"
	"hashtracker/internal/usecases/repo/polygon"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	polygonRepo := polygon.NewPolygonRepository(cfg)
	l := NewLoader(polygonRepo)
	filePath, getFileError := getFilePath()
	if getFileError != nil {
		panic(getFileError)
	}
	fmt.Println("--> dataset file: ", filePath)

	loadErr := l.Load(context.Background(), filePath)
	if loadErr != nil {
		panic(loadErr)
	}
}

func getFilePath() (string, error) {
	basePath, err := getGitHubPath()
	if err != nil {
		return "", err
	}
	fmt.Println("--> base path: ", basePath)
	return filepath.Join(basePath, "backend/migrations/datasets/addresses-darklist-001.json"), nil
}

func getGitHubPath() (string, error) {
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
