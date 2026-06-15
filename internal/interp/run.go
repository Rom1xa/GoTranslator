package interp

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/rom1xa/kyrsoTranslatorGo/internal/builder"
	"github.com/rom1xa/kyrsoTranslatorGo/internal/diag"
)

func RunSource(name, src string, stdout io.Writer) int {
	prog, bag := builder.Parse(name, src)
	if bag.HasErrors() {
		for _, e := range bag.Errors() {
			fmt.Fprintf(os.Stderr, "%s: %s\n", name, e.String())
		}
		return bag.ExitCode()
	}
	if err := New(stdout).Run(prog); err != nil {
		var de diag.Error
		if errors.As(err, &de) {
			fmt.Fprintf(os.Stderr, "%s: %s\n", name, de.String())
			return de.ExitCode()
		}
		fmt.Fprintf(os.Stderr, "%s: runtime error: %v\n", name, err)
		return 4
	}
	return 0
}
