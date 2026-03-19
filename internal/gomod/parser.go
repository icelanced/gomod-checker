package gomod

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"golang.org/x/mod/modfile"
)

type ModInfo struct {
	ModuleName string
	GoVersion  string
}

func findGoMod(root string) (string, error) {
	var result string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Name() == "go.mod" {
			result = path
			return fs.SkipAll
		}
		return nil
	})

	if err != nil {
		return "", fmt.Errorf("ошибка поиска go.mod: %w", err)
	}

	if result == "" {
		return "", fmt.Errorf("go.mod не найден в репозитории")
	}

	return result, nil
}

func Parse(dir string) (*ModInfo, error) {

	goModPath, err := findGoMod(dir)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(goModPath)
	if err != nil {
		return nil, fmt.Errorf("файл go.mod не найден: %w", err)
	}

	f, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга go.mod: %w", err)
	}

	goVer := "не указана"
	if f.Go != nil {
		goVer = f.Go.Version
	}

	moduleName := "не указан"
	if f.Module != nil {
		moduleName = f.Module.Mod.Path
	}

	return &ModInfo{
		ModuleName: moduleName,
		GoVersion:  goVer,
	}, nil
}
