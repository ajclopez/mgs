// Code generated from Query.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Query
import "github.com/antlr4-go/antlr/v4"

// QueryListener is a complete listener for a parse tree produced by QueryParser.
type QueryListener interface {
	antlr.ParseTreeListener

	// EnterInput is called when entering the input production.
	EnterInput(c *InputContext)

	// EnterAtomQuery is called when entering the atomQuery production.
	EnterAtomQuery(c *AtomQueryContext)

	// EnterPriorityQuery is called when entering the priorityQuery production.
	EnterPriorityQuery(c *PriorityQueryContext)

	// EnterOpQuery is called when entering the opQuery production.
	EnterOpQuery(c *OpQueryContext)

	// EnterCriteria is called when entering the criteria production.
	EnterCriteria(c *CriteriaContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterOp is called when entering the op production.
	EnterOp(c *OpContext)

	// ExitInput is called when exiting the input production.
	ExitInput(c *InputContext)

	// ExitAtomQuery is called when exiting the atomQuery production.
	ExitAtomQuery(c *AtomQueryContext)

	// ExitPriorityQuery is called when exiting the priorityQuery production.
	ExitPriorityQuery(c *PriorityQueryContext)

	// ExitOpQuery is called when exiting the opQuery production.
	ExitOpQuery(c *OpQueryContext)

	// ExitCriteria is called when exiting the criteria production.
	ExitCriteria(c *CriteriaContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitOp is called when exiting the op production.
	ExitOp(c *OpContext)
}
