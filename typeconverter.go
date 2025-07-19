package mgs

// TypeConverter defines a generic interface for converting string values into specific types, such as ObjectID, Decimal128, DateTime, etc.
// This allows users of the library to provide custom parsers depending on the types and MongoDB driver they are using.
type TypeConverter interface {
	Convert(value string) (interface{}, error)
}
