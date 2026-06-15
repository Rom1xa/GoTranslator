package tests

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/rom1xa/kyrsoTranslatorGo/internal/interp"
)

func TestNegativePrograms(t *testing.T) {
	matches, _ := filepath.Glob("testdata/negative/*.go")
	if len(matches) == 0 {
		t.Fatal("testdata/negative не содержит ни одного *.go файла")
	}
	for _, src := range matches {
		src := src
		t.Run(filepath.Base(src), func(t *testing.T) {
			code, err := os.ReadFile(src)
			if err != nil {
				t.Fatal(err)
			}
			var out bytes.Buffer
			rc := interp.RunSource(src, string(code), &out)
			if rc == 0 {
				t.Fatalf("ожидался ненулевой код возврата, получено 0")
			}
		})
	}
}

func TestPositivePrograms(t *testing.T) {
	matches, _ := filepath.Glob("testdata/positive/*.go")
	if len(matches) == 0 {
		t.Fatal("testdata/positive не содержит ни одного *.go файла")
	}
	for _, src := range matches {
		src := src
		t.Run(filepath.Base(src), func(t *testing.T) {
			code, err := os.ReadFile(src)
			if err != nil {
				t.Fatal(err)
			}
			want, err := os.ReadFile(src[:len(src)-len(".go")] + ".out")
			if err != nil {
				t.Fatal(err)
			}
			var out bytes.Buffer
			rc := interp.RunSource(src, string(code), &out)
			if rc != 0 {
				t.Fatalf("код возврата = %d, want 0", rc)
			}
			if out.String() != string(want) {
				t.Fatalf("вывод = %q, want %q", out.String(), string(want))
			}
		})
	}
}
