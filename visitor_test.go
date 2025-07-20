package mgs

import (
	"testing"

	"github.com/ajclopez/mgs/v2/parser"
	"github.com/antlr4-go/antlr/v4"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnErrForVisitInputWhenContextIsNil(t *testing.T) {
	caster := map[string]CastType{}
	visitor := NewGeneratorVisitor(queryHandler, &caster)
	result := visitor.VisitInput(nil)

	assert.Equal(t, ErrQueryVisitor, result)
}

func TestShouldReturnErrForVisitPriorityQueryWhenContextIsNil(t *testing.T) {
	caster := map[string]CastType{}
	visitor := NewGeneratorVisitor(queryHandler, &caster)
	result := visitor.VisitPriorityQuery(nil)

	assert.Equal(t, ErrQueryVisitor, result)
}

func TestShouldReturnErrForVisitAtomQueryWhenContextIsNil(t *testing.T) {
	caster := map[string]CastType{}
	visitor := NewGeneratorVisitor(queryHandler, &caster)
	result := visitor.VisitAtomQuery(nil)

	assert.Equal(t, ErrQueryVisitor, result)
}

func TestShouldReturnErrForVisitOpQueryWhenContextIsNil(t *testing.T) {
	caster := map[string]CastType{}
	visitor := NewGeneratorVisitor(queryHandler, &caster)
	result := visitor.VisitOpQuery(nil)

	assert.Equal(t, ErrQueryVisitor, result)
}

func TestShouldReturnErrForVisitCriteriaWhenContextIsNil(t *testing.T) {
	caster := map[string]CastType{}
	visitor := NewGeneratorVisitor(queryHandler, &caster)
	result := visitor.VisitCriteria(nil)

	assert.Equal(t, ErrQueryVisitor, result)
}

func TestShouldReturnErrForVisitCriteriaWhenKeyContextIsNil(t *testing.T) {
	caster := map[string]CastType{}
	visitor := NewGeneratorVisitor(queryHandler, &caster)

	criteriaContext := &parser.CriteriaContext{
		BaseParserRuleContext: *antlr.NewBaseParserRuleContext(nil, -1),
	}
	result := visitor.VisitCriteria(criteriaContext)

	assert.Equal(t, ErrQueryVisitor, result)
}
