package builder

import (
	"testing"

	"github.com/rom1xa/kyrsoTranslatorGo/internal/ast"
)

func TestParseVarAndExpr(t *testing.T) {
	src := "package main\nfunc main() {\n" +
		"  var x int = 2 + 3 * 4;\n" +
		"  y := x > 5;\n" +
		"  fmt.Println(x, y);\n" +
		"}\n"
	prog, bag := Parse("t.go", src)
	if bag.HasErrors() {
		t.Fatalf("неожиданные ошибки: %v", bag.Errors())
	}
	body := prog.Funcs[0].Body
	if len(body) != 3 {
		t.Fatalf("ожидалось 3 оператора, получено %d", len(body))
	}
	vd, ok := body[0].(*ast.VarDecl)
	if !ok || vd.Name != "x" || vd.Type != "int" {
		t.Fatalf("ожидался var x int, получено %+v", body[0])
	}
	// 2 + 3*4 => BinaryExpr("+", 2, BinaryExpr("*", 3, 4))
	add, ok := vd.Init.(*ast.BinaryExpr)
	if !ok || add.Op != "+" {
		t.Fatalf("ожидался '+', получено %+v", vd.Init)
	}
	if mul, ok := add.Right.(*ast.BinaryExpr); !ok || mul.Op != "*" {
		t.Fatalf("ожидался '*' справа (приоритет), получено %+v", add.Right)
	}
	if _, ok := body[1].(*ast.ShortVarDecl); !ok {
		t.Fatalf("ожидался ShortVarDecl, получено %+v", body[1])
	}
	ps, ok := body[2].(*ast.PrintStmt)
	if !ok || ps.Fn != "Println" || len(ps.Args) != 2 {
		t.Fatalf("ожидался Println с 2 арг., получено %+v", body[2])
	}
}

func TestParseLogicalPrecedence(t *testing.T) {
	src := "package main\nfunc main() {\n" +
		"  ok := true || false && false;\n" +
		"  fmt.Println(ok);\n" +
		"}\n"
	prog, bag := Parse("t.go", src)
	if bag.HasErrors() {
		t.Fatalf("неожиданные ошибки: %v", bag.Errors())
	}
	svd := prog.Funcs[0].Body[0].(*ast.ShortVarDecl)
	or, ok := svd.Value.(*ast.BinaryExpr)
	if !ok || or.Op != "||" {
		t.Fatalf("ожидался верхний '||', получено %+v", svd.Value)
	}
	if l, ok := or.Left.(*ast.BoolLit); !ok || l.Value != true {
		t.Fatalf("ожидался BoolLit(true) слева, получено %+v", or.Left)
	}
	if and, ok := or.Right.(*ast.BinaryExpr); !ok || and.Op != "&&" {
		t.Fatalf("ожидался '&&' справа (крепче ||), получено %+v", or.Right)
	}
}

func TestParseUnaryAndFloat(t *testing.T) {
	src := "package main\nfunc main() {\n  z := -3.5;\n}\n"
	prog, bag := Parse("t.go", src)
	if bag.HasErrors() {
		t.Fatalf("неожиданные ошибки: %v", bag.Errors())
	}
	svd := prog.Funcs[0].Body[0].(*ast.ShortVarDecl)
	un, ok := svd.Value.(*ast.UnaryExpr)
	if !ok || un.Op != "-" {
		t.Fatalf("ожидался унарный '-', получено %+v", svd.Value)
	}
	if f, ok := un.Operand.(*ast.FloatLit); !ok || f.Value != 3.5 {
		t.Fatalf("ожидался FloatLit(3.5), получено %+v", un.Operand)
	}
}

func TestParseControlFlow(t *testing.T) {
	src := "package main\nfunc main() {\n" +
		"  i := 0;\n" +
		"  for i < 3 {\n" +
		"    if i == 1 {\n" +
		"      fmt.Println(\"one\");\n" +
		"    } else {\n" +
		"      fmt.Println(\"other\");\n" +
		"    }\n" +
		"    i = i + 1;\n" +
		"  }\n" +
		"  for j := 0; j < 2; j = j + 1 {\n" +
		"    break;\n" +
		"  }\n" +
		"}\n"
	prog, bag := Parse("t.go", src)
	if bag.HasErrors() {
		t.Fatalf("неожиданные ошибки: %v", bag.Errors())
	}
	body := prog.Funcs[0].Body
	fw, ok := body[1].(*ast.ForStmt)
	if !ok || fw.Cond == nil || fw.Init != nil {
		t.Fatalf("ожидался for-while, получено %+v", body[1])
	}
	ifst, ok := fw.Body.Stmts[0].(*ast.IfStmt)
	if !ok || ifst.Else == nil {
		t.Fatalf("ожидался if с else, получено %+v", fw.Body.Stmts[0])
	}
	if _, ok := ifst.Else.(*ast.BlockStmt); !ok {
		t.Fatalf("ожидался else-блок, получено %T", ifst.Else)
	}
	fc, ok := body[2].(*ast.ForStmt)
	if !ok || fc.Init == nil || fc.Cond == nil || fc.Post == nil {
		t.Fatalf("ожидался for-clause со всеми частями, получено %+v", body[2])
	}
	if _, ok := fc.Body.Stmts[0].(*ast.BreakStmt); !ok {
		t.Fatalf("ожидался break в теле, получено %+v", fc.Body.Stmts[0])
	}
}

func TestParseForInfiniteAndElseIf(t *testing.T) {
	src := "package main\nfunc main() {\n" +
		"  for {\n    continue;\n  }\n" +
		"  if false {\n    fmt.Println(\"a\");\n" +
		"  } else if true {\n    fmt.Println(\"b\");\n" +
		"  } else {\n    fmt.Println(\"c\");\n  }\n" +
		"}\n"
	prog, bag := Parse("t.go", src)
	if bag.HasErrors() {
		t.Fatalf("неожиданные ошибки: %v", bag.Errors())
	}
	body := prog.Funcs[0].Body

	inf, ok := body[0].(*ast.ForStmt)
	if !ok || inf.Init != nil || inf.Cond != nil || inf.Post != nil {
		t.Fatalf("ожидался for-infinite, получено %+v", body[0])
	}
	if _, ok := inf.Body.Stmts[0].(*ast.ContinueStmt); !ok {
		t.Fatalf("ожидался continue в теле for, получено %+v", inf.Body.Stmts[0])
	}

	outer, ok := body[1].(*ast.IfStmt)
	if !ok {
		t.Fatalf("ожидался if, получено %+v", body[1])
	}
	elif, ok := outer.Else.(*ast.IfStmt)
	if !ok {
		t.Fatalf("ожидался else-if (*ast.IfStmt), получено %T", outer.Else)
	}
	if _, ok := elif.Else.(*ast.BlockStmt); !ok {
		t.Fatalf("ожидался финальный else-блок, получено %T", elif.Else)
	}
}
