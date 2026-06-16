package builder

import (
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"

	"github.com/rom1xa/kyrsoTranslatorGo/internal/ast"
	"github.com/rom1xa/kyrsoTranslatorGo/internal/diag"
	"github.com/rom1xa/kyrsoTranslatorGo/internal/gen"
)

type errListener struct {
	*antlr.DefaultErrorListener
	bag *diag.Bag
}

func (l *errListener) SyntaxError(_ antlr.Recognizer, _ any,
	line, col int, msg string, _ antlr.RecognitionException,
) {
	l.bag.Add(diag.Error{Line: line, Col: col, Level: diag.Syntax, Msg: msg})
}

func Parse(name, src string) (*ast.Program, *diag.Bag) {
	bag := &diag.Bag{}
	listener := &errListener{bag: bag}

	input := antlr.NewInputStream(src)
	lexer := gen.NewGoSubsetLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(listener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := gen.NewGoSubsetParser(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(listener)

	tree := parser.Program()
	if bag.HasErrors() {
		return nil, bag
	}
	return buildProgram(tree), bag
}

func buildProgram(ctx gen.IProgramContext) *ast.Program {
	p := &ast.Program{Package: ctx.IDENT().GetText()}
	for _, tl := range ctx.AllTopLevel() {
		if fd := tl.FuncDecl(); fd != nil {
			p.Funcs = append(p.Funcs, buildFunc(fd))
		} else if td := tl.TypeDecl(); td != nil {
			st, it := buildTypeDecl(td)
			if st != nil {
				p.Types = append(p.Types, st)
			}
			if it != nil {
				p.Interfaces = append(p.Interfaces, it)
			}
		}
	}
	return p
}

func buildTypeDecl(ctx gen.ITypeDeclContext) (*ast.StructType, *ast.InterfaceType) {
	name := ctx.IDENT().GetText()
	td := ctx.TypeDef()
	if fields := td.AllFieldDecl(); len(fields) > 0 || strings.HasPrefix(td.GetText(), "struct") {
		st := &ast.StructType{Name: name}
		for _, fd := range fields {
			st.Fields = append(st.Fields, ast.FieldDecl{
				Name: fd.IDENT().GetText(),
				Type: fd.Type_().GetText(),
			})
		}
		return st, nil
	}
	it := &ast.InterfaceType{Name: name}
	for _, md := range td.AllMethodDecl() {
		it.Methods = append(it.Methods, buildMethodDecl(md))
	}
	return nil, it
}

func buildMethodDecl(ctx gen.IMethodDeclContext) ast.MethodDecl {
	m := ast.MethodDecl{Name: ctx.IDENT().GetText()}
	if pl := ctx.ParamList(); pl != nil {
		for _, p := range pl.AllParam() {
			m.Params = append(m.Params, ast.Param{
				Name: p.IDENT().GetText(),
				Type: p.Type_().GetText(),
			})
		}
	}
	if rs := ctx.ReturnSpec(); rs != nil {
		for _, t := range rs.AllType_() {
			m.ReturnTypes = append(m.ReturnTypes, t.GetText())
		}
	}
	return m
}

func buildFieldAssign(ctx gen.IFieldAssignStmtContext) ast.Stmt {
	idents := ctx.AllIDENT()
	return &ast.FieldAssignStmt{
		Object: idents[0].GetText(),
		Field:  idents[1].GetText(),
		Value:  buildExpr(ctx.Expr()),
		Line:   ctx.GetStart().GetLine(),
		Col:    ctx.GetStart().GetColumn(),
	}
}

func buildIndexAssign(ctx gen.IIndexAssignStmtContext) ast.Stmt {
	return &ast.IndexAssignStmt{
		Object: ctx.IDENT().GetText(),
		Index:  buildExpr(ctx.Expr(0)),
		Value:  buildExpr(ctx.Expr(1)),
		Line:   ctx.GetStart().GetLine(),
		Col:    ctx.GetStart().GetColumn(),
	}
}

func buildFunc(ctx gen.IFuncDeclContext) *ast.FuncDecl {
	f := &ast.FuncDecl{Name: ctx.IDENT().GetText()}
	if pl := ctx.ParamList(); pl != nil {
		for _, p := range pl.AllParam() {
			f.Params = append(f.Params, ast.Param{
				Name: p.IDENT().GetText(),
				Type: p.Type_().GetText(),
			})
		}
	}
	if rs := ctx.ReturnSpec(); rs != nil {
		for _, t := range rs.AllType_() {
			f.ReturnTypes = append(f.ReturnTypes, t.GetText())
		}
	}
	for _, st := range ctx.Block().AllStatement() {
		f.Body = append(f.Body, buildStmt(st))
	}
	return f
}

func buildStmt(ctx gen.IStatementContext) ast.Stmt {
	switch {
	case ctx.VarDecl() != nil:
		return buildVarDecl(ctx.VarDecl())
	case ctx.GoStmt() != nil:
		return buildGoStmt(ctx.GoStmt())
	case ctx.ShortVarDecl() != nil:
		return buildShortVarDecl(ctx.ShortVarDecl())
	case ctx.FieldAssignStmt() != nil:
		return buildFieldAssign(ctx.FieldAssignStmt())
	case ctx.IndexAssignStmt() != nil:
		return buildIndexAssign(ctx.IndexAssignStmt())
	case ctx.AssignStmt() != nil:
		return buildAssign(ctx.AssignStmt())
	case ctx.IfStmt() != nil:
		return buildIf(ctx.IfStmt())
	case ctx.ForStmt() != nil:
		return buildFor(ctx.ForStmt())
	case ctx.BreakStmt() != nil:
		return &ast.BreakStmt{Line: ctx.GetStart().GetLine(), Col: ctx.GetStart().GetColumn()}
	case ctx.ContinueStmt() != nil:
		return &ast.ContinueStmt{Line: ctx.GetStart().GetLine(), Col: ctx.GetStart().GetColumn()}
	case ctx.ReturnStmt() != nil:
		return buildReturn(ctx.ReturnStmt())
	case ctx.CallStmt() != nil:
		return buildCallStmt(ctx.CallStmt())
	case ctx.MultiShortDecl() != nil:
		return buildMultiShortDecl(ctx.MultiShortDecl())
	case ctx.Block() != nil:
		return buildBlock(ctx.Block())
	case ctx.PrintStmt() != nil:
		return buildPrint(ctx.PrintStmt())
	default:
		panic("buildStmt: неизвестный оператор")
	}
}

func buildGoStmt(ctx gen.IGoStmtContext) ast.Stmt {
	gs := &ast.GoStmt{Line: ctx.GetStart().GetLine(), Col: ctx.GetStart().GetColumn()}
	if ident := ctx.IDENT(); ident != nil {
		gs.Name = ident.GetText()
	} else if strings.HasPrefix(ctx.GetText(), "gofmt.Printf") {
		gs.PrintFn = "Printf"
	} else {
		gs.PrintFn = "Println"
	}
	if list := ctx.ExprList(); list != nil {
		for _, e := range list.AllExpr() {
			gs.Args = append(gs.Args, buildExpr(e))
		}
	}
	return gs
}

func buildBlock(ctx gen.IBlockContext) *ast.BlockStmt {
	b := &ast.BlockStmt{Line: ctx.GetStart().GetLine(), Col: ctx.GetStart().GetColumn()}
	for _, st := range ctx.AllStatement() {
		b.Stmts = append(b.Stmts, buildStmt(st))
	}
	return b
}

func buildIf(ctx gen.IIfStmtContext) ast.Stmt {
	s := &ast.IfStmt{
		Cond: buildExpr(ctx.Expr()),
		Then: buildBlock(ctx.Block()),
		Line: ctx.GetStart().GetLine(),
		Col:  ctx.GetStart().GetColumn(),
	}
	if ep := ctx.ElsePart(); ep != nil {
		switch {
		case ep.IfStmt() != nil:
			s.Else = buildIf(ep.IfStmt())
		case ep.Block() != nil:
			s.Else = buildBlock(ep.Block())
		default:
			panic("buildIf: неизвестная форма else")
		}
	}
	return s
}

func buildFor(ctx gen.IForStmtContext) ast.Stmt {
	fs := &ast.ForStmt{Line: ctx.GetStart().GetLine(), Col: ctx.GetStart().GetColumn()}
	switch c := ctx.(type) {
	case *gen.ForInfiniteContext:
		fs.Body = buildBlock(c.Block())
	case *gen.ForWhileContext:
		fs.Cond = buildExpr(c.Expr())
		fs.Body = buildBlock(c.Block())
	case *gen.ForClauseContext:
		if c.GetInit() != nil {
			fs.Init = buildSimple(c.GetInit())
		}
		if c.GetCond() != nil {
			fs.Cond = buildExpr(c.GetCond())
		}
		if c.GetPost() != nil {
			fs.Post = buildSimple(c.GetPost())
		}
		fs.Body = buildBlock(c.Block())
	case *gen.ForRangeFullContext:
		idents := c.AllIDENT()
		return &ast.ForRangeStmt{
			Key:   idents[0].GetText(),
			Val:   idents[1].GetText(),
			Slice: buildExpr(c.Expr()),
			Body:  buildBlock(c.Block()),
			Line:  c.GetStart().GetLine(),
			Col:   c.GetStart().GetColumn(),
		}
	case *gen.ForRangeIndexContext:
		return &ast.ForRangeStmt{
			Key:   c.IDENT().GetText(),
			Slice: buildExpr(c.Expr()),
			Body:  buildBlock(c.Block()),
			Line:  c.GetStart().GetLine(),
			Col:   c.GetStart().GetColumn(),
		}
	default:
		panic("buildFor: неизвестная форма for")
	}
	return fs
}

func buildSimple(ctx gen.ISimpleStmtContext) ast.Stmt {
	line := ctx.GetStart().GetLine()
	col := ctx.GetStart().GetColumn()
	switch c := ctx.(type) {
	case *gen.SimpleShortContext:
		return &ast.ShortVarDecl{Name: c.IDENT().GetText(), Value: buildExpr(c.Expr()), Line: line, Col: col}
	case *gen.SimpleAssignContext:
		return &ast.AssignStmt{Name: c.IDENT().GetText(), Value: buildExpr(c.Expr()), Line: line, Col: col}
	default:
		panic("buildSimple: неизвестный простой оператор")
	}
}

func buildVarDecl(ctx gen.IVarDeclContext) ast.Stmt {
	vd := &ast.VarDecl{
		Name: ctx.IDENT().GetText(),
		Type: ctx.Type_().GetText(),
		Line: ctx.GetStart().GetLine(),
		Col:  ctx.GetStart().GetColumn(),
	}
	if e := ctx.Expr(); e != nil {
		vd.Init = buildExpr(e)
	}
	return vd
}

func buildShortVarDecl(ctx gen.IShortVarDeclContext) ast.Stmt {
	return &ast.ShortVarDecl{
		Name:  ctx.IDENT().GetText(),
		Value: buildExpr(ctx.Expr()),
		Line:  ctx.GetStart().GetLine(),
		Col:   ctx.GetStart().GetColumn(),
	}
}

func buildAssign(ctx gen.IAssignStmtContext) ast.Stmt {
	return &ast.AssignStmt{
		Name:  ctx.IDENT().GetText(),
		Value: buildExpr(ctx.Expr()),
		Line:  ctx.GetStart().GetLine(),
		Col:   ctx.GetStart().GetColumn(),
	}
}

func buildReturn(ctx gen.IReturnStmtContext) ast.Stmt {
	rs := &ast.ReturnStmt{Line: ctx.GetStart().GetLine(), Col: ctx.GetStart().GetColumn()}
	if list := ctx.ExprList(); list != nil {
		for _, e := range list.AllExpr() {
			rs.Values = append(rs.Values, buildExpr(e))
		}
	}
	return rs
}

func buildCallStmt(ctx gen.ICallStmtContext) ast.Stmt {
	cs := &ast.CallStmt{
		Name: ctx.IDENT().GetText(),
		Line: ctx.GetStart().GetLine(),
		Col:  ctx.GetStart().GetColumn(),
	}
	if list := ctx.ExprList(); list != nil {
		for _, e := range list.AllExpr() {
			cs.Args = append(cs.Args, buildExpr(e))
		}
	}
	return cs
}

func buildMultiShortDecl(ctx gen.IMultiShortDeclContext) ast.Stmt {
	ms := &ast.MultiShortVarDecl{
		Func: ctx.IDENT(len(ctx.AllIDENT()) - 1).GetText(),
		Line: ctx.GetStart().GetLine(),
		Col:  ctx.GetStart().GetColumn(),
	}
	for k := 0; k < len(ctx.AllIDENT())-1; k++ {
		ms.Names = append(ms.Names, ctx.IDENT(k).GetText())
	}
	if list := ctx.ExprList(); list != nil {
		for _, e := range list.AllExpr() {
			ms.Args = append(ms.Args, buildExpr(e))
		}
	}
	return ms
}

func buildPrint(ctx gen.IPrintStmtContext) ast.Stmt {
	fn := "Println"
	if strings.HasPrefix(ctx.GetText(), "fmt.Printf") {
		fn = "Printf"
	}
	ps := &ast.PrintStmt{
		Fn:   fn,
		Line: ctx.GetStart().GetLine(),
		Col:  ctx.GetStart().GetColumn(),
	}
	if list := ctx.ExprList(); list != nil {
		for _, e := range list.AllExpr() {
			ps.Args = append(ps.Args, buildExpr(e))
		}
	}
	return ps
}

func buildExpr(ctx gen.IExprContext) ast.Expr {
	line := ctx.GetStart().GetLine()
	col := ctx.GetStart().GetColumn()
	switch c := ctx.(type) {
	case *gen.ParenExprContext:
		return buildExpr(c.Expr())
	case *gen.UnaryExprContext:
		return &ast.UnaryExpr{Op: c.GetOp().GetText(), Operand: buildExpr(c.Expr()), Line: line, Col: col}
	case *gen.MulExprContext:
		return binary(c.GetOp().GetText(), c.Expr(0), c.Expr(1), line, col)
	case *gen.AddExprContext:
		return binary(c.GetOp().GetText(), c.Expr(0), c.Expr(1), line, col)
	case *gen.RelExprContext:
		return binary(c.GetOp().GetText(), c.Expr(0), c.Expr(1), line, col)
	case *gen.EqExprContext:
		return binary(c.GetOp().GetText(), c.Expr(0), c.Expr(1), line, col)
	case *gen.AndExprContext:
		return binary("&&", c.Expr(0), c.Expr(1), line, col)
	case *gen.OrExprContext:
		return binary("||", c.Expr(0), c.Expr(1), line, col)
	case *gen.IntExprContext:
		return &ast.IntLit{Value: atoi(c.INT().GetText())}
	case *gen.FloatExprContext:
		return &ast.FloatLit{Value: atof(c.FLOAT().GetText())}
	case *gen.StringExprContext:
		return &ast.StringLit{Value: unquote(c.STRING().GetText())}
	case *gen.TrueExprContext:
		return &ast.BoolLit{Value: true}
	case *gen.FalseExprContext:
		return &ast.BoolLit{Value: false}
	case *gen.CallExprContext:
		ce := &ast.CallExpr{Name: c.IDENT().GetText(), Line: line, Col: col}
		if list := c.ExprList(); list != nil {
			for _, e := range list.AllExpr() {
				ce.Args = append(ce.Args, buildExpr(e))
			}
		}
		return ce
	case *gen.SelectorExprContext:
		return &ast.SelectorExpr{
			Object: buildExpr(c.Expr()),
			Field:  c.IDENT().GetText(),
			Line:   line,
			Col:    col,
		}
	case *gen.StructLitExprContext:
		sl := &ast.StructLit{
			TypeName: c.IDENT().GetText(),
			Line:     line,
			Col:      col,
		}
		if list := c.FieldInitList(); list != nil {
			for _, fi := range list.AllFieldInit() {
				sl.Fields = append(sl.Fields, ast.FieldInit{
					Name:  fi.IDENT().GetText(),
					Value: buildExpr(fi.Expr()),
				})
			}
		}
		return sl
	case *gen.SliceLitExprContext:
		sl := &ast.SliceLit{
			ElemType: c.Type_().GetText(),
			Line:     line,
			Col:      col,
		}
		if list := c.ExprList(); list != nil {
			for _, e := range list.AllExpr() {
				sl.Elems = append(sl.Elems, buildExpr(e))
			}
		}
		return sl
	case *gen.IndexExprContext:
		return &ast.IndexExpr{
			Object: buildExpr(c.Expr(0)),
			Index:  buildExpr(c.Expr(1)),
			Line:   line,
			Col:    col,
		}
	case *gen.IdentExprContext:
		return &ast.Ident{Name: c.IDENT().GetText(), Line: line, Col: col}
	default:
		panic("buildExpr: неизвестное выражение")
	}
}

func binary(op string, l, r gen.IExprContext, line, col int) ast.Expr {
	return &ast.BinaryExpr{Op: op, Left: buildExpr(l), Right: buildExpr(r), Line: line, Col: col}
}

func unquote(s string) string {
	if u, err := strconv.Unquote(s); err == nil {
		return u
	}
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}

func atoi(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}

func atof(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
