package value

import "testing"

func TestStringRepr(t *testing.T) {
	if got := Int(42).String(); got != "42" {
		t.Fatalf("Int.String() = %q, want %q", got, "42")
	}
	if got := String("hi").String(); got != "hi" {
		t.Fatalf("String.String() = %q, want %q", got, "hi")
	}
	if got := Int(0).String(); got != "0" {
		t.Fatalf("Int(0).String() = %q, want %q", got, "0")
	}
	if got := Int(-1).String(); got != "-1" {
		t.Fatalf("Int(-1).String() = %q, want %q", got, "-1")
	}
}

func TestNative(t *testing.T) {
	cases := []struct {
		v    Value
		want any
	}{
		{Int(42), int64(42)},
		{Float(3.5), float64(3.5)},
		{String("hi"), "hi"},
		{Bool(true), true},
		{Bool(false), false},
	}
	for _, c := range cases {
		if got := c.v.Native(); got != c.want {
			t.Fatalf("%T.Native() = %#v, want %#v", c.v, got, c.want)
		}
	}
}

func TestFloatAndBoolString(t *testing.T) {
	if got := Float(3.5).String(); got != "3.5" {
		t.Fatalf("Float(3.5).String() = %q, want %q", got, "3.5")
	}
	if got := Bool(true).String(); got != "true" {
		t.Fatalf("Bool(true).String() = %q, want %q", got, "true")
	}
	if got := Bool(false).String(); got != "false" {
		t.Fatalf("Bool(false).String() = %q, want %q", got, "false")
	}
}
