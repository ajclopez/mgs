package mgs

// Primitives defines methods to convert to specific types.
type Primitives interface {
	ObjectID(val string) (oid interface{}, err error)
}
