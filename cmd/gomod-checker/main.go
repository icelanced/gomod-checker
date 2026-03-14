package main

import (
	"fmt"
	"gomod-checker/internal/git"
	"gomod-checker/internal/gomod"
	"gomod-checker/internal/updater"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Использование: gomod-checker <git-repo-url>")
		os.Exit(1)
	}

	repoUrl := os.Args[1]

	fmt.Printf("Клонирую репозиторий %s\n\n", repoUrl)
	repoDir, cleanup, err := git.Clone(repoUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка клонирования: %v\n", err)
		os.Exit(1)
	}
	defer cleanup()

	modInfo, err := gomod.Parse(repoDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка чтения go.mod: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Проверяю наличие обновлений...")
	updates, err := updater.GetUpdates(repoDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка получения обновлений: %v\n", err)
		os.Exit(1)
	}

	printResult(modInfo, updates)
}
