package mgs

import (
	"errors"
	"net/url"
	"strings"

	"github.com/ajclopez/mgs/parser"
	"github.com/antlr4-go/antlr/v4"
)

var (
	// ErrUnescapeCharacters is returned when cannot converter an unescape string
	ErrUnescapeCharacters = errors.New("unescape characters")
	// ErrValueNoMatch is returned when the converter cannot match a string to
	// an unsigned integer.
	ErrValueNoMatch = errors.New("value does not match")
	// ErrQueryVisitor
	ErrQueryVisitor = errors.New("context does not exists")
)

// NewQueryHandler creates a new QueryHandler instance.
func NewQueryHandler(converter TypeConverter) *QueryHandler {
	return &QueryHandler{
		Converter: converter,
	}
}

// MongoGoSearch converts query into a MongoDB query object.
func (qh *QueryHandler) MongoGoSearch(query string, opts *FindOptions) (Query, error) {

	res := Query{}

	if strings.TrimSpace(query) == "" {
		return res, nil
	}

	var err error
	query, err = url.QueryUnescape(strings.Replace(query, "+", "%2B", -1))

	if err != nil {
		return res, ErrUnescapeCharacters
	}

	var filterAdvanced map[string][]interface{}
	var filters []interface{}

	for _, criteria := range Parse(query, opts.Caster) {
		switch criteria.Key {
		case "filter":
			filterAdvanced = parseFilterAdvanced(criteria.Value, qh, opts.Caster)
		case "skip":
			err = parseSkip(&res, criteria.Value)
		case "limit":
			err = parseLimit(&res, criteria.Value, opts)
		case "sort":
			parseSort(&res, criteria.Value)
		case "fields":
			parseFields(&res, criteria.Value)
		default:
			filters = append(filters, Convert(criteria, qh))
		}

		if err != nil {
			return res, err
		}
	}

	parseDefaultFilter(&res, filters, filterAdvanced)

	return res, err
}

func parseDefaultFilter(res *Query, filters []interface{}, filterAdvanced map[string][]interface{}) {
	if filterAdvanced != nil {
		res.Filter = filterAdvanced
	} else {
		if filters == nil {
			res.Filter = map[string][]interface{}{}
		} else {
			res.Filter = map[string][]interface{}{
				"$and": filters,
			}
		}
	}
}

func parseSkip(res *Query, value string) error {
	skip, err := parseIntValueToInt(value)
	if err != nil {
		return err
	}

	res.Skip = skip
	return nil
}

func parseLimit(res *Query, value string, opts *FindOptions) error {
	limit, err := parseIntValueToInt(value)

	if err != nil {
		if opts.DefaultLimit != nil {
			res.Limit = *opts.DefaultLimit
			return nil
		} else {
			return err
		}
	}

	maxLimit := opts.MaxLimit
	if maxLimit != nil && limit > *maxLimit {
		res.Limit = *maxLimit
	} else {
		res.Limit = limit
	}

	return nil
}

func parseSort(res *Query, value string) {
	sort := make(map[string]int)
	values := strings.Split(value, ",")

	for i := range values {
		if strings.HasPrefix(values[i], "+") {
			sort[values[i][1:]] = 1
		} else if strings.HasPrefix(values[i], "-") {
			sort[values[i][1:]] = -1
		} else {
			sort[values[i]] = 1
		}
	}

	res.Sort = sort
}

func parseFields(res *Query, value string) {
	projection := make(map[string]int)
	values := strings.Split(value, ",")
	for i := range values {
		projection[values[i]] = 1
	}

	res.Projection = projection
}

func parseFilterAdvanced(value string, qh *QueryHandler, caster *map[string]CastType) map[string][]interface{} {

	is := antlr.NewInputStream(value)
	lexer := parser.NewQueryLexer(is)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := parser.NewQueryParser(tokens)
	visitor := NewGeneratorVisitor(qh, caster)

	tree := parser.Input()
	result := visitor.Visit(tree)

	if _, ok := result.(map[string]interface{}); ok {
		return map[string][]interface{}{
			"$and": {
				result.(map[string]interface{}),
			},
		}
	} else {
		return result.(map[string][]interface{})
	}
}
