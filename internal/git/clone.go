package git

import (
	"fmt"
	"os"
	"os/exec"
)

func Clone(url string) (dir string, cleanup func(), err error) {
	tempDir, err := os.MkdirTemp("", "gomod-checker-*")
	if err != nil {
		return "", nil, fmt.Errorf("не удалось создать temp-директорию: %w", err)
	}

	cleanup = func() {
		os.RemoveAll(tempDir)
	}

	commands := [][]string{
		{"git", "init", tempDir},
		{"git", "-C", tempDir, "remote", "add", "origin", url},
		{"git", "-C", tempDir, "sparse-checkout", "init"},
		{"git", "-C", tempDir, "sparse-checkout", "set", "go.mod", "go.sum"},
		{"git", "-C", tempDir, "pull", "--depth=1", "origin", "HEAD"},
	}

	for _, args := range commands {
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			cleanup()
			return "", nil, fmt.Errorf("команда %v завершилась с ошибкой: %w", args, err)
		}
	}

	return tempDir, cleanup, nil
}
