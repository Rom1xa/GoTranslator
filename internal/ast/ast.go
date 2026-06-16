package ast

type FieldDecl struct {
	Name string
	Type string
}

type StructType struct {
	Name   string
	Fields []FieldDecl
}

type MethodDecl struct {
	Name        string
	Params      []Param
	ReturnTypes []string
}

type InterfaceType struct {
	Name    string
	Methods []MethodDecl
}

type Program struct {
	Package    string
	Types      []*StructType
	Interfaces []*InterfaceType
	Funcs      []*FuncDecl
}

type Param struct {
	Name string
	Type string
}

type FuncDecl struct {
	Name        string
	Params      []Param
	ReturnTypes []string // пустой срез => void
	Body        []Stmt
}

type Stmt interface{ stmtNode() }

// PrintStmt: fmt.Println(Args...) или fmt.Printf(Args...).
type PrintStmt struct {
	Fn   string // "Println" | "Printf"
	Args []Expr
	Line int
	Col  int
}

func (*PrintStmt) stmtNode() {}

// GoStmt: go foo(args...) или go fmt.Println/Printf(args...).
type GoStmt struct {
	Name    string
	PrintFn string
	Args    []Expr
	Line    int
	Col     int
}

func (*GoStmt) stmtNode() {}

// VarDecl: var Name Type [= Init]
type VarDecl struct {
	Name string
	Type string // "int" | "float64" | "string" | "bool"
	Init Expr   // nil => нулевое значение типа
	Line int
	Col  int
}

// ShortVarDecl: Name := Value
type ShortVarDecl struct {
	Name  string
	Value Expr
	Line  int
	Col   int
}

// AssignStmt: Name = Value
type AssignStmt struct {
	Name  string
	Value Expr
	Line  int
	Col   int
}

func (*VarDecl) stmtNode()      {}
func (*ShortVarDecl) stmtNode() {}
func (*AssignStmt) stmtNode()   {}

type Expr interface{ exprNode() }

type (
	StringLit struct{ Value string }
	IntLit    struct{ Value int64 }
)

func (*StringLit) exprNode() {}
func (*IntLit) exprNode()    {}

// Ident — обращение к переменной по имени.
type Ident struct {
	Name string
	Line int
	Col  int
}

// Left Op Right (Op — текст оператора: "+","*","==","&&",...).
type BinaryExpr struct {
	Op          string
	Left, Right Expr
	Line        int
	Col         int
}

// Op — "-" | "!".
type UnaryExpr struct {
	Op      string
	Operand Expr
	Line    int
	Col     int
}

type (
	FloatLit struct{ Value float64 }
	BoolLit  struct{ Value bool }
)

func (*Ident) exprNode()      {}
func (*BinaryExpr) exprNode() {}
func (*UnaryExpr) exprNode()  {}
func (*FloatLit) exprNode()   {}
func (*BoolLit) exprNode()    {}

// Последовательность операторов в своей области видимости.
type BlockStmt struct {
	Stmts []Stmt
	Line  int
	Col   int
}

// if Cond Then [else Else]. Else — nil | *IfStmt (else if) | *BlockStmt (else).
type IfStmt struct {
	Cond Expr
	Then *BlockStmt
	Else Stmt
	Line int
	Col  int
}

// for {}                      => Init=nil, Cond=nil, Post=nil
// for cond {}                 => только Cond
// for init; cond; post {}     => любые из трёх
type ForStmt struct {
	Init Stmt // *ShortVarDecl | *AssignStmt | nil
	Cond Expr
	Post Stmt // *ShortVarDecl | *AssignStmt | nil
	Body *BlockStmt
	Line int
	Col  int
}

type BreakStmt struct {
	Line int
	Col  int
}

type ContinueStmt struct {
	Line int
	Col  int
}

func (*BlockStmt) stmtNode()    {}
func (*IfStmt) stmtNode()       {}
func (*ForStmt) stmtNode()      {}
func (*BreakStmt) stmtNode()    {}
func (*ContinueStmt) stmtNode() {}

// return [expr, ...];
type ReturnStmt struct {
	Values []Expr
	Line   int
	Col    int
}

// foo(args...);
type CallStmt struct {
	Name string
	Args []Expr
	Line int
	Col  int
}

// a, b := foo(args...);
type MultiShortVarDecl struct {
	Names []string
	Func  string
	Args  []Expr
	Line  int
	Col   int
}

// foo(args...)
type CallExpr struct {
	Name string
	Args []Expr
	Line int
	Col  int
}

func (*ReturnStmt) stmtNode()        {}
func (*CallStmt) stmtNode()          {}
func (*MultiShortVarDecl) stmtNode() {}
func (*CallExpr) exprNode()          {}

// Object.Field = Value;
type FieldAssignStmt struct {
	Object string
	Field  string
	Value  Expr
	Line   int
	Col    int
}

// одна пара поле:значение в литерале структуры.
type FieldInit struct {
	Name  string
	Value Expr
}

// TypeName{field: val, ...}
type StructLit struct {
	TypeName string
	Fields   []FieldInit
	Line     int
	Col      int
}

// Object.Field
type SelectorExpr struct {
	Object Expr
	Field  string
	Line   int
	Col    int
}

func (*FieldAssignStmt) stmtNode() {}
func (*StructLit) exprNode()       {}
func (*SelectorExpr) exprNode()    {}

// []ElemType{elem, elem, ...}
type SliceLit struct {
	ElemType string // "int", "[]int", "Point", ...
	Elems    []Expr
	Line     int
	Col      int
}

// slice[index]
type IndexExpr struct {
	Object Expr
	Index  Expr
	Line   int
	Col    int
}

// name[index] = value;
type IndexAssignStmt struct {
	Object string
	Index  Expr
	Value  Expr
	Line   int
	Col    int
}

// ForRangeStmt: for [Key,] [Val] := range Slice { ... }
// Key == "" или "_"   не объявляется в scope.
// Val == "" или "_"   не объявляется в scope.
type ForRangeStmt struct {
	Key   string
	Val   string
	Slice Expr
	Body  *BlockStmt
	Line  int
	Col   int
}

func (*SliceLit) exprNode()        {}
func (*IndexExpr) exprNode()       {}
func (*IndexAssignStmt) stmtNode() {}
func (*ForRangeStmt) stmtNode()    {}
