package interp

import (
	"bytes"
	"testing"
)

func TestVarsAndAssign(t *testing.T) {
	src := "package main\nfunc main() {\n" +
		"  var x int = 10;\n" +
		"  y := 5;\n" +
		"  x = x + y;\n" +
		"  fmt.Println(x);\n" +
		"}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 0 {
		t.Fatalf("код = %d, want 0", code)
	}
	if out.String() != "15\n" {
		t.Fatalf("вывод = %q, want %q", out.String(), "15\n")
	}
}

func TestAssignUndefined(t *testing.T) {
	src := "package main\nfunc main() {\n  x = 1;\n}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 3 {
		t.Fatalf("код = %d, want 3 (semantic)", code)
	}
}

func TestArithmeticAndPrintf(t *testing.T) {
	src := "package main\nfunc main() {\n" +
		"  fmt.Println(2 + 3 * 4);\n" +
		"  fmt.Println(10 / 3);\n" +
		"  fmt.Println(7 % 3);\n" +
		"  fmt.Println(1.5 + 2.0);\n" +
		"  fmt.Println(2 > 1 && 1 == 1);\n" +
		"  fmt.Printf(\"%d-%s\\n\", 5, \"x\");\n" +
		"}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 0 {
		t.Fatalf("код = %d, want 0; вывод=%q", code, out.String())
	}
	want := "14\n3\n1\n3.5\ntrue\n5-x\n"
	if out.String() != want {
		t.Fatalf("вывод = %q, want %q", out.String(), want)
	}
}

func TestDivByZero(t *testing.T) {
	src := "package main\nfunc main() {\n  fmt.Println(1 / 0);\n}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 4 {
		t.Fatalf("код = %d, want 4 (runtime)", code)
	}
}

func TestTypeMismatchAdd(t *testing.T) {
	src := "package main\nfunc main() {\n  fmt.Println(1 + 1.5);\n}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 3 {
		t.Fatalf("код = %d, want 3 (semantic)", code)
	}
}

func TestAssignTypeMismatch(t *testing.T) {
	src := "package main\nfunc main() {\n  var x int = 5;\n  x = \"hi\";\n}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 3 {
		t.Fatalf("код = %d, want 3 (semantic): присваивание строки в int", code)
	}
}

func TestStringConcatAndUnary(t *testing.T) {
	src := "package main\nfunc main() {\n" +
		"  s := \"ab\" + \"cd\";\n" +
		"  fmt.Println(s);\n" +
		"  fmt.Println(!false);\n" +
		"  fmt.Println(-3 + 1);\n" +
		"}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 0 {
		t.Fatalf("код = %d, want 0; вывод=%q", code, out.String())
	}
	want := "abcd\ntrue\n-2\n"
	if out.String() != want {
		t.Fatalf("вывод = %q, want %q", out.String(), want)
	}
}

func TestVarZeroValue(t *testing.T) {
	src := "package main\nfunc main() {\n  var n int;\n  var f float64;\n  fmt.Println(n, f);\n}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 0 {
		t.Fatalf("код = %d, want 0", code)
	}
	if out.String() != "0 0\n" {
		t.Fatalf("вывод = %q, want %q", out.String(), "0 0\n")
	}
}

func TestRedeclareError(t *testing.T) {
	src := "package main\nfunc main() {\n  x := 1;\n  x := 2;\n}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 3 {
		t.Fatalf("код = %d, want 3 (повторное объявление)", code)
	}
}

func TestShadowingInBlock(t *testing.T) {
	src := "package main\nfunc main() {\n" +
		"  x := 1;\n" +
		"  {\n" +
		"    x := 2;\n" +
		"    fmt.Println(x);\n" +
		"  }\n" +
		"  fmt.Println(x);\n" +
		"}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 0 {
		t.Fatalf("код = %d, want 0; вывод=%q", code, out.String())
	}
	if out.String() != "2\n1\n" {
		t.Fatalf("вывод = %q, want %q", out.String(), "2\n1\n")
	}
}

func TestIfElseChain(t *testing.T) {
	src := "package main\nfunc main() {\n" +
		"  x := 2;\n" +
		"  if x == 1 {\n    fmt.Println(\"one\");\n" +
		"  } else if x == 2 {\n    fmt.Println(\"two\");\n" +
		"  } else {\n    fmt.Println(\"many\");\n  }\n" +
		"}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 0 {
		t.Fatalf("код = %d, want 0; вывод=%q", code, out.String())
	}
	if out.String() != "two\n" {
		t.Fatalf("вывод = %q, want %q", out.String(), "two\n")
	}
}

func TestIfNonBoolCond(t *testing.T) {
	src := "package main\nfunc main() {\n  if 1 {\n    fmt.Println(\"x\");\n  }\n}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 3 {
		t.Fatalf("код = %d, want 3 (условие не bool)", code)
	}
}

func TestForClause(t *testing.T) {
	src := "package main\nfunc main() {\n" +
		"  sum := 0;\n" +
		"  for i := 1; i <= 3; i = i + 1 {\n" +
		"    sum = sum + i;\n" +
		"  }\n" +
		"  fmt.Println(sum);\n" +
		"}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 0 {
		t.Fatalf("код = %d, want 0; вывод=%q", code, out.String())
	}
	if out.String() != "6\n" {
		t.Fatalf("вывод = %q, want %q", out.String(), "6\n")
	}
}

func TestForBreakContinue(t *testing.T) {
	src := "package main\nfunc main() {\n" +
		"  i := 0;\n" +
		"  for {\n" +
		"    i = i + 1;\n" +
		"    if i == 2 {\n      continue;\n    }\n" +
		"    if i > 4 {\n      break;\n    }\n" +
		"    fmt.Println(i);\n" +
		"  }\n" +
		"}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 0 {
		t.Fatalf("код = %d, want 0; вывод=%q", code, out.String())
	}
	if out.String() != "1\n3\n4\n" {
		t.Fatalf("вывод = %q, want %q", out.String(), "1\n3\n4\n")
	}
}

func TestBreakOutsideLoop(t *testing.T) {
	src := "package main\nfunc main() {\n  break;\n}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 3 {
		t.Fatalf("код = %d, want 3 (break вне цикла)", code)
	}
}

func TestContinueOutsideLoop(t *testing.T) {
	src := "package main\nfunc main() {\n  continue;\n}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 3 {
		t.Fatalf("код = %d, want 3 (continue вне цикла)", code)
	}
}

func TestFuncBasic(t *testing.T) {
	src := "package main\n" +
		"func add(a int, b int) int {\n  return a + b;\n}\n" +
		"func main() {\n  fmt.Println(add(3, 4));\n  x := add(10, 5);\n  fmt.Println(x);\n}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 0 {
		t.Fatalf("код = %d, want 0; вывод=%q", code, out.String())
	}
	if out.String() != "7\n15\n" {
		t.Fatalf("вывод = %q, want %q", out.String(), "7\n15\n")
	}
}

func TestFuncVoid(t *testing.T) {
	src := "package main\n" +
		"func greet(name string) {\n  fmt.Println(\"Hello\", name);\n}\n" +
		"func main() {\n  greet(\"World\");\n  greet(\"Go\");\n}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 0 {
		t.Fatalf("код = %d, want 0; вывод=%q", code, out.String())
	}
	if out.String() != "Hello World\nHello Go\n" {
		t.Fatalf("вывод = %q, want %q", out.String(), "Hello World\nHello Go\n")
	}
}

func TestFuncRecursion(t *testing.T) {
	src := "package main\n" +
		"func fib(n int) int {\n" +
		"  if n <= 1 {\n    return n;\n  }\n" +
		"  return fib(n-1) + fib(n-2);\n}\n" +
		"func main() {\n  fmt.Println(fib(0));\n  fmt.Println(fib(1));\n  fmt.Println(fib(7));\n}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 0 {
		t.Fatalf("код = %d, want 0; вывод=%q", code, out.String())
	}
	if out.String() != "0\n1\n13\n" {
		t.Fatalf("вывод = %q, want %q", out.String(), "0\n1\n13\n")
	}
}

func TestFuncMultiReturn(t *testing.T) {
	src := "package main\n" +
		"func minmax(a int, b int) (int, int) {\n" +
		"  if a < b {\n    return a, b;\n  }\n" +
		"  return b, a;\n}\n" +
		"func main() {\n" +
		"  lo, hi := minmax(5, 3);\n" +
		"  fmt.Println(lo, hi);\n" +
		"}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 0 {
		t.Fatalf("код = %d, want 0; вывод=%q", code, out.String())
	}
	if out.String() != "3 5\n" {
		t.Fatalf("вывод = %q, want %q", out.String(), "3 5\n")
	}
}

func TestFuncWrongArgCount(t *testing.T) {
	src := "package main\n" +
		"func add(a int, b int) int {\n  return a + b;\n}\n" +
		"func main() {\n  fmt.Println(add(1));\n}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 3 {
		t.Fatalf("код = %d, want 3 (неправильное число аргументов)", code)
	}
}

func TestFuncWrongArgType(t *testing.T) {
	src := "package main\n" +
		"func double(x int) int {\n  return x * 2;\n}\n" +
		"func main() {\n  fmt.Println(double(true));\n}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 3 {
		t.Fatalf("код = %d, want 3 (несоответствие типа аргумента)", code)
	}
}

func TestFuncUndefFunc(t *testing.T) {
	src := "package main\nfunc main() {\n  foo();\n}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 3 {
		t.Fatalf("код = %d, want 3 (вызов необъявленной функции)", code)
	}
}

func TestFuncReturnTypeMismatch(t *testing.T) {
	src := "package main\n" +
		"func getInt() int {\n  return true;\n}\n" +
		"func main() {\n  fmt.Println(getInt());\n}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 3 {
		t.Fatalf("код = %d, want 3 (несоответствие типа возврата)", code)
	}
}

func TestStructBasic(t *testing.T) {
	src := "package main\n" +
		"type Point struct { x int; y int; }\n" +
		"func main() {\n" +
		"  p := Point{x: 3, y: 4};\n" +
		"  fmt.Println(p.x);\n" +
		"  fmt.Println(p.y);\n" +
		"  p.x = 10;\n" +
		"  fmt.Println(p.x);\n" +
		"}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 0 {
		t.Fatalf("код = %d, want 0; вывод=%q", code, out.String())
	}
	if out.String() != "3\n4\n10\n" {
		t.Fatalf("вывод = %q, want %q", out.String(), "3\n4\n10\n")
	}
}

func TestStructZeroInit(t *testing.T) {
	src := "package main\n" +
		"type Pair struct { a int; b string; }\n" +
		"func main() {\n" +
		"  var p Pair;\n" +
		"  fmt.Println(p.a);\n" +
		"  fmt.Println(p.b);\n" +
		"}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 0 {
		t.Fatalf("код = %d, want 0; вывод=%q", code, out.String())
	}
	if out.String() != "0\n\n" {
		t.Fatalf("вывод = %q, want %q", out.String(), "0\n\n")
	}
}

func TestStructInFunc(t *testing.T) {
	src := "package main\n" +
		"type Vec struct { x int; y int; }\n" +
		"func makeVec(x int, y int) Vec {\n" +
		"  return Vec{x: x, y: y};\n" +
		"}\n" +
		"func sumVec(v Vec) int {\n" +
		"  return v.x + v.y;\n" +
		"}\n" +
		"func main() {\n" +
		"  v := makeVec(3, 4);\n" +
		"  fmt.Println(sumVec(v));\n" +
		"}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 0 {
		t.Fatalf("код = %d, want 0; вывод=%q", code, out.String())
	}
	if out.String() != "7\n" {
		t.Fatalf("вывод = %q, want %q", out.String(), "7\n")
	}
}

func TestStructUndefField(t *testing.T) {
	src := "package main\n" +
		"type Box struct { w int; }\n" +
		"func main() {\n" +
		"  b := Box{w: 1};\n" +
		"  fmt.Println(b.h);\n" +
		"}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 3 {
		t.Fatalf("код = %d, want 3 (несуществующее поле)", code)
	}
}

func TestStructUndefType(t *testing.T) {
	src := "package main\n" +
		"func main() {\n" +
		"  p := Ghost{x: 1};\n" +
		"  fmt.Println(p.x);\n" +
		"}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 3 {
		t.Fatalf("код = %d, want 3 (неизвестный тип)", code)
	}
}

func TestSliceBasic(t *testing.T) {
	src := "package main\nfunc main() {\n" +
		"  s := []int{10, 20, 30};\n" +
		"  fmt.Println(s[0]);\n" +
		"  fmt.Println(len(s));\n" +
		"  s[1] = 99;\n" +
		"  fmt.Println(s[1]);\n" +
		"}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 0 {
		t.Fatalf("код = %d, want 0; вывод=%q", code, out.String())
	}
	if out.String() != "10\n3\n99\n" {
		t.Fatalf("вывод = %q, want %q", out.String(), "10\n3\n99\n")
	}
}

func TestSliceAppend(t *testing.T) {
	src := "package main\nfunc main() {\n" +
		"  s := []int{1, 2};\n" +
		"  s = append(s, 3);\n" +
		"  s = append(s, 4);\n" +
		"  fmt.Println(len(s));\n" +
		"  fmt.Println(s[3]);\n" +
		"}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 0 {
		t.Fatalf("код = %d, want 0; вывод=%q", code, out.String())
	}
	if out.String() != "4\n4\n" {
		t.Fatalf("вывод = %q, want %q", out.String(), "4\n4\n")
	}
}

func TestSliceRange(t *testing.T) {
	src := "package main\nfunc main() {\n" +
		"  s := []int{5, 6, 7};\n" +
		"  sum := 0;\n" +
		"  for i, v := range s {\n" +
		"    sum = sum + v;\n" +
		"    fmt.Println(i);\n" +
		"  }\n" +
		"  fmt.Println(sum);\n" +
		"}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 0 {
		t.Fatalf("код = %d, want 0; вывод=%q", code, out.String())
	}
	if out.String() != "0\n1\n2\n18\n" {
		t.Fatalf("вывод = %q, want %q", out.String(), "0\n1\n2\n18\n")
	}
}

func TestSliceRangeIndex(t *testing.T) {
	src := "package main\nfunc main() {\n" +
		"  s := []string{\"a\", \"b\", \"c\"};\n" +
		"  for i := range s {\n" +
		"    fmt.Println(i);\n" +
		"  }\n" +
		"}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 0 {
		t.Fatalf("код = %d, want 0; вывод=%q", code, out.String())
	}
	if out.String() != "0\n1\n2\n" {
		t.Fatalf("вывод = %q, want %q", out.String(), "0\n1\n2\n")
	}
}

func TestSliceOOB(t *testing.T) {
	src := "package main\nfunc main() {\n" +
		"  s := []int{1, 2};\n" +
		"  fmt.Println(s[5]);\n" +
		"}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 4 {
		t.Fatalf("код = %d, want 4 (runtime: out of bounds)", code)
	}
}

func TestSliceAppendTypeMismatch(t *testing.T) {
	src := "package main\nfunc main() {\n" +
		"  s := []int{1, 2};\n" +
		"  s = append(s, true);\n" +
		"}\n"
	var out bytes.Buffer
	if code := RunSource("t.go", src, &out); code != 3 {
		t.Fatalf("код = %d, want 3 (semantic: type mismatch in append)", code)
	}
}
