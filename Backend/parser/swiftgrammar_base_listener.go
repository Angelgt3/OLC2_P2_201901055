// Code generated from SwiftGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// BaseSwiftGrammarListener is a complete listener for a parse tree produced by SwiftGrammarParser.
type BaseSwiftGrammarListener struct{}

var _ SwiftGrammarListener = &BaseSwiftGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSwiftGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSwiftGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSwiftGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSwiftGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterS is called when production s is entered.
func (s *BaseSwiftGrammarListener) EnterS(ctx *SContext) {}

// ExitS is called when production s is exited.
func (s *BaseSwiftGrammarListener) ExitS(ctx *SContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseSwiftGrammarListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseSwiftGrammarListener) ExitBlock(ctx *BlockContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseSwiftGrammarListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseSwiftGrammarListener) ExitInstruction(ctx *InstructionContext) {}

// EnterStructCreation is called when production structCreation is entered.
func (s *BaseSwiftGrammarListener) EnterStructCreation(ctx *StructCreationContext) {}

// ExitStructCreation is called when production structCreation is exited.
func (s *BaseSwiftGrammarListener) ExitStructCreation(ctx *StructCreationContext) {}

// EnterListStructDec is called when production listStructDec is entered.
func (s *BaseSwiftGrammarListener) EnterListStructDec(ctx *ListStructDecContext) {}

// ExitListStructDec is called when production listStructDec is exited.
func (s *BaseSwiftGrammarListener) ExitListStructDec(ctx *ListStructDecContext) {}

// EnterFuncstmt is called when production funcstmt is entered.
func (s *BaseSwiftGrammarListener) EnterFuncstmt(ctx *FuncstmtContext) {}

// ExitFuncstmt is called when production funcstmt is exited.
func (s *BaseSwiftGrammarListener) ExitFuncstmt(ctx *FuncstmtContext) {}

// EnterListParamsFunc is called when production listParamsFunc is entered.
func (s *BaseSwiftGrammarListener) EnterListParamsFunc(ctx *ListParamsFuncContext) {}

// ExitListParamsFunc is called when production listParamsFunc is exited.
func (s *BaseSwiftGrammarListener) ExitListParamsFunc(ctx *ListParamsFuncContext) {}

// EnterFuncallstmt is called when production funcallstmt is entered.
func (s *BaseSwiftGrammarListener) EnterFuncallstmt(ctx *FuncallstmtContext) {}

// ExitFuncallstmt is called when production funcallstmt is exited.
func (s *BaseSwiftGrammarListener) ExitFuncallstmt(ctx *FuncallstmtContext) {}

// EnterFuncallestmt is called when production funcallestmt is entered.
func (s *BaseSwiftGrammarListener) EnterFuncallestmt(ctx *FuncallestmtContext) {}

// ExitFuncallestmt is called when production funcallestmt is exited.
func (s *BaseSwiftGrammarListener) ExitFuncallestmt(ctx *FuncallestmtContext) {}

// EnterPrintstmt is called when production printstmt is entered.
func (s *BaseSwiftGrammarListener) EnterPrintstmt(ctx *PrintstmtContext) {}

// ExitPrintstmt is called when production printstmt is exited.
func (s *BaseSwiftGrammarListener) ExitPrintstmt(ctx *PrintstmtContext) {}

// EnterVariablestmt is called when production variablestmt is entered.
func (s *BaseSwiftGrammarListener) EnterVariablestmt(ctx *VariablestmtContext) {}

// ExitVariablestmt is called when production variablestmt is exited.
func (s *BaseSwiftGrammarListener) ExitVariablestmt(ctx *VariablestmtContext) {}

// EnterTypestmt is called when production typestmt is entered.
func (s *BaseSwiftGrammarListener) EnterTypestmt(ctx *TypestmtContext) {}

// ExitTypestmt is called when production typestmt is exited.
func (s *BaseSwiftGrammarListener) ExitTypestmt(ctx *TypestmtContext) {}

// EnterType2stmt is called when production type2stmt is entered.
func (s *BaseSwiftGrammarListener) EnterType2stmt(ctx *Type2stmtContext) {}

// ExitType2stmt is called when production type2stmt is exited.
func (s *BaseSwiftGrammarListener) ExitType2stmt(ctx *Type2stmtContext) {}

// EnterFunvecstmt is called when production funvecstmt is entered.
func (s *BaseSwiftGrammarListener) EnterFunvecstmt(ctx *FunvecstmtContext) {}

// ExitFunvecstmt is called when production funvecstmt is exited.
func (s *BaseSwiftGrammarListener) ExitFunvecstmt(ctx *FunvecstmtContext) {}

// EnterGuardstmt is called when production guardstmt is entered.
func (s *BaseSwiftGrammarListener) EnterGuardstmt(ctx *GuardstmtContext) {}

// ExitGuardstmt is called when production guardstmt is exited.
func (s *BaseSwiftGrammarListener) ExitGuardstmt(ctx *GuardstmtContext) {}

// EnterIfstmt is called when production ifstmt is entered.
func (s *BaseSwiftGrammarListener) EnterIfstmt(ctx *IfstmtContext) {}

// ExitIfstmt is called when production ifstmt is exited.
func (s *BaseSwiftGrammarListener) ExitIfstmt(ctx *IfstmtContext) {}

// EnterElifs is called when production elifs is entered.
func (s *BaseSwiftGrammarListener) EnterElifs(ctx *ElifsContext) {}

// ExitElifs is called when production elifs is exited.
func (s *BaseSwiftGrammarListener) ExitElifs(ctx *ElifsContext) {}

// EnterSwitchstmt is called when production switchstmt is entered.
func (s *BaseSwiftGrammarListener) EnterSwitchstmt(ctx *SwitchstmtContext) {}

// ExitSwitchstmt is called when production switchstmt is exited.
func (s *BaseSwiftGrammarListener) ExitSwitchstmt(ctx *SwitchstmtContext) {}

// EnterCases is called when production cases is entered.
func (s *BaseSwiftGrammarListener) EnterCases(ctx *CasesContext) {}

// ExitCases is called when production cases is exited.
func (s *BaseSwiftGrammarListener) ExitCases(ctx *CasesContext) {}

// EnterWhilestmt is called when production whilestmt is entered.
func (s *BaseSwiftGrammarListener) EnterWhilestmt(ctx *WhilestmtContext) {}

// ExitWhilestmt is called when production whilestmt is exited.
func (s *BaseSwiftGrammarListener) ExitWhilestmt(ctx *WhilestmtContext) {}

// EnterForstmt is called when production forstmt is entered.
func (s *BaseSwiftGrammarListener) EnterForstmt(ctx *ForstmtContext) {}

// ExitForstmt is called when production forstmt is exited.
func (s *BaseSwiftGrammarListener) ExitForstmt(ctx *ForstmtContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSwiftGrammarListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSwiftGrammarListener) ExitExpr(ctx *ExprContext) {}

// EnterPrimitives is called when production primitives is entered.
func (s *BaseSwiftGrammarListener) EnterPrimitives(ctx *PrimitivesContext) {}

// ExitPrimitives is called when production primitives is exited.
func (s *BaseSwiftGrammarListener) ExitPrimitives(ctx *PrimitivesContext) {}

// EnterListParams is called when production listParams is entered.
func (s *BaseSwiftGrammarListener) EnterListParams(ctx *ListParamsContext) {}

// ExitListParams is called when production listParams is exited.
func (s *BaseSwiftGrammarListener) ExitListParams(ctx *ListParamsContext) {}

// EnterListArray is called when production listArray is entered.
func (s *BaseSwiftGrammarListener) EnterListArray(ctx *ListArrayContext) {}

// ExitListArray is called when production listArray is exited.
func (s *BaseSwiftGrammarListener) ExitListArray(ctx *ListArrayContext) {}

// EnterListAceso is called when production listAceso is entered.
func (s *BaseSwiftGrammarListener) EnterListAceso(ctx *ListAcesoContext) {}

// ExitListAceso is called when production listAceso is exited.
func (s *BaseSwiftGrammarListener) ExitListAceso(ctx *ListAcesoContext) {}

// EnterListStructExp is called when production listStructExp is entered.
func (s *BaseSwiftGrammarListener) EnterListStructExp(ctx *ListStructExpContext) {}

// ExitListStructExp is called when production listStructExp is exited.
func (s *BaseSwiftGrammarListener) ExitListStructExp(ctx *ListStructExpContext) {}
