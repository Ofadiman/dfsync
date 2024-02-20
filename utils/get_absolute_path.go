package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func GetAbsolutePath(path string) string {
	if strings.HasPrefix(path, "~/") {
		home, _ := os.UserHomeDir()

		return filepath.Clean(filepath.Join(home, strings.TrimPrefix(path, "~/")))
	}

	if strings.HasPrefix(path, "~") {
		home, _ := os.UserHomeDir()

		return filepath.Clean(filepath.Join(home, strings.TrimPrefix(path, "~")))
	}

	if strings.HasPrefix(path, "/") {
		return filepath.Clean(path)
	}

	abs, _ := filepath.Abs(path)
	return abs
}
