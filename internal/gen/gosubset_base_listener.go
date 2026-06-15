// Code generated from grammar/GoSubset.g4 by ANTLR 4.13.2. DO NOT EDIT.

package gen // GoSubset
import "github.com/antlr4-go/antlr/v4"

// BaseGoSubsetListener is a complete listener for a parse tree produced by GoSubsetParser.
type BaseGoSubsetListener struct{}

var _ GoSubsetListener = &BaseGoSubsetListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGoSubsetListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGoSubsetListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGoSubsetListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGoSubsetListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseGoSubsetListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseGoSubsetListener) ExitProgram(ctx *ProgramContext) {}

// EnterTopLevel is called when production topLevel is entered.
func (s *BaseGoSubsetListener) EnterTopLevel(ctx *TopLevelContext) {}

// ExitTopLevel is called when production topLevel is exited.
func (s *BaseGoSubsetListener) ExitTopLevel(ctx *TopLevelContext) {}

// EnterFuncDecl is called when production funcDecl is entered.
func (s *BaseGoSubsetListener) EnterFuncDecl(ctx *FuncDeclContext) {}

// ExitFuncDecl is called when production funcDecl is exited.
func (s *BaseGoSubsetListener) ExitFuncDecl(ctx *FuncDeclContext) {}

// EnterTypeDecl is called when production typeDecl is entered.
func (s *BaseGoSubsetListener) EnterTypeDecl(ctx *TypeDeclContext) {}

// ExitTypeDecl is called when production typeDecl is exited.
func (s *BaseGoSubsetListener) ExitTypeDecl(ctx *TypeDeclContext) {}

// EnterFieldDecl is called when production fieldDecl is entered.
func (s *BaseGoSubsetListener) EnterFieldDecl(ctx *FieldDeclContext) {}

// ExitFieldDecl is called when production fieldDecl is exited.
func (s *BaseGoSubsetListener) ExitFieldDecl(ctx *FieldDeclContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BaseGoSubsetListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BaseGoSubsetListener) ExitParamList(ctx *ParamListContext) {}

// EnterParam is called when production param is entered.
func (s *BaseGoSubsetListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseGoSubsetListener) ExitParam(ctx *ParamContext) {}

// EnterReturnSpec is called when production returnSpec is entered.
func (s *BaseGoSubsetListener) EnterReturnSpec(ctx *ReturnSpecContext) {}

// ExitReturnSpec is called when production returnSpec is exited.
func (s *BaseGoSubsetListener) ExitReturnSpec(ctx *ReturnSpecContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseGoSubsetListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseGoSubsetListener) ExitBlock(ctx *BlockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseGoSubsetListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseGoSubsetListener) ExitStatement(ctx *StatementContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *BaseGoSubsetListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BaseGoSubsetListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterShortVarDecl is called when production shortVarDecl is entered.
func (s *BaseGoSubsetListener) EnterShortVarDecl(ctx *ShortVarDeclContext) {}

// ExitShortVarDecl is called when production shortVarDecl is exited.
func (s *BaseGoSubsetListener) ExitShortVarDecl(ctx *ShortVarDeclContext) {}

// EnterAssignStmt is called when production assignStmt is entered.
func (s *BaseGoSubsetListener) EnterAssignStmt(ctx *AssignStmtContext) {}

// ExitAssignStmt is called when production assignStmt is exited.
func (s *BaseGoSubsetListener) ExitAssignStmt(ctx *AssignStmtContext) {}

// EnterFieldAssignStmt is called when production fieldAssignStmt is entered.
func (s *BaseGoSubsetListener) EnterFieldAssignStmt(ctx *FieldAssignStmtContext) {}

// ExitFieldAssignStmt is called when production fieldAssignStmt is exited.
func (s *BaseGoSubsetListener) ExitFieldAssignStmt(ctx *FieldAssignStmtContext) {}

// EnterIndexAssignStmt is called when production indexAssignStmt is entered.
func (s *BaseGoSubsetListener) EnterIndexAssignStmt(ctx *IndexAssignStmtContext) {}

// ExitIndexAssignStmt is called when production indexAssignStmt is exited.
func (s *BaseGoSubsetListener) ExitIndexAssignStmt(ctx *IndexAssignStmtContext) {}

// EnterReturnStmt is called when production returnStmt is entered.
func (s *BaseGoSubsetListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production returnStmt is exited.
func (s *BaseGoSubsetListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterCallStmt is called when production callStmt is entered.
func (s *BaseGoSubsetListener) EnterCallStmt(ctx *CallStmtContext) {}

// ExitCallStmt is called when production callStmt is exited.
func (s *BaseGoSubsetListener) ExitCallStmt(ctx *CallStmtContext) {}

// EnterMultiShortDecl is called when production multiShortDecl is entered.
func (s *BaseGoSubsetListener) EnterMultiShortDecl(ctx *MultiShortDeclContext) {}

// ExitMultiShortDecl is called when production multiShortDecl is exited.
func (s *BaseGoSubsetListener) ExitMultiShortDecl(ctx *MultiShortDeclContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseGoSubsetListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseGoSubsetListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterElsePart is called when production elsePart is entered.
func (s *BaseGoSubsetListener) EnterElsePart(ctx *ElsePartContext) {}

// ExitElsePart is called when production elsePart is exited.
func (s *BaseGoSubsetListener) ExitElsePart(ctx *ElsePartContext) {}

// EnterForInfinite is called when production ForInfinite is entered.
func (s *BaseGoSubsetListener) EnterForInfinite(ctx *ForInfiniteContext) {}

// ExitForInfinite is called when production ForInfinite is exited.
func (s *BaseGoSubsetListener) ExitForInfinite(ctx *ForInfiniteContext) {}

// EnterForRangeFull is called when production ForRangeFull is entered.
func (s *BaseGoSubsetListener) EnterForRangeFull(ctx *ForRangeFullContext) {}

// ExitForRangeFull is called when production ForRangeFull is exited.
func (s *BaseGoSubsetListener) ExitForRangeFull(ctx *ForRangeFullContext) {}

// EnterForRangeIndex is called when production ForRangeIndex is entered.
func (s *BaseGoSubsetListener) EnterForRangeIndex(ctx *ForRangeIndexContext) {}

// ExitForRangeIndex is called when production ForRangeIndex is exited.
func (s *BaseGoSubsetListener) ExitForRangeIndex(ctx *ForRangeIndexContext) {}

// EnterForClause is called when production ForClause is entered.
func (s *BaseGoSubsetListener) EnterForClause(ctx *ForClauseContext) {}

// ExitForClause is called when production ForClause is exited.
func (s *BaseGoSubsetListener) ExitForClause(ctx *ForClauseContext) {}

// EnterForWhile is called when production ForWhile is entered.
func (s *BaseGoSubsetListener) EnterForWhile(ctx *ForWhileContext) {}

// ExitForWhile is called when production ForWhile is exited.
func (s *BaseGoSubsetListener) ExitForWhile(ctx *ForWhileContext) {}

// EnterSimpleShort is called when production SimpleShort is entered.
func (s *BaseGoSubsetListener) EnterSimpleShort(ctx *SimpleShortContext) {}

// ExitSimpleShort is called when production SimpleShort is exited.
func (s *BaseGoSubsetListener) ExitSimpleShort(ctx *SimpleShortContext) {}

// EnterSimpleAssign is called when production SimpleAssign is entered.
func (s *BaseGoSubsetListener) EnterSimpleAssign(ctx *SimpleAssignContext) {}

// ExitSimpleAssign is called when production SimpleAssign is exited.
func (s *BaseGoSubsetListener) ExitSimpleAssign(ctx *SimpleAssignContext) {}

// EnterBreakStmt is called when production breakStmt is entered.
func (s *BaseGoSubsetListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production breakStmt is exited.
func (s *BaseGoSubsetListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterContinueStmt is called when production continueStmt is entered.
func (s *BaseGoSubsetListener) EnterContinueStmt(ctx *ContinueStmtContext) {}

// ExitContinueStmt is called when production continueStmt is exited.
func (s *BaseGoSubsetListener) ExitContinueStmt(ctx *ContinueStmtContext) {}

// EnterType is called when production type is entered.
func (s *BaseGoSubsetListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseGoSubsetListener) ExitType(ctx *TypeContext) {}

// EnterFieldInitList is called when production fieldInitList is entered.
func (s *BaseGoSubsetListener) EnterFieldInitList(ctx *FieldInitListContext) {}

// ExitFieldInitList is called when production fieldInitList is exited.
func (s *BaseGoSubsetListener) ExitFieldInitList(ctx *FieldInitListContext) {}

// EnterFieldInit is called when production fieldInit is entered.
func (s *BaseGoSubsetListener) EnterFieldInit(ctx *FieldInitContext) {}

// ExitFieldInit is called when production fieldInit is exited.
func (s *BaseGoSubsetListener) ExitFieldInit(ctx *FieldInitContext) {}

// EnterPrintStmt is called when production printStmt is entered.
func (s *BaseGoSubsetListener) EnterPrintStmt(ctx *PrintStmtContext) {}

// ExitPrintStmt is called when production printStmt is exited.
func (s *BaseGoSubsetListener) ExitPrintStmt(ctx *PrintStmtContext) {}

// EnterExprList is called when production exprList is entered.
func (s *BaseGoSubsetListener) EnterExprList(ctx *ExprListContext) {}

// ExitExprList is called when production exprList is exited.
func (s *BaseGoSubsetListener) ExitExprList(ctx *ExprListContext) {}

// EnterMulExpr is called when production MulExpr is entered.
func (s *BaseGoSubsetListener) EnterMulExpr(ctx *MulExprContext) {}

// ExitMulExpr is called when production MulExpr is exited.
func (s *BaseGoSubsetListener) ExitMulExpr(ctx *MulExprContext) {}

// EnterAndExpr is called when production AndExpr is entered.
func (s *BaseGoSubsetListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production AndExpr is exited.
func (s *BaseGoSubsetListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterStringExpr is called when production StringExpr is entered.
func (s *BaseGoSubsetListener) EnterStringExpr(ctx *StringExprContext) {}

// ExitStringExpr is called when production StringExpr is exited.
func (s *BaseGoSubsetListener) ExitStringExpr(ctx *StringExprContext) {}

// EnterIdentExpr is called when production IdentExpr is entered.
func (s *BaseGoSubsetListener) EnterIdentExpr(ctx *IdentExprContext) {}

// ExitIdentExpr is called when production IdentExpr is exited.
func (s *BaseGoSubsetListener) ExitIdentExpr(ctx *IdentExprContext) {}

// EnterFloatExpr is called when production FloatExpr is entered.
func (s *BaseGoSubsetListener) EnterFloatExpr(ctx *FloatExprContext) {}

// ExitFloatExpr is called when production FloatExpr is exited.
func (s *BaseGoSubsetListener) ExitFloatExpr(ctx *FloatExprContext) {}

// EnterTrueExpr is called when production TrueExpr is entered.
func (s *BaseGoSubsetListener) EnterTrueExpr(ctx *TrueExprContext) {}

// ExitTrueExpr is called when production TrueExpr is exited.
func (s *BaseGoSubsetListener) ExitTrueExpr(ctx *TrueExprContext) {}

// EnterRelExpr is called when production RelExpr is entered.
func (s *BaseGoSubsetListener) EnterRelExpr(ctx *RelExprContext) {}

// ExitRelExpr is called when production RelExpr is exited.
func (s *BaseGoSubsetListener) ExitRelExpr(ctx *RelExprContext) {}

// EnterAddExpr is called when production AddExpr is entered.
func (s *BaseGoSubsetListener) EnterAddExpr(ctx *AddExprContext) {}

// ExitAddExpr is called when production AddExpr is exited.
func (s *BaseGoSubsetListener) ExitAddExpr(ctx *AddExprContext) {}

// EnterUnaryExpr is called when production UnaryExpr is entered.
func (s *BaseGoSubsetListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production UnaryExpr is exited.
func (s *BaseGoSubsetListener) ExitUnaryExpr(ctx *UnaryExprContext) {}

// EnterOrExpr is called when production OrExpr is entered.
func (s *BaseGoSubsetListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production OrExpr is exited.
func (s *BaseGoSubsetListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterIndexExpr is called when production IndexExpr is entered.
func (s *BaseGoSubsetListener) EnterIndexExpr(ctx *IndexExprContext) {}

// ExitIndexExpr is called when production IndexExpr is exited.
func (s *BaseGoSubsetListener) ExitIndexExpr(ctx *IndexExprContext) {}

// EnterFalseExpr is called when production FalseExpr is entered.
func (s *BaseGoSubsetListener) EnterFalseExpr(ctx *FalseExprContext) {}

// ExitFalseExpr is called when production FalseExpr is exited.
func (s *BaseGoSubsetListener) ExitFalseExpr(ctx *FalseExprContext) {}

// EnterSliceLitExpr is called when production SliceLitExpr is entered.
func (s *BaseGoSubsetListener) EnterSliceLitExpr(ctx *SliceLitExprContext) {}

// ExitSliceLitExpr is called when production SliceLitExpr is exited.
func (s *BaseGoSubsetListener) ExitSliceLitExpr(ctx *SliceLitExprContext) {}

// EnterSelectorExpr is called when production SelectorExpr is entered.
func (s *BaseGoSubsetListener) EnterSelectorExpr(ctx *SelectorExprContext) {}

// ExitSelectorExpr is called when production SelectorExpr is exited.
func (s *BaseGoSubsetListener) ExitSelectorExpr(ctx *SelectorExprContext) {}

// EnterEqExpr is called when production EqExpr is entered.
func (s *BaseGoSubsetListener) EnterEqExpr(ctx *EqExprContext) {}

// ExitEqExpr is called when production EqExpr is exited.
func (s *BaseGoSubsetListener) ExitEqExpr(ctx *EqExprContext) {}

// EnterCallExpr is called when production CallExpr is entered.
func (s *BaseGoSubsetListener) EnterCallExpr(ctx *CallExprContext) {}

// ExitCallExpr is called when production CallExpr is exited.
func (s *BaseGoSubsetListener) ExitCallExpr(ctx *CallExprContext) {}

// EnterIntExpr is called when production IntExpr is entered.
func (s *BaseGoSubsetListener) EnterIntExpr(ctx *IntExprContext) {}

// ExitIntExpr is called when production IntExpr is exited.
func (s *BaseGoSubsetListener) ExitIntExpr(ctx *IntExprContext) {}

// EnterParenExpr is called when production ParenExpr is entered.
func (s *BaseGoSubsetListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production ParenExpr is exited.
func (s *BaseGoSubsetListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterStructLitExpr is called when production StructLitExpr is entered.
func (s *BaseGoSubsetListener) EnterStructLitExpr(ctx *StructLitExprContext) {}

// ExitStructLitExpr is called when production StructLitExpr is exited.
func (s *BaseGoSubsetListener) ExitStructLitExpr(ctx *StructLitExprContext) {}
