package mgs

type CastType int

// list of allowed cast type.
const (
	NUMBER CastType = iota + 1
	DATE
	BOOLEAN
	PATTERN
	STRING
)

// FindOptions is a structure that allows to use advanced options.
type FindOptions struct {
	// Caster map to set object type on values by key (BOOLEAN, NUMBER, PATTERN, DATE, STRING).
	Caster *map[string]CastType
	// DefaultLimit default value for limit key.
	DefaultLimit *int64
	// MaxLimit maximum value for limit key.
	MaxLimit *int64
}

// FindOption creates a new FindOptions instance.
func FindOption() *FindOptions {
	return &FindOptions{}
}

// SetCaster sets the value for the object type on values by key.
func (f *FindOptions) SetCaster(m map[string]CastType) *FindOptions {
	f.Caster = &m
	return f
}

// SetDefaultLimit sets the value for the default value for limit key.
func (f *FindOptions) SetDefaultLimit(i int64) *FindOptions {
	f.DefaultLimit = &i
	return f
}

// SetMaxLimit sets the value for the maximum value for limit key.
func (f *FindOptions) SetMaxLimit(i int64) *FindOptions {
	f.MaxLimit = &i
	return f
}
