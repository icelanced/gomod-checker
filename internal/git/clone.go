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

	cmd := exec.Command("git", "clone", "--depth=1", url, tempDir)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		cleanup()
		return "", nil, fmt.Errorf("git-clone завершился с ошибкой: %w", err)
	}
	return tempDir, cleanup, nil
}
