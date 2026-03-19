package main

import (
	"fmt"
	"gomod-checker/internal/git"
	"gomod-checker/internal/gomod"
	"gomod-checker/internal/updater"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("использование: gomod-checker <git-repo-url>")
	}

	repoUrl := os.Args[1]

	fmt.Printf("Клонирую репозиторий %s\n\n", repoUrl)
	repoDir, cleanup, err := git.Clone(repoUrl)
	if err != nil {
		return fmt.Errorf("ошибка клонирования: %w", err)
	}
	defer cleanup()

	modInfo, err := gomod.Parse(repoDir)
	if err != nil {
		return fmt.Errorf("ошибка чтения go.mod: %w", err)
	}

	fmt.Println("Проверяю наличие обновлений...")
	updates, err := updater.GetUpdates(repoDir)
	if err != nil {
		return fmt.Errorf("ошибка получения обновлений: %w", err)
	}

	printResult(modInfo, updates)
	return nil
}
