package interp

import (
	"bytes"
	"testing"
)

func TestRunSourceHello(t *testing.T) {
	src := "package main\nfunc main() {\n  fmt.Println(\"hi\");\n}\n"
	var out bytes.Buffer
	code := RunSource("hello.go", src, &out)
	if code != 0 {
		t.Fatalf("код возврата = %d, want 0", code)
	}
	if out.String() != "hi\n" {
		t.Fatalf("вывод = %q", out.String())
	}
}

func TestRunSourceSyntaxError(t *testing.T) {
	src := "package main\nfunc main() {\n  fmt.Println(\"hi\")\n}\n" // нет ;
	var out bytes.Buffer
	code := RunSource("bad.go", src, &out)
	if code != 2 {
		t.Fatalf("код возврата = %d, want 2", code)
	}
}
