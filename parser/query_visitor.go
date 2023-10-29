// Code generated from Query.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Query
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by QueryParser.
type QueryVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by QueryParser#input.
	VisitInput(ctx *InputContext) interface{}

	// Visit a parse tree produced by QueryParser#atomQuery.
	VisitAtomQuery(ctx *AtomQueryContext) interface{}

	// Visit a parse tree produced by QueryParser#priorityQuery.
	VisitPriorityQuery(ctx *PriorityQueryContext) interface{}

	// Visit a parse tree produced by QueryParser#opQuery.
	VisitOpQuery(ctx *OpQueryContext) interface{}

	// Visit a parse tree produced by QueryParser#criteria.
	VisitCriteria(ctx *CriteriaContext) interface{}

	// Visit a parse tree produced by QueryParser#key.
	VisitKey(ctx *KeyContext) interface{}

	// Visit a parse tree produced by QueryParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by QueryParser#op.
	VisitOp(ctx *OpContext) interface{}
}
