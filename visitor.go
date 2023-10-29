package mgs

import (
	"fmt"
	"strings"

	"github.com/ajclopez/mgs/parser"
	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	*parser.BaseQueryVisitor
	Caster *map[string]CastType
}

func NewGeneratorVisitor(caster *map[string]CastType) *Visitor {
	return &Visitor{
		Caster: caster,
	}
}

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *Visitor) VisitInput(ctx *parser.InputContext) interface{} {
	if ctx == nil {
		return ErrQueryVisitor
	}
	return v.Visit(ctx.Query())
}

func (v *Visitor) VisitPriorityQuery(ctx *parser.PriorityQueryContext) interface{} {
	if ctx == nil {
		return ErrQueryVisitor
	}
	return v.Visit(ctx.Query())
}

func (v *Visitor) VisitAtomQuery(ctx *parser.AtomQueryContext) interface{} {
	if ctx == nil {
		return ErrQueryVisitor
	}
	return v.Visit(ctx.Criteria())
}

func (v *Visitor) VisitOpQuery(ctx *parser.OpQueryContext) interface{} {

	if ctx == nil {
		return ErrQueryVisitor
	}

	left := v.Visit(ctx.GetLeft())
	right := v.Visit(ctx.GetRight())
	op := ctx.GetLogicalOp().GetText()
	var operator string

	switch op {
	case "and", "AND":
		operator = "$and"
	case "or", "OR":
		operator = "$or"
	default:
		operator = "$and"
	}

	return map[string][]interface{}{
		operator: {
			left,
			right,
		},
	}
}

func (v *Visitor) VisitCriteria(ctx *parser.CriteriaContext) interface{} {

	if ctx == nil {
		return ErrQueryVisitor
	}

	key := ctx.Key()
	if key == nil {
		return ErrQueryVisitor
	}

	op := ctx.Op()
	if op == nil {
		return ErrQueryVisitor
	}

	value := ctx.Value()
	if value == nil {
		return ErrQueryVisitor
	}

	expression := fmt.Sprintf("%s%s%s", strings.TrimSpace(key.GetText()), strings.TrimSpace(op.GetText()), strings.TrimSpace(value.GetText()))
	criteria := CriteriaParser(expression, v.Caster)

	return Convert(criteria)
}
