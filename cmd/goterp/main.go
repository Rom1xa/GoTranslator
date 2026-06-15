package main

import (
	"fmt"
	"os"

	"github.com/rom1xa/kyrsoTranslatorGo/internal/interp"
)

func usage() {
	fmt.Fprintln(os.Stderr, "Использование:")
	fmt.Fprintln(os.Stderr, "  goterp <файл.go>")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Коды возврата:")
	fmt.Fprintln(os.Stderr, "  0  — успешное выполнение")
	fmt.Fprintln(os.Stderr, "  1  — ошибка аргументов или файла")
	fmt.Fprintln(os.Stderr, "  2  — лексическая / синтаксическая ошибка")
	fmt.Fprintln(os.Stderr, "  3  — семантическая ошибка")
	fmt.Fprintln(os.Stderr, "  4  — ошибка времени выполнения")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Пример:")
	fmt.Fprintln(os.Stderr, "  goterp program.go")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	path := os.Args[1]
	src, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "не удалось открыть файл %s: %v\n", path, err)
		os.Exit(1)
	}
	os.Exit(interp.RunSource(path, string(src), os.Stdout))
}
