package mgs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateInstanceFindOpetions(t *testing.T) {

	assert.Equal(t, &FindOptions{}, FindOption())
}

func TestShouldCreateInstanceFindOpetionsWithValues(t *testing.T) {

	caster := map[string]CastType{
		"mobile": STRING,
	}

	opts := FindOption()
	opts.SetCaster(caster)
	opts.SetDefaultLimit(100)
	opts.SetMaxLimit(100)

	dl := int64(100)
	ml := int64(100)
	expected := FindOptions{
		Caster:       &caster,
		DefaultLimit: &dl,
		MaxLimit:     &ml,
	}

	assert.Equal(t, &expected, opts)
}
