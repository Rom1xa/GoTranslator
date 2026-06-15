// Code generated from grammar/GoSubset.g4 by ANTLR 4.13.2. DO NOT EDIT.

package gen // GoSubset
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by GoSubsetParser.
type GoSubsetVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GoSubsetParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#topLevel.
	VisitTopLevel(ctx *TopLevelContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#funcDecl.
	VisitFuncDecl(ctx *FuncDeclContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#typeDecl.
	VisitTypeDecl(ctx *TypeDeclContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#fieldDecl.
	VisitFieldDecl(ctx *FieldDeclContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#paramList.
	VisitParamList(ctx *ParamListContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#param.
	VisitParam(ctx *ParamContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#returnSpec.
	VisitReturnSpec(ctx *ReturnSpecContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#varDecl.
	VisitVarDecl(ctx *VarDeclContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#shortVarDecl.
	VisitShortVarDecl(ctx *ShortVarDeclContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#assignStmt.
	VisitAssignStmt(ctx *AssignStmtContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#fieldAssignStmt.
	VisitFieldAssignStmt(ctx *FieldAssignStmtContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#indexAssignStmt.
	VisitIndexAssignStmt(ctx *IndexAssignStmtContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#returnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#callStmt.
	VisitCallStmt(ctx *CallStmtContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#multiShortDecl.
	VisitMultiShortDecl(ctx *MultiShortDeclContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#elsePart.
	VisitElsePart(ctx *ElsePartContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#ForInfinite.
	VisitForInfinite(ctx *ForInfiniteContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#ForRangeFull.
	VisitForRangeFull(ctx *ForRangeFullContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#ForRangeIndex.
	VisitForRangeIndex(ctx *ForRangeIndexContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#ForClause.
	VisitForClause(ctx *ForClauseContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#ForWhile.
	VisitForWhile(ctx *ForWhileContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#SimpleShort.
	VisitSimpleShort(ctx *SimpleShortContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#SimpleAssign.
	VisitSimpleAssign(ctx *SimpleAssignContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#breakStmt.
	VisitBreakStmt(ctx *BreakStmtContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#continueStmt.
	VisitContinueStmt(ctx *ContinueStmtContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#fieldInitList.
	VisitFieldInitList(ctx *FieldInitListContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#fieldInit.
	VisitFieldInit(ctx *FieldInitContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#printStmt.
	VisitPrintStmt(ctx *PrintStmtContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#exprList.
	VisitExprList(ctx *ExprListContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#MulExpr.
	VisitMulExpr(ctx *MulExprContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#AndExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#StringExpr.
	VisitStringExpr(ctx *StringExprContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#IdentExpr.
	VisitIdentExpr(ctx *IdentExprContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#FloatExpr.
	VisitFloatExpr(ctx *FloatExprContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#TrueExpr.
	VisitTrueExpr(ctx *TrueExprContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#RelExpr.
	VisitRelExpr(ctx *RelExprContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#AddExpr.
	VisitAddExpr(ctx *AddExprContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#UnaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#OrExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#IndexExpr.
	VisitIndexExpr(ctx *IndexExprContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#FalseExpr.
	VisitFalseExpr(ctx *FalseExprContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#SliceLitExpr.
	VisitSliceLitExpr(ctx *SliceLitExprContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#SelectorExpr.
	VisitSelectorExpr(ctx *SelectorExprContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#EqExpr.
	VisitEqExpr(ctx *EqExprContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#CallExpr.
	VisitCallExpr(ctx *CallExprContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#IntExpr.
	VisitIntExpr(ctx *IntExprContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#ParenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by GoSubsetParser#StructLitExpr.
	VisitStructLitExpr(ctx *StructLitExprContext) interface{}
}
