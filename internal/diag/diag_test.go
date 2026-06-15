package diag

import "testing"

func TestErrorImplementsError(t *testing.T) {
	var err error = Error{Line: 1, Col: 2, Level: Semantic, Msg: "undefined: x"}
	if err.Error() != "1:2: semantic error: undefined: x" {
		t.Fatalf("Error() = %q", err.Error())
	}
}

func TestExitCodeOnError(t *testing.T) {
	if got := (Error{Level: Semantic}).ExitCode(); got != 3 {
		t.Fatalf("ExitCode Semantic = %d, want 3", got)
	}
	if got := (Error{Level: Runtime}).ExitCode(); got != 4 {
		t.Fatalf("ExitCode Runtime = %d, want 4", got)
	}
}

func TestBagCollectsAndFormats(t *testing.T) {
	var b Bag
	b.Add(Error{Line: 3, Col: 5, Level: Syntax, Msg: "missing ;"})
	if !b.HasErrors() {
		t.Fatal("ожидались ошибки")
	}
	if got := b.Errors()[0].String(); got != "3:5: syntax error: missing ;" {
		t.Fatalf("формат = %q", got)
	}
	if b.ExitCode() != 2 {
		t.Fatalf("ExitCode = %d, want 2", b.ExitCode())
	}
}

func TestExitCodeSemantic(t *testing.T) {
	var b Bag
	b.Add(Error{Line: 1, Col: 0, Level: Semantic, Msg: "undefined: x"})
	if got := b.ExitCode(); got != 3 {
		t.Fatalf("ExitCode для Semantic = %d, ожидалось 3", got)
	}
}

func TestExitCodeRuntime(t *testing.T) {
	var b Bag
	b.Add(Error{Line: 2, Col: 0, Level: Runtime, Msg: "division by zero"})
	if got := b.ExitCode(); got != 4 {
		t.Fatalf("ExitCode для Runtime = %d, ожидалось 4", got)
	}
}

func TestExitCodeMaxOfMultiple(t *testing.T) {
	var b Bag
	b.Add(Error{Line: 1, Col: 0, Level: Syntax, Msg: "missing ;"})
	b.Add(Error{Line: 2, Col: 0, Level: Runtime, Msg: "division by zero"})
	if got := b.ExitCode(); got != 4 {
		t.Fatalf("ExitCode (Syntax+Runtime) = %d, ожидалось 4", got)
	}
}
