// Code generated from SwiftGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// SwiftGrammarListener is a complete listener for a parse tree produced by SwiftGrammarParser.
type SwiftGrammarListener interface {
	antlr.ParseTreeListener

	// EnterS is called when entering the s production.
	EnterS(c *SContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterStructCreation is called when entering the structCreation production.
	EnterStructCreation(c *StructCreationContext)

	// EnterListStructDec is called when entering the listStructDec production.
	EnterListStructDec(c *ListStructDecContext)

	// EnterFuncstmt is called when entering the funcstmt production.
	EnterFuncstmt(c *FuncstmtContext)

	// EnterListParamsFunc is called when entering the listParamsFunc production.
	EnterListParamsFunc(c *ListParamsFuncContext)

	// EnterFuncallstmt is called when entering the funcallstmt production.
	EnterFuncallstmt(c *FuncallstmtContext)

	// EnterFuncallestmt is called when entering the funcallestmt production.
	EnterFuncallestmt(c *FuncallestmtContext)

	// EnterPrintstmt is called when entering the printstmt production.
	EnterPrintstmt(c *PrintstmtContext)

	// EnterVariablestmt is called when entering the variablestmt production.
	EnterVariablestmt(c *VariablestmtContext)

	// EnterTypestmt is called when entering the typestmt production.
	EnterTypestmt(c *TypestmtContext)

	// EnterType2stmt is called when entering the type2stmt production.
	EnterType2stmt(c *Type2stmtContext)

	// EnterFunvecstmt is called when entering the funvecstmt production.
	EnterFunvecstmt(c *FunvecstmtContext)

	// EnterGuardstmt is called when entering the guardstmt production.
	EnterGuardstmt(c *GuardstmtContext)

	// EnterIfstmt is called when entering the ifstmt production.
	EnterIfstmt(c *IfstmtContext)

	// EnterElifs is called when entering the elifs production.
	EnterElifs(c *ElifsContext)

	// EnterSwitchstmt is called when entering the switchstmt production.
	EnterSwitchstmt(c *SwitchstmtContext)

	// EnterCases is called when entering the cases production.
	EnterCases(c *CasesContext)

	// EnterWhilestmt is called when entering the whilestmt production.
	EnterWhilestmt(c *WhilestmtContext)

	// EnterForstmt is called when entering the forstmt production.
	EnterForstmt(c *ForstmtContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterPrimitives is called when entering the primitives production.
	EnterPrimitives(c *PrimitivesContext)

	// EnterListParams is called when entering the listParams production.
	EnterListParams(c *ListParamsContext)

	// EnterListArray is called when entering the listArray production.
	EnterListArray(c *ListArrayContext)

	// EnterListAceso is called when entering the listAceso production.
	EnterListAceso(c *ListAcesoContext)

	// EnterListStructExp is called when entering the listStructExp production.
	EnterListStructExp(c *ListStructExpContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitStructCreation is called when exiting the structCreation production.
	ExitStructCreation(c *StructCreationContext)

	// ExitListStructDec is called when exiting the listStructDec production.
	ExitListStructDec(c *ListStructDecContext)

	// ExitFuncstmt is called when exiting the funcstmt production.
	ExitFuncstmt(c *FuncstmtContext)

	// ExitListParamsFunc is called when exiting the listParamsFunc production.
	ExitListParamsFunc(c *ListParamsFuncContext)

	// ExitFuncallstmt is called when exiting the funcallstmt production.
	ExitFuncallstmt(c *FuncallstmtContext)

	// ExitFuncallestmt is called when exiting the funcallestmt production.
	ExitFuncallestmt(c *FuncallestmtContext)

	// ExitPrintstmt is called when exiting the printstmt production.
	ExitPrintstmt(c *PrintstmtContext)

	// ExitVariablestmt is called when exiting the variablestmt production.
	ExitVariablestmt(c *VariablestmtContext)

	// ExitTypestmt is called when exiting the typestmt production.
	ExitTypestmt(c *TypestmtContext)

	// ExitType2stmt is called when exiting the type2stmt production.
	ExitType2stmt(c *Type2stmtContext)

	// ExitFunvecstmt is called when exiting the funvecstmt production.
	ExitFunvecstmt(c *FunvecstmtContext)

	// ExitGuardstmt is called when exiting the guardstmt production.
	ExitGuardstmt(c *GuardstmtContext)

	// ExitIfstmt is called when exiting the ifstmt production.
	ExitIfstmt(c *IfstmtContext)

	// ExitElifs is called when exiting the elifs production.
	ExitElifs(c *ElifsContext)

	// ExitSwitchstmt is called when exiting the switchstmt production.
	ExitSwitchstmt(c *SwitchstmtContext)

	// ExitCases is called when exiting the cases production.
	ExitCases(c *CasesContext)

	// ExitWhilestmt is called when exiting the whilestmt production.
	ExitWhilestmt(c *WhilestmtContext)

	// ExitForstmt is called when exiting the forstmt production.
	ExitForstmt(c *ForstmtContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitPrimitives is called when exiting the primitives production.
	ExitPrimitives(c *PrimitivesContext)

	// ExitListParams is called when exiting the listParams production.
	ExitListParams(c *ListParamsContext)

	// ExitListArray is called when exiting the listArray production.
	ExitListArray(c *ListArrayContext)

	// ExitListAceso is called when exiting the listAceso production.
	ExitListAceso(c *ListAcesoContext)

	// ExitListStructExp is called when exiting the listStructExp production.
	ExitListStructExp(c *ListStructExpContext)
}
