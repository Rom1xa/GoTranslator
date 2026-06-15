// Code generated from grammar/GoSubset.g4 by ANTLR 4.13.2. DO NOT EDIT.

package gen // GoSubset
import "github.com/antlr4-go/antlr/v4"

type BaseGoSubsetVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGoSubsetVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitTopLevel(ctx *TopLevelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitFuncDecl(ctx *FuncDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitTypeDecl(ctx *TypeDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitFieldDecl(ctx *FieldDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitParamList(ctx *ParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitParam(ctx *ParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitReturnSpec(ctx *ReturnSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitVarDecl(ctx *VarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitShortVarDecl(ctx *ShortVarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitAssignStmt(ctx *AssignStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitFieldAssignStmt(ctx *FieldAssignStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitIndexAssignStmt(ctx *IndexAssignStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitCallStmt(ctx *CallStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitMultiShortDecl(ctx *MultiShortDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitElsePart(ctx *ElsePartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitForInfinite(ctx *ForInfiniteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitForRangeFull(ctx *ForRangeFullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitForRangeIndex(ctx *ForRangeIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitForClause(ctx *ForClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitForWhile(ctx *ForWhileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitSimpleShort(ctx *SimpleShortContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitSimpleAssign(ctx *SimpleAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitFieldInitList(ctx *FieldInitListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitFieldInit(ctx *FieldInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitPrintStmt(ctx *PrintStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitExprList(ctx *ExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitMulExpr(ctx *MulExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitStringExpr(ctx *StringExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitIdentExpr(ctx *IdentExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitFloatExpr(ctx *FloatExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitTrueExpr(ctx *TrueExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitRelExpr(ctx *RelExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitAddExpr(ctx *AddExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitUnaryExpr(ctx *UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitIndexExpr(ctx *IndexExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitFalseExpr(ctx *FalseExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitSliceLitExpr(ctx *SliceLitExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitSelectorExpr(ctx *SelectorExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitEqExpr(ctx *EqExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitCallExpr(ctx *CallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitIntExpr(ctx *IntExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoSubsetVisitor) VisitStructLitExpr(ctx *StructLitExprContext) interface{} {
	return v.VisitChildren(ctx)
}
