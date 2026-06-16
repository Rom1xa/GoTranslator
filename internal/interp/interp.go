package interp

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/rom1xa/kyrsoTranslatorGo/internal/ast"
	"github.com/rom1xa/kyrsoTranslatorGo/internal/diag"
	"github.com/rom1xa/kyrsoTranslatorGo/internal/value"
)

var (
	errBreak    = errors.New("break")
	errContinue = errors.New("continue")
)

// сигнал возврата из функции
type errReturn struct {
	values []value.Value
}

func (e errReturn) Error() string { return "return" }

const maxIterations = 10_000_000

type scope struct {
	vars        map[string]value.Value
	parent      *scope
	isFrameRoot bool
}

func newScope(parent *scope) *scope {
	return &scope{vars: map[string]value.Value{}, parent: parent}
}

func (s *scope) get(name string) (value.Value, bool) {
	for c := s; c != nil; c = c.parent {
		if v, ok := c.vars[name]; ok {
			return v, true
		}
		if c.isFrameRoot {
			break
		}
	}
	return nil, false
}

func (s *scope) setExisting(name string, v value.Value) bool {
	for c := s; c != nil; c = c.parent {
		if _, ok := c.vars[name]; ok {
			c.vars[name] = v
			return true
		}
		if c.isFrameRoot {
			break
		}
	}
	return false
}

func (s *scope) hasLocal(name string) bool          { _, ok := s.vars[name]; return ok }
func (s *scope) declare(name string, v value.Value) { s.vars[name] = v }

type Interp struct {
	out        io.Writer
	cur        *scope
	funcs      map[string]*ast.FuncDecl
	types      map[string]*ast.StructType
	interfaces map[string]*ast.InterfaceType
	async      *asyncState
}

type asyncState struct {
	wg    sync.WaitGroup
	mu    sync.Mutex
	outMu sync.Mutex
	err   error
}

func New(out io.Writer) *Interp {
	return &Interp{
		out:        out,
		cur:        newScope(nil),
		funcs:      map[string]*ast.FuncDecl{},
		types:      map[string]*ast.StructType{},
		interfaces: map[string]*ast.InterfaceType{},
		async:      &asyncState{},
	}
}

func (i *Interp) push() { i.cur = newScope(i.cur) }
func (i *Interp) pop()  { i.cur = i.cur.parent }

// Run исполняет программр: егистрирует типы и функции и вызывает main.
func (i *Interp) Run(p *ast.Program) error {
	for _, t := range p.Types {
		i.types[t.Name] = t
	}
	for _, t := range p.Interfaces {
		i.interfaces[t.Name] = t
	}
	for _, f := range p.Funcs {
		i.funcs[f.Name] = f
	}
	main, ok := i.funcs["main"]
	if !ok {
		return diag.Error{Level: diag.Semantic, Msg: "функция main не найдена"}
	}
	_, err := i.callFunc(main, nil, 0, 0)
	i.async.wg.Wait()
	if err != nil {
		return err
	}
	i.async.mu.Lock()
	defer i.async.mu.Unlock()
	return i.async.err
}

func (i *Interp) callFunc(fn *ast.FuncDecl, args []value.Value, line, col int) ([]value.Value, error) {
	if len(args) != len(fn.Params) {
		return nil, semErr(line, col, fmt.Sprintf(
			"функция %s: ожидалось %d аргументов, получено %d",
			fn.Name, len(fn.Params), len(args),
		))
	}

	saved := i.cur
	i.cur = &scope{vars: map[string]value.Value{}, isFrameRoot: true}
	defer func() { i.cur = saved }()

	for k, p := range fn.Params {
		if !i.typeMatches(p.Type, args[k]) {
			return nil, semErr(line, col, fmt.Sprintf(
				"параметр %s: ожидался %s, получено %s",
				p.Name, p.Type, typeNameOf(args[k]),
			))
		}
		i.cur.declare(p.Name, i.wrapForType(p.Type, args[k]))
	}

	for _, st := range fn.Body {
		if err := i.execStmt(st); err != nil {
			var ret errReturn
			if errors.As(err, &ret) {
				if err2 := i.checkReturn(fn, ret, line, col); err2 != nil {
					return nil, err2
				}
				return ret.values, nil
			}
			if errors.Is(err, errBreak) || errors.Is(err, errContinue) {
				return nil, semErr(line, col, "break/continue вне цикла")
			}
			return nil, err
		}
	}

	if len(fn.ReturnTypes) > 0 {
		return nil, semErr(line, col, fmt.Sprintf(
			"функция %s: не все пути возвращают значение", fn.Name,
		))
	}
	return nil, nil
}

func (i *Interp) checkReturn(fn *ast.FuncDecl, ret errReturn, line, col int) error {
	if len(ret.values) != len(fn.ReturnTypes) {
		return semErr(line, col, fmt.Sprintf(
			"функция %s: ожидается %d возвращаемых значений, получено %d",
			fn.Name, len(fn.ReturnTypes), len(ret.values),
		))
	}
	for k, t := range fn.ReturnTypes {
		if !i.typeMatches(t, ret.values[k]) {
			return semErr(line, col, fmt.Sprintf(
				"функция %s: возврат %d: ожидался %s, получено %s",
				fn.Name, k+1, t, typeNameOf(ret.values[k]),
			))
		}
		ret.values[k] = i.wrapForType(t, ret.values[k])
	}
	return nil
}

func (i *Interp) execStmt(s ast.Stmt) error {
	switch st := s.(type) {
	case *ast.VarDecl:
		return i.execVarDecl(st)
	case *ast.ShortVarDecl:
		if i.cur.hasLocal(st.Name) {
			return semErr(st.Line, st.Col, "повторное объявление: "+st.Name)
		}
		v, err := i.eval(st.Value)
		if err != nil {
			return err
		}
		i.cur.declare(st.Name, v)
		return nil
	case *ast.AssignStmt:
		old, ok := i.cur.get(st.Name)
		if !ok {
			return semErr(st.Line, st.Col, "undefined: "+st.Name)
		}
		v, err := i.eval(st.Value)
		if err != nil {
			return err
		}
		targetType := typeNameOf(old)
		if !i.typeMatches(targetType, v) {
			return semErr(st.Line, st.Col, fmt.Sprintf(
				"несоответствие типов при присваивании %s: %s = %s",
				st.Name, targetType, typeNameOf(v),
			))
		}
		if !i.cur.setExisting(st.Name, i.wrapForType(targetType, v)) {
			return semErr(st.Line, st.Col, "undefined: "+st.Name)
		}
		return nil
	case *ast.BlockStmt:
		return i.execBlock(st)
	case *ast.IfStmt:
		return i.execIf(st)
	case *ast.ForStmt:
		return i.execFor(st)
	case *ast.BreakStmt:
		return errBreak
	case *ast.ContinueStmt:
		return errContinue
	case *ast.ReturnStmt:
		return i.execReturn(st)
	case *ast.CallStmt:
		return i.execCallStmt(st)
	case *ast.MultiShortVarDecl:
		return i.execMultiShortDecl(st)
	case *ast.PrintStmt:
		return i.execPrint(st)
	case *ast.GoStmt:
		return i.execGo(st)
	case *ast.FieldAssignStmt:
		return i.execFieldAssign(st)
	case *ast.IndexAssignStmt:
		return i.execIndexAssign(st)
	case *ast.ForRangeStmt:
		return i.execForRange(st)
	default:
		return diag.Error{Level: diag.Runtime, Msg: fmt.Sprintf("неизвестный оператор %T", s)}
	}
}

func (i *Interp) execVarDecl(vd *ast.VarDecl) error {
	if i.cur.hasLocal(vd.Name) {
		return semErr(vd.Line, vd.Col, "повторное объявление: "+vd.Name)
	}
	if vd.Init == nil {
		v, err := i.zeroValue(vd.Type, vd.Line, vd.Col)
		if err != nil {
			return err
		}
		i.cur.declare(vd.Name, v)
		return nil
	}
	v, err := i.eval(vd.Init)
	if err != nil {
		return err
	}
	if !i.typeMatches(vd.Type, v) {
		return diag.Error{
			Line: vd.Line, Col: vd.Col, Level: diag.Semantic,
			Msg: fmt.Sprintf("несоответствие типов: %s = %s", vd.Type, typeNameOf(v)),
		}
	}
	i.cur.declare(vd.Name, i.wrapForType(vd.Type, v))
	return nil
}

func (i *Interp) execBlock(b *ast.BlockStmt) error {
	i.push()
	defer i.pop()
	for _, st := range b.Stmts {
		if err := i.execStmt(st); err != nil {
			return err
		}
	}
	return nil
}

func (i *Interp) execIf(s *ast.IfStmt) error {
	c, err := i.eval(s.Cond)
	if err != nil {
		return err
	}
	cb, ok := c.(value.Bool)
	if !ok {
		return semErr(s.Line, s.Col, "условие if должно быть bool")
	}
	if bool(cb) {
		return i.execBlock(s.Then)
	}
	if s.Else != nil {
		return i.execStmt(s.Else)
	}
	return nil
}

func (i *Interp) execFor(fs *ast.ForStmt) error {
	i.push()
	defer i.pop()

	if fs.Init != nil {
		if err := i.execStmt(fs.Init); err != nil {
			return err
		}
	}

	iter := 0
	for {
		if fs.Cond != nil {
			c, err := i.eval(fs.Cond)
			if err != nil {
				return err
			}
			cb, ok := c.(value.Bool)
			if !ok {
				return semErr(fs.Line, fs.Col, "условие for должно быть bool")
			}
			if !bool(cb) {
				break
			}
		}

		iter++
		if iter > maxIterations {
			return rtErr(fs.Line, fs.Col, "превышен лимит итераций (возможный бесконечный цикл)")
		}

		if err := i.execBlock(fs.Body); err != nil {
			if errors.Is(err, errBreak) {
				break
			}
			if !errors.Is(err, errContinue) {
				return err
			}
		}

		if fs.Post != nil {
			if err := i.execStmt(fs.Post); err != nil {
				return err
			}
		}
	}
	return nil
}

func (i *Interp) zeroValue(t string, line, col int) (value.Value, error) {
	switch t {
	case "int":
		return value.Int(0), nil
	case "float64":
		return value.Float(0), nil
	case "string":
		return value.String(""), nil
	case "bool":
		return value.Bool(false), nil
	default:
		if _, ok := i.interfaces[t]; ok {
			return value.NewInterface(t, nil), nil
		}
		if strings.HasPrefix(t, "[]") {
			return value.NewSlice(t[2:]), nil
		}
		st, ok := i.types[t]
		if !ok {
			return nil, semErr(line, col, "неизвестный тип: "+t)
		}
		s := value.NewStruct(t)
		for _, fd := range st.Fields {
			fv, err := i.zeroValue(fd.Type, line, col)
			if err != nil {
				return nil, err
			}
			s.Set(fd.Name, fv)
		}
		return s, nil
	}
}

func typeNameOf(v value.Value) string {
	switch n := v.(type) {
	case value.Int:
		return "int"
	case value.Float:
		return "float64"
	case value.String:
		return "string"
	case value.Bool:
		return "bool"
	case *value.Interface:
		return n.TypeName
	case *value.Struct:
		return n.TypeName
	case *value.Slice:
		return "[]" + n.ElemType
	}
	return "unknown"
}

func (i *Interp) typeMatches(t string, v value.Value) bool {
	if iv, ok := v.(*value.Interface); ok {
		if iv.Value == nil {
			return t == iv.TypeName
		}
		if t == iv.TypeName {
			return true
		}
		v = iv.Value
	}
	if i.isInterfaceType(t) {
		return true
	}
	switch n := v.(type) {
	case value.Int:
		return t == "int"
	case value.Float:
		return t == "float64"
	case value.String:
		return t == "string"
	case value.Bool:
		return t == "bool"
	case *value.Struct:
		return t == n.TypeName
	case *value.Slice:
		return t == "[]"+n.ElemType
	}
	return false
}

func (i *Interp) isInterfaceType(t string) bool {
	_, ok := i.interfaces[t]
	return ok
}

func (i *Interp) wrapForType(t string, v value.Value) value.Value {
	if iv, ok := v.(*value.Interface); ok && iv.TypeName == t {
		return iv
	}
	if i.isInterfaceType(t) {
		return value.NewInterface(t, unwrapInterface(v))
	}
	return v
}

func unwrapInterface(v value.Value) value.Value {
	if iv, ok := v.(*value.Interface); ok {
		return iv.Value
	}
	return v
}

func (i *Interp) execPrint(ps *ast.PrintStmt) error {
	args := make([]any, 0, len(ps.Args))
	for _, e := range ps.Args {
		v, err := i.eval(e)
		if err != nil {
			return err
		}
		args = append(args, v.Native())
	}
	i.async.outMu.Lock()
	defer i.async.outMu.Unlock()
	if ps.Fn == "Printf" {
		if len(args) == 0 {
			return diag.Error{
				Line: ps.Line, Col: ps.Col, Level: diag.Semantic,
				Msg: "Printf требует строку формата",
			}
		}
		format, ok := args[0].(string)
		if !ok {
			return diag.Error{
				Line: ps.Line, Col: ps.Col, Level: diag.Semantic,
				Msg: "первый аргумент Printf должен быть строкой",
			}
		}
		fmt.Fprintf(i.out, format, args[1:]...)
		return nil
	}
	fmt.Fprintln(i.out, args...)
	return nil
}

func (i *Interp) execGo(gs *ast.GoStmt) error {
	args, err := i.evalArgs(gs.Args, gs.Line, gs.Col)
	if err != nil {
		return err
	}
	if gs.PrintFn != "" {
		i.async.wg.Add(1)
		go func() {
			defer i.async.wg.Done()
			i.async.outMu.Lock()
			defer i.async.outMu.Unlock()
			native := make([]any, 0, len(args))
			for _, v := range args {
				native = append(native, v.Native())
			}
			if gs.PrintFn == "Printf" {
				if len(native) == 0 {
					i.setAsyncErr(diag.Error{Line: gs.Line, Col: gs.Col, Level: diag.Semantic, Msg: "Printf требует строку формата"})
					return
				}
				format, ok := native[0].(string)
				if !ok {
					i.setAsyncErr(diag.Error{Line: gs.Line, Col: gs.Col, Level: diag.Semantic, Msg: "первый аргумент Printf должен быть строкой"})
					return
				}
				fmt.Fprintf(i.out, format, native[1:]...)
				return
			}
			fmt.Fprintln(i.out, native...)
		}()
		return nil
	}
	fn, ok := i.funcs[gs.Name]
	if !ok {
		return semErr(gs.Line, gs.Col, "undefined function: "+gs.Name)
	}
	i.async.wg.Add(1)
	go func() {
		defer i.async.wg.Done()
		child := &Interp{
			out:        i.out,
			cur:        newScope(nil),
			funcs:      i.funcs,
			types:      i.types,
			interfaces: i.interfaces,
			async:      i.async,
		}
		if _, err := child.callFunc(fn, args, gs.Line, gs.Col); err != nil {
			child.setAsyncErr(err)
		}
	}()
	return nil
}

func (i *Interp) setAsyncErr(err error) {
	i.async.mu.Lock()
	defer i.async.mu.Unlock()
	if i.async.err == nil {
		i.async.err = err
	}
}

func (i *Interp) execReturn(rs *ast.ReturnStmt) error {
	values := make([]value.Value, 0, len(rs.Values))
	for _, e := range rs.Values {
		v, err := i.eval(e)
		if err != nil {
			return err
		}
		values = append(values, v)
	}
	return errReturn{values: values}
}

func (i *Interp) execCallStmt(cs *ast.CallStmt) error {
	if cs.Name == "len" || cs.Name == "append" {
		fakeExpr := &ast.CallExpr{Name: cs.Name, Args: cs.Args, Line: cs.Line, Col: cs.Col}
		_, err := i.evalCall(fakeExpr)
		return err
	}
	fn, ok := i.funcs[cs.Name]
	if !ok {
		return semErr(cs.Line, cs.Col, "undefined function: "+cs.Name)
	}
	args, err := i.evalArgs(cs.Args, cs.Line, cs.Col)
	if err != nil {
		return err
	}
	_, err = i.callFunc(fn, args, cs.Line, cs.Col)
	return err
}

func (i *Interp) execMultiShortDecl(ms *ast.MultiShortVarDecl) error {
	fn, ok := i.funcs[ms.Func]
	if !ok {
		return semErr(ms.Line, ms.Col, "undefined function: "+ms.Func)
	}
	args, err := i.evalArgs(ms.Args, ms.Line, ms.Col)
	if err != nil {
		return err
	}
	results, err := i.callFunc(fn, args, ms.Line, ms.Col)
	if err != nil {
		return err
	}
	if len(results) != len(ms.Names) {
		return semErr(ms.Line, ms.Col, fmt.Sprintf(
			"функция возвращает %d значений, ожидается %d", len(results), len(ms.Names),
		))
	}
	for k, name := range ms.Names {
		if i.cur.hasLocal(name) {
			return semErr(ms.Line, ms.Col, "повторное объявление: "+name)
		}
		i.cur.declare(name, results[k])
	}
	return nil
}

func (i *Interp) evalArgs(exprs []ast.Expr, line, col int) ([]value.Value, error) {
	args := make([]value.Value, 0, len(exprs))
	for _, e := range exprs {
		v, err := i.eval(e)
		if err != nil {
			return nil, err
		}
		args = append(args, v)
	}
	return args, nil
}

func (i *Interp) evalCall(ex *ast.CallExpr) (value.Value, error) {
	switch ex.Name {
	case "len":
		return i.evalBuiltinLen(ex)
	case "append":
		return i.evalBuiltinAppend(ex)
	}
	fn, ok := i.funcs[ex.Name]
	if !ok {
		return nil, semErr(ex.Line, ex.Col, "undefined function: "+ex.Name)
	}
	args, err := i.evalArgs(ex.Args, ex.Line, ex.Col)
	if err != nil {
		return nil, err
	}
	results, err := i.callFunc(fn, args, ex.Line, ex.Col)
	if err != nil {
		return nil, err
	}
	if len(results) != 1 {
		return nil, semErr(ex.Line, ex.Col, fmt.Sprintf(
			"функция %s возвращает %d значений в контексте одного значения",
			ex.Name, len(results),
		))
	}
	return results[0], nil
}

func (i *Interp) eval(e ast.Expr) (value.Value, error) {
	switch ex := e.(type) {
	case *ast.IntLit:
		return value.Int(ex.Value), nil
	case *ast.FloatLit:
		return value.Float(ex.Value), nil
	case *ast.StringLit:
		return value.String(ex.Value), nil
	case *ast.BoolLit:
		return value.Bool(ex.Value), nil
	case *ast.Ident:
		v, ok := i.cur.get(ex.Name)
		if !ok {
			return nil, semErr(ex.Line, ex.Col, "undefined: "+ex.Name)
		}
		return v, nil
	case *ast.CallExpr:
		return i.evalCall(ex)
	case *ast.UnaryExpr:
		return i.evalUnary(ex)
	case *ast.BinaryExpr:
		return i.evalBinary(ex)
	case *ast.StructLit:
		return i.evalStructLit(ex)
	case *ast.SelectorExpr:
		return i.evalSelector(ex)
	case *ast.SliceLit:
		return i.evalSliceLit(ex)
	case *ast.IndexExpr:
		return i.evalIndex(ex)
	default:
		return nil, diag.Error{Level: diag.Runtime, Msg: fmt.Sprintf("неизвестное выражение %T", e)}
	}
}

func (i *Interp) evalUnary(ex *ast.UnaryExpr) (value.Value, error) {
	v, err := i.eval(ex.Operand)
	if err != nil {
		return nil, err
	}
	switch ex.Op {
	case "-":
		switch n := v.(type) {
		case value.Int:
			return -n, nil
		case value.Float:
			return -n, nil
		}
		return nil, semErr(ex.Line, ex.Col, "оператор - требует число")
	case "!":
		if b, ok := v.(value.Bool); ok {
			return !b, nil
		}
		return nil, semErr(ex.Line, ex.Col, "оператор ! требует bool")
	}
	return nil, semErr(ex.Line, ex.Col, "неизвестный унарный оператор "+ex.Op)
}

func semErr(line, col int, msg string) error {
	return diag.Error{Line: line, Col: col, Level: diag.Semantic, Msg: msg}
}

func rtErr(line, col int, msg string) error {
	return diag.Error{Line: line, Col: col, Level: diag.Runtime, Msg: msg}
}

func (i *Interp) evalBinary(ex *ast.BinaryExpr) (value.Value, error) {
	if ex.Op == "&&" || ex.Op == "||" {
		l, err := i.eval(ex.Left)
		if err != nil {
			return nil, err
		}
		lb, ok := l.(value.Bool)
		if !ok {
			return nil, semErr(ex.Line, ex.Col, "логический оператор требует bool")
		}
		if ex.Op == "&&" && !bool(lb) {
			return value.Bool(false), nil
		}
		if ex.Op == "||" && bool(lb) {
			return value.Bool(true), nil
		}
		r, err := i.eval(ex.Right)
		if err != nil {
			return nil, err
		}
		rb, ok := r.(value.Bool)
		if !ok {
			return nil, semErr(ex.Line, ex.Col, "логический оператор требует bool")
		}
		return rb, nil
	}

	l, err := i.eval(ex.Left)
	if err != nil {
		return nil, err
	}
	r, err := i.eval(ex.Right)
	if err != nil {
		return nil, err
	}

	// Сравнения == и !=
	if ex.Op == "==" || ex.Op == "!=" {
		if !sameType(l, r) {
			return nil, semErr(ex.Line, ex.Col, "сравнение разных типов")
		}
		eq := l.Native() == r.Native()
		if ex.Op == "!=" {
			eq = !eq
		}
		return value.Bool(eq), nil
	}

	// Строковая конкатенация
	if ls, ok := l.(value.String); ok {
		rs, ok := r.(value.String)
		if !ok || ex.Op != "+" {
			return nil, semErr(ex.Line, ex.Col, "для строк допустим только +")
		}
		return value.String(string(ls) + string(rs)), nil
	}

	// Числовые операции: оба операнда одного числового типа
	switch lv := l.(type) {
	case value.Int:
		rv, ok := r.(value.Int)
		if !ok {
			return nil, semErr(ex.Line, ex.Col, "несоответствие типов в выражении")
		}
		return intOp(ex, int64(lv), int64(rv))
	case value.Float:
		rv, ok := r.(value.Float)
		if !ok {
			return nil, semErr(ex.Line, ex.Col, "несоответствие типов в выражении")
		}
		return floatOp(ex, float64(lv), float64(rv))
	}
	return nil, semErr(ex.Line, ex.Col, "оператор "+ex.Op+" неприменим к этим типам")
}

func sameType(a, b value.Value) bool {
	return typeNameOf(a) == typeNameOf(b)
}

func intOp(ex *ast.BinaryExpr, a, b int64) (value.Value, error) {
	switch ex.Op {
	case "+":
		return value.Int(a + b), nil
	case "-":
		return value.Int(a - b), nil
	case "*":
		return value.Int(a * b), nil
	case "/":
		if b == 0 {
			return nil, rtErr(ex.Line, ex.Col, "деление на ноль")
		}
		return value.Int(a / b), nil
	case "%":
		if b == 0 {
			return nil, rtErr(ex.Line, ex.Col, "деление на ноль")
		}
		return value.Int(a % b), nil
	case "<":
		return value.Bool(a < b), nil
	case "<=":
		return value.Bool(a <= b), nil
	case ">":
		return value.Bool(a > b), nil
	case ">=":
		return value.Bool(a >= b), nil
	}
	return nil, semErr(ex.Line, ex.Col, "неизвестный оператор "+ex.Op)
}

func floatOp(ex *ast.BinaryExpr, a, b float64) (value.Value, error) {
	switch ex.Op {
	case "+":
		return value.Float(a + b), nil
	case "-":
		return value.Float(a - b), nil
	case "*":
		return value.Float(a * b), nil
	case "/":
		if b == 0 {
			return nil, rtErr(ex.Line, ex.Col, "деление на ноль")
		}
		return value.Float(a / b), nil
	case "<":
		return value.Bool(a < b), nil
	case "<=":
		return value.Bool(a <= b), nil
	case ">":
		return value.Bool(a > b), nil
	case ">=":
		return value.Bool(a >= b), nil
	}
	return nil, semErr(ex.Line, ex.Col, "оператор "+ex.Op+" неприменим к float64")
}

func (i *Interp) evalStructLit(sl *ast.StructLit) (value.Value, error) {
	st, ok := i.types[sl.TypeName]
	if !ok {
		return nil, semErr(sl.Line, sl.Col, "неизвестный тип: "+sl.TypeName)
	}
	s := value.NewStruct(sl.TypeName)
	for _, fd := range st.Fields {
		fv, err := i.zeroValue(fd.Type, sl.Line, sl.Col)
		if err != nil {
			return nil, err
		}
		s.Set(fd.Name, fv)
	}
	for _, fi := range sl.Fields {
		found := false
		var declType string
		for _, fd := range st.Fields {
			if fd.Name == fi.Name {
				found = true
				declType = fd.Type
				break
			}
		}
		if !found {
			return nil, semErr(sl.Line, sl.Col,
				fmt.Sprintf("структура %s не имеет поля %s", sl.TypeName, fi.Name))
		}
		v, err := i.eval(fi.Value)
		if err != nil {
			return nil, err
		}
		if !i.typeMatches(declType, v) {
			return nil, semErr(sl.Line, sl.Col, fmt.Sprintf(
				"поле %s.%s: ожидался %s, получено %s",
				sl.TypeName, fi.Name, declType, typeNameOf(v),
			))
		}
		s.Set(fi.Name, v)
	}
	return s, nil
}

func (i *Interp) evalSelector(se *ast.SelectorExpr) (value.Value, error) {
	obj, err := i.eval(se.Object)
	if err != nil {
		return nil, err
	}
	s, ok := obj.(*value.Struct)
	if !ok {
		return nil, semErr(se.Line, se.Col,
			fmt.Sprintf("обращение к полю %s у не-структуры (тип %s)", se.Field, typeNameOf(obj)))
	}
	v, ok := s.Get(se.Field)
	if !ok {
		return nil, semErr(se.Line, se.Col,
			fmt.Sprintf("структура %s не имеет поля %s", s.TypeName, se.Field))
	}
	return v, nil
}

func (i *Interp) evalSliceLit(sl *ast.SliceLit) (value.Value, error) {
	s := value.NewSlice(sl.ElemType)
	for _, e := range sl.Elems {
		v, err := i.eval(e)
		if err != nil {
			return nil, err
		}
		if !i.typeMatches(sl.ElemType, v) {
			return nil, semErr(sl.Line, sl.Col, fmt.Sprintf(
				"элемент среза: ожидался %s, получено %s", sl.ElemType, typeNameOf(v),
			))
		}
		s.AppendInPlace(v)
	}
	return s, nil
}

func (i *Interp) evalIndex(ex *ast.IndexExpr) (value.Value, error) {
	obj, err := i.eval(ex.Object)
	if err != nil {
		return nil, err
	}
	idx, err := i.eval(ex.Index)
	if err != nil {
		return nil, err
	}
	idxInt, ok := idx.(value.Int)
	if !ok {
		return nil, semErr(ex.Line, ex.Col, "индекс должен быть int")
	}
	n := int64(idxInt)
	switch sv := obj.(type) {
	case *value.Slice:
		if n < 0 || n >= int64(sv.Len()) {
			return nil, rtErr(ex.Line, ex.Col, fmt.Sprintf(
				"индекс %d выходит за пределы среза длиной %d", n, sv.Len(),
			))
		}
		return sv.Get(int(n)), nil
	case value.String:
		s := string(sv)
		if n < 0 || n >= int64(len(s)) {
			return nil, rtErr(ex.Line, ex.Col, fmt.Sprintf(
				"индекс %d выходит за пределы строки длиной %d", n, len(s),
			))
		}
		return value.Int(int64(s[n])), nil
	}
	return nil, semErr(ex.Line, ex.Col, fmt.Sprintf(
		"индексирование неприменимо к %s", typeNameOf(obj),
	))
}

func (i *Interp) execIndexAssign(ia *ast.IndexAssignStmt) error {
	obj, ok := i.cur.get(ia.Object)
	if !ok {
		return semErr(ia.Line, ia.Col, "undefined: "+ia.Object)
	}
	sl, ok := obj.(*value.Slice)
	if !ok {
		return semErr(ia.Line, ia.Col, fmt.Sprintf("%s не является срезом", ia.Object))
	}
	idx, err := i.eval(ia.Index)
	if err != nil {
		return err
	}
	idxInt, ok := idx.(value.Int)
	if !ok {
		return semErr(ia.Line, ia.Col, "индекс должен быть int")
	}
	n := int64(idxInt)
	if n < 0 || n >= int64(sl.Len()) {
		return rtErr(ia.Line, ia.Col, fmt.Sprintf(
			"индекс %d выходит за пределы среза длиной %d", n, sl.Len(),
		))
	}
	val, err := i.eval(ia.Value)
	if err != nil {
		return err
	}
	if !i.typeMatches(sl.ElemType, val) {
		return semErr(ia.Line, ia.Col, fmt.Sprintf(
			"элемент среза: ожидался %s, получено %s", sl.ElemType, typeNameOf(val),
		))
	}
	sl.Set(int(n), val)
	return nil
}

func (i *Interp) execForRange(fs *ast.ForRangeStmt) error {
	sv, err := i.eval(fs.Slice)
	if err != nil {
		return err
	}
	sl, ok := sv.(*value.Slice)
	if !ok {
		return semErr(fs.Line, fs.Col, "range применимо только к срезу")
	}
	for idx := 0; idx < sl.Len(); idx++ {
		i.push()
		if fs.Key != "" && fs.Key != "_" {
			i.cur.declare(fs.Key, value.Int(int64(idx)))
		}
		if fs.Val != "" && fs.Val != "_" {
			i.cur.declare(fs.Val, sl.Get(idx))
		}
		err := i.execBlock(fs.Body)
		i.pop()
		if err != nil {
			if errors.Is(err, errBreak) {
				return nil
			}
			if errors.Is(err, errContinue) {
				continue
			}
			return err
		}
	}
	return nil
}

func (i *Interp) evalBuiltinLen(ex *ast.CallExpr) (value.Value, error) {
	if len(ex.Args) != 1 {
		return nil, semErr(ex.Line, ex.Col, "len: ожидается 1 аргумент")
	}
	v, err := i.eval(ex.Args[0])
	if err != nil {
		return nil, err
	}
	switch sv := v.(type) {
	case *value.Slice:
		return value.Int(int64(sv.Len())), nil
	case value.String:
		return value.Int(int64(len(string(sv)))), nil
	}
	return nil, semErr(ex.Line, ex.Col, fmt.Sprintf("len неприменим к %s", typeNameOf(v)))
}

func (i *Interp) evalBuiltinAppend(ex *ast.CallExpr) (value.Value, error) {
	if len(ex.Args) != 2 {
		return nil, semErr(ex.Line, ex.Col, "append: ожидается 2 аргумента")
	}
	sv, err := i.eval(ex.Args[0])
	if err != nil {
		return nil, err
	}
	sl, ok := sv.(*value.Slice)
	if !ok {
		return nil, semErr(ex.Line, ex.Col, "append: первый аргумент должен быть срезом")
	}
	val, err := i.eval(ex.Args[1])
	if err != nil {
		return nil, err
	}
	if !i.typeMatches(sl.ElemType, val) {
		return nil, semErr(ex.Line, ex.Col, fmt.Sprintf(
			"append: тип элемента %s, получено %s", sl.ElemType, typeNameOf(val),
		))
	}
	return sl.Append(val), nil
}

func (i *Interp) execFieldAssign(fa *ast.FieldAssignStmt) error {
	obj, ok := i.cur.get(fa.Object)
	if !ok {
		return semErr(fa.Line, fa.Col, "undefined: "+fa.Object)
	}
	s, ok := obj.(*value.Struct)
	if !ok {
		return semErr(fa.Line, fa.Col,
			fmt.Sprintf("%s не является структурой", fa.Object))
	}
	st, ok := i.types[s.TypeName]
	if !ok {
		return semErr(fa.Line, fa.Col, "неизвестный тип структуры: "+s.TypeName)
	}
	var declType string
	fieldFound := false
	for _, fd := range st.Fields {
		if fd.Name == fa.Field {
			declType = fd.Type
			fieldFound = true
			break
		}
	}
	if !fieldFound {
		return semErr(fa.Line, fa.Col,
			fmt.Sprintf("структура %s не имеет поля %s", s.TypeName, fa.Field))
	}
	v, err := i.eval(fa.Value)
	if err != nil {
		return err
	}
	if !i.typeMatches(declType, v) {
		return semErr(fa.Line, fa.Col, fmt.Sprintf(
			"поле %s.%s: ожидался %s, получено %s",
			s.TypeName, fa.Field, declType, typeNameOf(v),
		))
	}
	s.Set(fa.Field, v)
	return nil
}
