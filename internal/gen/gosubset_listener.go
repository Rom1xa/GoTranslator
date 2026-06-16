// Code generated from grammar/GoSubset.g4 by ANTLR 4.13.2. DO NOT EDIT.

package gen // GoSubset
import "github.com/antlr4-go/antlr/v4"

// GoSubsetListener is a complete listener for a parse tree produced by GoSubsetParser.
type GoSubsetListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterTopLevel is called when entering the topLevel production.
	EnterTopLevel(c *TopLevelContext)

	// EnterFuncDecl is called when entering the funcDecl production.
	EnterFuncDecl(c *FuncDeclContext)

	// EnterTypeDecl is called when entering the typeDecl production.
	EnterTypeDecl(c *TypeDeclContext)

	// EnterTypeDef is called when entering the typeDef production.
	EnterTypeDef(c *TypeDefContext)

	// EnterFieldDecl is called when entering the fieldDecl production.
	EnterFieldDecl(c *FieldDeclContext)

	// EnterMethodDecl is called when entering the methodDecl production.
	EnterMethodDecl(c *MethodDeclContext)

	// EnterParamList is called when entering the paramList production.
	EnterParamList(c *ParamListContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterReturnSpec is called when entering the returnSpec production.
	EnterReturnSpec(c *ReturnSpecContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterVarDecl is called when entering the varDecl production.
	EnterVarDecl(c *VarDeclContext)

	// EnterShortVarDecl is called when entering the shortVarDecl production.
	EnterShortVarDecl(c *ShortVarDeclContext)

	// EnterAssignStmt is called when entering the assignStmt production.
	EnterAssignStmt(c *AssignStmtContext)

	// EnterFieldAssignStmt is called when entering the fieldAssignStmt production.
	EnterFieldAssignStmt(c *FieldAssignStmtContext)

	// EnterIndexAssignStmt is called when entering the indexAssignStmt production.
	EnterIndexAssignStmt(c *IndexAssignStmtContext)

	// EnterReturnStmt is called when entering the returnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterCallStmt is called when entering the callStmt production.
	EnterCallStmt(c *CallStmtContext)

	// EnterMultiShortDecl is called when entering the multiShortDecl production.
	EnterMultiShortDecl(c *MultiShortDeclContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterElsePart is called when entering the elsePart production.
	EnterElsePart(c *ElsePartContext)

	// EnterForInfinite is called when entering the ForInfinite production.
	EnterForInfinite(c *ForInfiniteContext)

	// EnterForRangeFull is called when entering the ForRangeFull production.
	EnterForRangeFull(c *ForRangeFullContext)

	// EnterForRangeIndex is called when entering the ForRangeIndex production.
	EnterForRangeIndex(c *ForRangeIndexContext)

	// EnterForClause is called when entering the ForClause production.
	EnterForClause(c *ForClauseContext)

	// EnterForWhile is called when entering the ForWhile production.
	EnterForWhile(c *ForWhileContext)

	// EnterSimpleShort is called when entering the SimpleShort production.
	EnterSimpleShort(c *SimpleShortContext)

	// EnterSimpleAssign is called when entering the SimpleAssign production.
	EnterSimpleAssign(c *SimpleAssignContext)

	// EnterBreakStmt is called when entering the breakStmt production.
	EnterBreakStmt(c *BreakStmtContext)

	// EnterContinueStmt is called when entering the continueStmt production.
	EnterContinueStmt(c *ContinueStmtContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterFieldInitList is called when entering the fieldInitList production.
	EnterFieldInitList(c *FieldInitListContext)

	// EnterFieldInit is called when entering the fieldInit production.
	EnterFieldInit(c *FieldInitContext)

	// EnterPrintStmt is called when entering the printStmt production.
	EnterPrintStmt(c *PrintStmtContext)

	// EnterGoStmt is called when entering the goStmt production.
	EnterGoStmt(c *GoStmtContext)

	// EnterExprList is called when entering the exprList production.
	EnterExprList(c *ExprListContext)

	// EnterMulExpr is called when entering the MulExpr production.
	EnterMulExpr(c *MulExprContext)

	// EnterAndExpr is called when entering the AndExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterStringExpr is called when entering the StringExpr production.
	EnterStringExpr(c *StringExprContext)

	// EnterIdentExpr is called when entering the IdentExpr production.
	EnterIdentExpr(c *IdentExprContext)

	// EnterFloatExpr is called when entering the FloatExpr production.
	EnterFloatExpr(c *FloatExprContext)

	// EnterTrueExpr is called when entering the TrueExpr production.
	EnterTrueExpr(c *TrueExprContext)

	// EnterRelExpr is called when entering the RelExpr production.
	EnterRelExpr(c *RelExprContext)

	// EnterAddExpr is called when entering the AddExpr production.
	EnterAddExpr(c *AddExprContext)

	// EnterUnaryExpr is called when entering the UnaryExpr production.
	EnterUnaryExpr(c *UnaryExprContext)

	// EnterOrExpr is called when entering the OrExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterIndexExpr is called when entering the IndexExpr production.
	EnterIndexExpr(c *IndexExprContext)

	// EnterFalseExpr is called when entering the FalseExpr production.
	EnterFalseExpr(c *FalseExprContext)

	// EnterSliceLitExpr is called when entering the SliceLitExpr production.
	EnterSliceLitExpr(c *SliceLitExprContext)

	// EnterSelectorExpr is called when entering the SelectorExpr production.
	EnterSelectorExpr(c *SelectorExprContext)

	// EnterEqExpr is called when entering the EqExpr production.
	EnterEqExpr(c *EqExprContext)

	// EnterCallExpr is called when entering the CallExpr production.
	EnterCallExpr(c *CallExprContext)

	// EnterIntExpr is called when entering the IntExpr production.
	EnterIntExpr(c *IntExprContext)

	// EnterParenExpr is called when entering the ParenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterStructLitExpr is called when entering the StructLitExpr production.
	EnterStructLitExpr(c *StructLitExprContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitTopLevel is called when exiting the topLevel production.
	ExitTopLevel(c *TopLevelContext)

	// ExitFuncDecl is called when exiting the funcDecl production.
	ExitFuncDecl(c *FuncDeclContext)

	// ExitTypeDecl is called when exiting the typeDecl production.
	ExitTypeDecl(c *TypeDeclContext)

	// ExitTypeDef is called when exiting the typeDef production.
	ExitTypeDef(c *TypeDefContext)

	// ExitFieldDecl is called when exiting the fieldDecl production.
	ExitFieldDecl(c *FieldDeclContext)

	// ExitMethodDecl is called when exiting the methodDecl production.
	ExitMethodDecl(c *MethodDeclContext)

	// ExitParamList is called when exiting the paramList production.
	ExitParamList(c *ParamListContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitReturnSpec is called when exiting the returnSpec production.
	ExitReturnSpec(c *ReturnSpecContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitVarDecl is called when exiting the varDecl production.
	ExitVarDecl(c *VarDeclContext)

	// ExitShortVarDecl is called when exiting the shortVarDecl production.
	ExitShortVarDecl(c *ShortVarDeclContext)

	// ExitAssignStmt is called when exiting the assignStmt production.
	ExitAssignStmt(c *AssignStmtContext)

	// ExitFieldAssignStmt is called when exiting the fieldAssignStmt production.
	ExitFieldAssignStmt(c *FieldAssignStmtContext)

	// ExitIndexAssignStmt is called when exiting the indexAssignStmt production.
	ExitIndexAssignStmt(c *IndexAssignStmtContext)

	// ExitReturnStmt is called when exiting the returnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitCallStmt is called when exiting the callStmt production.
	ExitCallStmt(c *CallStmtContext)

	// ExitMultiShortDecl is called when exiting the multiShortDecl production.
	ExitMultiShortDecl(c *MultiShortDeclContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitElsePart is called when exiting the elsePart production.
	ExitElsePart(c *ElsePartContext)

	// ExitForInfinite is called when exiting the ForInfinite production.
	ExitForInfinite(c *ForInfiniteContext)

	// ExitForRangeFull is called when exiting the ForRangeFull production.
	ExitForRangeFull(c *ForRangeFullContext)

	// ExitForRangeIndex is called when exiting the ForRangeIndex production.
	ExitForRangeIndex(c *ForRangeIndexContext)

	// ExitForClause is called when exiting the ForClause production.
	ExitForClause(c *ForClauseContext)

	// ExitForWhile is called when exiting the ForWhile production.
	ExitForWhile(c *ForWhileContext)

	// ExitSimpleShort is called when exiting the SimpleShort production.
	ExitSimpleShort(c *SimpleShortContext)

	// ExitSimpleAssign is called when exiting the SimpleAssign production.
	ExitSimpleAssign(c *SimpleAssignContext)

	// ExitBreakStmt is called when exiting the breakStmt production.
	ExitBreakStmt(c *BreakStmtContext)

	// ExitContinueStmt is called when exiting the continueStmt production.
	ExitContinueStmt(c *ContinueStmtContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitFieldInitList is called when exiting the fieldInitList production.
	ExitFieldInitList(c *FieldInitListContext)

	// ExitFieldInit is called when exiting the fieldInit production.
	ExitFieldInit(c *FieldInitContext)

	// ExitPrintStmt is called when exiting the printStmt production.
	ExitPrintStmt(c *PrintStmtContext)

	// ExitGoStmt is called when exiting the goStmt production.
	ExitGoStmt(c *GoStmtContext)

	// ExitExprList is called when exiting the exprList production.
	ExitExprList(c *ExprListContext)

	// ExitMulExpr is called when exiting the MulExpr production.
	ExitMulExpr(c *MulExprContext)

	// ExitAndExpr is called when exiting the AndExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitStringExpr is called when exiting the StringExpr production.
	ExitStringExpr(c *StringExprContext)

	// ExitIdentExpr is called when exiting the IdentExpr production.
	ExitIdentExpr(c *IdentExprContext)

	// ExitFloatExpr is called when exiting the FloatExpr production.
	ExitFloatExpr(c *FloatExprContext)

	// ExitTrueExpr is called when exiting the TrueExpr production.
	ExitTrueExpr(c *TrueExprContext)

	// ExitRelExpr is called when exiting the RelExpr production.
	ExitRelExpr(c *RelExprContext)

	// ExitAddExpr is called when exiting the AddExpr production.
	ExitAddExpr(c *AddExprContext)

	// ExitUnaryExpr is called when exiting the UnaryExpr production.
	ExitUnaryExpr(c *UnaryExprContext)

	// ExitOrExpr is called when exiting the OrExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitIndexExpr is called when exiting the IndexExpr production.
	ExitIndexExpr(c *IndexExprContext)

	// ExitFalseExpr is called when exiting the FalseExpr production.
	ExitFalseExpr(c *FalseExprContext)

	// ExitSliceLitExpr is called when exiting the SliceLitExpr production.
	ExitSliceLitExpr(c *SliceLitExprContext)

	// ExitSelectorExpr is called when exiting the SelectorExpr production.
	ExitSelectorExpr(c *SelectorExprContext)

	// ExitEqExpr is called when exiting the EqExpr production.
	ExitEqExpr(c *EqExprContext)

	// ExitCallExpr is called when exiting the CallExpr production.
	ExitCallExpr(c *CallExprContext)

	// ExitIntExpr is called when exiting the IntExpr production.
	ExitIntExpr(c *IntExprContext)

	// ExitParenExpr is called when exiting the ParenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitStructLitExpr is called when exiting the StructLitExpr production.
	ExitStructLitExpr(c *StructLitExprContext)
}
