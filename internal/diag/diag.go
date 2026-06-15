package diag

import "fmt"

type Level int

const (
	Lexical Level = iota
	Syntax
	Semantic
	Runtime
)

func (l Level) label() string {
	switch l {
	case Lexical:
		return "lexical error"
	case Syntax:
		return "syntax error"
	case Semantic:
		return "semantic error"
	default:
		return "runtime error"
	}
}

// exitCode: 2 — lexical/syntax, 3 — semantic, 4 — runtime.
func (l Level) exitCode() int {
	switch l {
	case Semantic:
		return 3
	case Runtime:
		return 4
	default:
		return 2
	}
}

type Error struct {
	Line, Col int
	Level     Level
	Msg       string
}

func (e Error) String() string {
	return fmt.Sprintf("%d:%d: %s: %s", e.Line, e.Col, e.Level.label(), e.Msg)
}

func (e Error) Error() string { return e.String() }

func (e Error) ExitCode() int { return e.Level.exitCode() }

type Bag struct{ errs []Error }

func (b *Bag) Add(e Error)     { b.errs = append(b.errs, e) }
func (b *Bag) HasErrors() bool { return len(b.errs) > 0 }
func (b *Bag) Errors() []Error { return b.errs }

func (b *Bag) ExitCode() int {
	code := 0
	for _, e := range b.errs {
		if c := e.Level.exitCode(); c > code {
			code = c
		}
	}
	return code
}
