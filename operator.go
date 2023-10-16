package mgs

type SearchOperator int

// list of allowed operators.
const (
	EQUAL SearchOperator = iota + 1
	NOT_EQUAL
	GREATER_THAN
	GREATER_THAN_EQUAL
	LESS_THAN
	LESS_THAN_EQUAL
	EXISTS
)

// String gets operator string.
func (s SearchOperator) String() string {
	return [...]string{"EQUAL", "NOT_EQUAL", "GREATER_THAN", "GREATER_THAN_EQUAL", "LESS_THAN", "LESS_THAN_EQUAL", "EXISTS"}[s-1]
}

// GetOperation allows get operation type to filter query
func (s SearchOperator) GetOperation(input string) SearchOperator {
	switch input {
	case "=":
		return EQUAL
	case "!=":
		return NOT_EQUAL
	case ">":
		return GREATER_THAN
	case ">=":
		return GREATER_THAN_EQUAL
	case "<":
		return LESS_THAN
	case "<=":
		return LESS_THAN_EQUAL
	default:
		return EXISTS
	}
}
