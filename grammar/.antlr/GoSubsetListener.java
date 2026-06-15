// Generated from /home/user/mirea/translators/kyrsoTranslatorGo/grammar/GoSubset.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link GoSubsetParser}.
 */
public interface GoSubsetListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#program}.
	 * @param ctx the parse tree
	 */
	void enterProgram(GoSubsetParser.ProgramContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#program}.
	 * @param ctx the parse tree
	 */
	void exitProgram(GoSubsetParser.ProgramContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#topLevel}.
	 * @param ctx the parse tree
	 */
	void enterTopLevel(GoSubsetParser.TopLevelContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#topLevel}.
	 * @param ctx the parse tree
	 */
	void exitTopLevel(GoSubsetParser.TopLevelContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#funcDecl}.
	 * @param ctx the parse tree
	 */
	void enterFuncDecl(GoSubsetParser.FuncDeclContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#funcDecl}.
	 * @param ctx the parse tree
	 */
	void exitFuncDecl(GoSubsetParser.FuncDeclContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#typeDecl}.
	 * @param ctx the parse tree
	 */
	void enterTypeDecl(GoSubsetParser.TypeDeclContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#typeDecl}.
	 * @param ctx the parse tree
	 */
	void exitTypeDecl(GoSubsetParser.TypeDeclContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#fieldDecl}.
	 * @param ctx the parse tree
	 */
	void enterFieldDecl(GoSubsetParser.FieldDeclContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#fieldDecl}.
	 * @param ctx the parse tree
	 */
	void exitFieldDecl(GoSubsetParser.FieldDeclContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#paramList}.
	 * @param ctx the parse tree
	 */
	void enterParamList(GoSubsetParser.ParamListContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#paramList}.
	 * @param ctx the parse tree
	 */
	void exitParamList(GoSubsetParser.ParamListContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#param}.
	 * @param ctx the parse tree
	 */
	void enterParam(GoSubsetParser.ParamContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#param}.
	 * @param ctx the parse tree
	 */
	void exitParam(GoSubsetParser.ParamContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SingleReturnSpec}
	 * labeled alternative in {@link GoSubsetParser#returnSpec}.
	 * @param ctx the parse tree
	 */
	void enterSingleReturnSpec(GoSubsetParser.SingleReturnSpecContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SingleReturnSpec}
	 * labeled alternative in {@link GoSubsetParser#returnSpec}.
	 * @param ctx the parse tree
	 */
	void exitSingleReturnSpec(GoSubsetParser.SingleReturnSpecContext ctx);
	/**
	 * Enter a parse tree produced by the {@code MultiReturnSpec}
	 * labeled alternative in {@link GoSubsetParser#returnSpec}.
	 * @param ctx the parse tree
	 */
	void enterMultiReturnSpec(GoSubsetParser.MultiReturnSpecContext ctx);
	/**
	 * Exit a parse tree produced by the {@code MultiReturnSpec}
	 * labeled alternative in {@link GoSubsetParser#returnSpec}.
	 * @param ctx the parse tree
	 */
	void exitMultiReturnSpec(GoSubsetParser.MultiReturnSpecContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(GoSubsetParser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(GoSubsetParser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatement(GoSubsetParser.StatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatement(GoSubsetParser.StatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#varDecl}.
	 * @param ctx the parse tree
	 */
	void enterVarDecl(GoSubsetParser.VarDeclContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#varDecl}.
	 * @param ctx the parse tree
	 */
	void exitVarDecl(GoSubsetParser.VarDeclContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#shortVarDecl}.
	 * @param ctx the parse tree
	 */
	void enterShortVarDecl(GoSubsetParser.ShortVarDeclContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#shortVarDecl}.
	 * @param ctx the parse tree
	 */
	void exitShortVarDecl(GoSubsetParser.ShortVarDeclContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#assignStmt}.
	 * @param ctx the parse tree
	 */
	void enterAssignStmt(GoSubsetParser.AssignStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#assignStmt}.
	 * @param ctx the parse tree
	 */
	void exitAssignStmt(GoSubsetParser.AssignStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#fieldAssignStmt}.
	 * @param ctx the parse tree
	 */
	void enterFieldAssignStmt(GoSubsetParser.FieldAssignStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#fieldAssignStmt}.
	 * @param ctx the parse tree
	 */
	void exitFieldAssignStmt(GoSubsetParser.FieldAssignStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#indexAssignStmt}.
	 * @param ctx the parse tree
	 */
	void enterIndexAssignStmt(GoSubsetParser.IndexAssignStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#indexAssignStmt}.
	 * @param ctx the parse tree
	 */
	void exitIndexAssignStmt(GoSubsetParser.IndexAssignStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#returnStmt}.
	 * @param ctx the parse tree
	 */
	void enterReturnStmt(GoSubsetParser.ReturnStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#returnStmt}.
	 * @param ctx the parse tree
	 */
	void exitReturnStmt(GoSubsetParser.ReturnStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#callStmt}.
	 * @param ctx the parse tree
	 */
	void enterCallStmt(GoSubsetParser.CallStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#callStmt}.
	 * @param ctx the parse tree
	 */
	void exitCallStmt(GoSubsetParser.CallStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#multiShortDecl}.
	 * @param ctx the parse tree
	 */
	void enterMultiShortDecl(GoSubsetParser.MultiShortDeclContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#multiShortDecl}.
	 * @param ctx the parse tree
	 */
	void exitMultiShortDecl(GoSubsetParser.MultiShortDeclContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#ifStmt}.
	 * @param ctx the parse tree
	 */
	void enterIfStmt(GoSubsetParser.IfStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#ifStmt}.
	 * @param ctx the parse tree
	 */
	void exitIfStmt(GoSubsetParser.IfStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ElseIf}
	 * labeled alternative in {@link GoSubsetParser#elsePart}.
	 * @param ctx the parse tree
	 */
	void enterElseIf(GoSubsetParser.ElseIfContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ElseIf}
	 * labeled alternative in {@link GoSubsetParser#elsePart}.
	 * @param ctx the parse tree
	 */
	void exitElseIf(GoSubsetParser.ElseIfContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ElseBlock}
	 * labeled alternative in {@link GoSubsetParser#elsePart}.
	 * @param ctx the parse tree
	 */
	void enterElseBlock(GoSubsetParser.ElseBlockContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ElseBlock}
	 * labeled alternative in {@link GoSubsetParser#elsePart}.
	 * @param ctx the parse tree
	 */
	void exitElseBlock(GoSubsetParser.ElseBlockContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ForInfinite}
	 * labeled alternative in {@link GoSubsetParser#forStmt}.
	 * @param ctx the parse tree
	 */
	void enterForInfinite(GoSubsetParser.ForInfiniteContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ForInfinite}
	 * labeled alternative in {@link GoSubsetParser#forStmt}.
	 * @param ctx the parse tree
	 */
	void exitForInfinite(GoSubsetParser.ForInfiniteContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ForRangeFull}
	 * labeled alternative in {@link GoSubsetParser#forStmt}.
	 * @param ctx the parse tree
	 */
	void enterForRangeFull(GoSubsetParser.ForRangeFullContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ForRangeFull}
	 * labeled alternative in {@link GoSubsetParser#forStmt}.
	 * @param ctx the parse tree
	 */
	void exitForRangeFull(GoSubsetParser.ForRangeFullContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ForRangeIndex}
	 * labeled alternative in {@link GoSubsetParser#forStmt}.
	 * @param ctx the parse tree
	 */
	void enterForRangeIndex(GoSubsetParser.ForRangeIndexContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ForRangeIndex}
	 * labeled alternative in {@link GoSubsetParser#forStmt}.
	 * @param ctx the parse tree
	 */
	void exitForRangeIndex(GoSubsetParser.ForRangeIndexContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ForClause}
	 * labeled alternative in {@link GoSubsetParser#forStmt}.
	 * @param ctx the parse tree
	 */
	void enterForClause(GoSubsetParser.ForClauseContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ForClause}
	 * labeled alternative in {@link GoSubsetParser#forStmt}.
	 * @param ctx the parse tree
	 */
	void exitForClause(GoSubsetParser.ForClauseContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ForWhile}
	 * labeled alternative in {@link GoSubsetParser#forStmt}.
	 * @param ctx the parse tree
	 */
	void enterForWhile(GoSubsetParser.ForWhileContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ForWhile}
	 * labeled alternative in {@link GoSubsetParser#forStmt}.
	 * @param ctx the parse tree
	 */
	void exitForWhile(GoSubsetParser.ForWhileContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SimpleShort}
	 * labeled alternative in {@link GoSubsetParser#simpleStmt}.
	 * @param ctx the parse tree
	 */
	void enterSimpleShort(GoSubsetParser.SimpleShortContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SimpleShort}
	 * labeled alternative in {@link GoSubsetParser#simpleStmt}.
	 * @param ctx the parse tree
	 */
	void exitSimpleShort(GoSubsetParser.SimpleShortContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SimpleAssign}
	 * labeled alternative in {@link GoSubsetParser#simpleStmt}.
	 * @param ctx the parse tree
	 */
	void enterSimpleAssign(GoSubsetParser.SimpleAssignContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SimpleAssign}
	 * labeled alternative in {@link GoSubsetParser#simpleStmt}.
	 * @param ctx the parse tree
	 */
	void exitSimpleAssign(GoSubsetParser.SimpleAssignContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#breakStmt}.
	 * @param ctx the parse tree
	 */
	void enterBreakStmt(GoSubsetParser.BreakStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#breakStmt}.
	 * @param ctx the parse tree
	 */
	void exitBreakStmt(GoSubsetParser.BreakStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#continueStmt}.
	 * @param ctx the parse tree
	 */
	void enterContinueStmt(GoSubsetParser.ContinueStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#continueStmt}.
	 * @param ctx the parse tree
	 */
	void exitContinueStmt(GoSubsetParser.ContinueStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#type}.
	 * @param ctx the parse tree
	 */
	void enterType(GoSubsetParser.TypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#type}.
	 * @param ctx the parse tree
	 */
	void exitType(GoSubsetParser.TypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#fieldInitList}.
	 * @param ctx the parse tree
	 */
	void enterFieldInitList(GoSubsetParser.FieldInitListContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#fieldInitList}.
	 * @param ctx the parse tree
	 */
	void exitFieldInitList(GoSubsetParser.FieldInitListContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#fieldInit}.
	 * @param ctx the parse tree
	 */
	void enterFieldInit(GoSubsetParser.FieldInitContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#fieldInit}.
	 * @param ctx the parse tree
	 */
	void exitFieldInit(GoSubsetParser.FieldInitContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#printStmt}.
	 * @param ctx the parse tree
	 */
	void enterPrintStmt(GoSubsetParser.PrintStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#printStmt}.
	 * @param ctx the parse tree
	 */
	void exitPrintStmt(GoSubsetParser.PrintStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoSubsetParser#exprList}.
	 * @param ctx the parse tree
	 */
	void enterExprList(GoSubsetParser.ExprListContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoSubsetParser#exprList}.
	 * @param ctx the parse tree
	 */
	void exitExprList(GoSubsetParser.ExprListContext ctx);
	/**
	 * Enter a parse tree produced by the {@code MulExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterMulExpr(GoSubsetParser.MulExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code MulExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitMulExpr(GoSubsetParser.MulExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AndExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterAndExpr(GoSubsetParser.AndExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AndExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitAndExpr(GoSubsetParser.AndExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StringExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterStringExpr(GoSubsetParser.StringExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StringExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitStringExpr(GoSubsetParser.StringExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IdentExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterIdentExpr(GoSubsetParser.IdentExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IdentExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitIdentExpr(GoSubsetParser.IdentExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FloatExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterFloatExpr(GoSubsetParser.FloatExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FloatExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitFloatExpr(GoSubsetParser.FloatExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code TrueExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterTrueExpr(GoSubsetParser.TrueExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code TrueExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitTrueExpr(GoSubsetParser.TrueExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code RelExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterRelExpr(GoSubsetParser.RelExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code RelExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitRelExpr(GoSubsetParser.RelExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AddExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterAddExpr(GoSubsetParser.AddExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AddExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitAddExpr(GoSubsetParser.AddExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code UnaryExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterUnaryExpr(GoSubsetParser.UnaryExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code UnaryExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitUnaryExpr(GoSubsetParser.UnaryExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code OrExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterOrExpr(GoSubsetParser.OrExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code OrExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitOrExpr(GoSubsetParser.OrExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IndexExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterIndexExpr(GoSubsetParser.IndexExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IndexExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitIndexExpr(GoSubsetParser.IndexExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FalseExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterFalseExpr(GoSubsetParser.FalseExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FalseExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitFalseExpr(GoSubsetParser.FalseExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SliceLitExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterSliceLitExpr(GoSubsetParser.SliceLitExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SliceLitExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitSliceLitExpr(GoSubsetParser.SliceLitExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SelectorExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterSelectorExpr(GoSubsetParser.SelectorExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SelectorExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitSelectorExpr(GoSubsetParser.SelectorExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code EqExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterEqExpr(GoSubsetParser.EqExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code EqExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitEqExpr(GoSubsetParser.EqExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code CallExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterCallExpr(GoSubsetParser.CallExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code CallExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitCallExpr(GoSubsetParser.CallExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IntExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterIntExpr(GoSubsetParser.IntExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IntExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitIntExpr(GoSubsetParser.IntExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ParenExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterParenExpr(GoSubsetParser.ParenExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ParenExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitParenExpr(GoSubsetParser.ParenExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StructLitExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterStructLitExpr(GoSubsetParser.StructLitExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructLitExpr}
	 * labeled alternative in {@link GoSubsetParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitStructLitExpr(GoSubsetParser.StructLitExprContext ctx);
}