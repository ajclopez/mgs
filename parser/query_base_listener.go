// Code generated from Query.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Query
import "github.com/antlr4-go/antlr/v4"

// BaseQueryListener is a complete listener for a parse tree produced by QueryParser.
type BaseQueryListener struct{}

var _ QueryListener = &BaseQueryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseQueryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseQueryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseQueryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseQueryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterInput is called when production input is entered.
func (s *BaseQueryListener) EnterInput(ctx *InputContext) {}

// ExitInput is called when production input is exited.
func (s *BaseQueryListener) ExitInput(ctx *InputContext) {}

// EnterAtomQuery is called when production atomQuery is entered.
func (s *BaseQueryListener) EnterAtomQuery(ctx *AtomQueryContext) {}

// ExitAtomQuery is called when production atomQuery is exited.
func (s *BaseQueryListener) ExitAtomQuery(ctx *AtomQueryContext) {}

// EnterPriorityQuery is called when production priorityQuery is entered.
func (s *BaseQueryListener) EnterPriorityQuery(ctx *PriorityQueryContext) {}

// ExitPriorityQuery is called when production priorityQuery is exited.
func (s *BaseQueryListener) ExitPriorityQuery(ctx *PriorityQueryContext) {}

// EnterOpQuery is called when production opQuery is entered.
func (s *BaseQueryListener) EnterOpQuery(ctx *OpQueryContext) {}

// ExitOpQuery is called when production opQuery is exited.
func (s *BaseQueryListener) ExitOpQuery(ctx *OpQueryContext) {}

// EnterCriteria is called when production criteria is entered.
func (s *BaseQueryListener) EnterCriteria(ctx *CriteriaContext) {}

// ExitCriteria is called when production criteria is exited.
func (s *BaseQueryListener) ExitCriteria(ctx *CriteriaContext) {}

// EnterKey is called when production key is entered.
func (s *BaseQueryListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BaseQueryListener) ExitKey(ctx *KeyContext) {}

// EnterValue is called when production value is entered.
func (s *BaseQueryListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseQueryListener) ExitValue(ctx *ValueContext) {}

// EnterOp is called when production op is entered.
func (s *BaseQueryListener) EnterOp(ctx *OpContext) {}

// ExitOp is called when production op is exited.
func (s *BaseQueryListener) ExitOp(ctx *OpContext) {}
