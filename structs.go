package mgs

// Query is a structure that holds information about DB request.
type Query struct {
	// Filter is a document containing query operators.
	Filter map[string][]interface{}
	// Sort is a document specifying the order in which documents should be returned.
	Sort map[string]int
	// Limit is the maximum number of documents to return.
	Limit int64
	// Skip is a number of documents to be skipped.
	Skip int64
	// Projection is the limit fields to return in each records.
	Projection map[string]int
}

// SearchCriteria is a structure that holds search criteria specification.
type SearchCriteria struct {
	// Prefix is a bool that allows to check that the search criteria contain prefixes.
	Prefix bool
	// key is a string that allows the key to be identified in the search criteria.
	Key string
	// Operation is a structure that containing operator type.
	Operation SearchOperator
	// Value is a string that allows the value to be identified in the search criteria.
	Value string
	// Caster is a cast type value.
	Caster *CastType
}

// Regex is a structure that holds RegExp specification.
type Regex struct {
	// Pattern is a string contains regex pattern should be used.
	Pattern string
	// Opetion is a string contains regex options should be used.
	Option string
}
