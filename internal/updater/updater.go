package updater

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type Update struct {
	Path    string `json:"Path"`
	Version string `json:"Version"`
}

type Module struct {
	Path    string  `json:"Path"`
	Version string  `json:"Version"`
	Update  *Update `json:"Update"`
	Main    bool    `json:"Main"`
}

func GetUpdates(dir string) ([]Module, error) {

	cmd := exec.Command("go", "list", "-m", "-u", "-json", "all")

	cmd.Dir = dir

	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("`go list` завершился с ошибкой: %w", err)
	}

	var modules []Module
	decoder := json.NewDecoder(strings.NewReader(string(out)))

	for decoder.More() {
		var mod Module
		if err := decoder.Decode(&mod); err != nil {
			return nil, fmt.Errorf("ошибка декодирования json: %w", err)
		}

		if !mod.Main && mod.Update != nil {
			modules = append(modules, mod)
		}
	}
	return modules, nil
}
