package main

import (
	"fmt"
	"gomod-checker/internal/gomod"
	"gomod-checker/internal/updater"
	"os"
	"text/tabwriter"
)

func printResult(info *gomod.ModInfo, updates []updater.Module) {

	fmt.Println()
	fmt.Println()
	fmt.Printf("Модуль: %s\n", info.ModuleName)
	fmt.Printf("Версия Go:  %s\n", info.GoVersion)
	fmt.Println()
	fmt.Println()

	if len(updates) == 0 {
		fmt.Println("\nВсе зависимости актуальны, обновлений нет.")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "Зависимость\tТекущая версия\tНовая версия")
	fmt.Fprintln(w, "───────────\t──────────────\t────────────")

	for _, mod := range updates {
		fmt.Fprintf(w, "%s\t%s\t%s\n",
			mod.Path,
			mod.Version,
			mod.Update.Version,
		)
	}
	w.Flush()
	fmt.Println()
}
