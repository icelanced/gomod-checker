package gomod

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/mod/modfile"
)

type ModInfo struct {
	ModuleName string
	GoVersion  string
}

func Parse(dir string) (*ModInfo, error) {

	goModPath := filepath.Join(dir, "go.mod")

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
