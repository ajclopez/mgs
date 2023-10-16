package mgs

import (
	"errors"
	"net/url"
	"strings"
)

var (
	// ErrUnescapeCharacters is returned when cannot converter an unescape string
	ErrUnescapeCharacters = errors.New("unescape characters")
	// ErrValueNoMatch is returned when the converter cannot match a string to
	// an unsigned integer.
	ErrValueNoMatch = errors.New("value does not match")
)

// MongoGoSearch converts query into a MongoDB query object.
func MongoGoSearch(query string, opts *FindOptions) (Query, error) {

	res := Query{}

	if strings.TrimSpace(query) == "" {
		return res, nil
	}

	var err error
	query, err = url.QueryUnescape(strings.Replace(query, "+", "%2B", -1))

	if err != nil {
		return res, ErrUnescapeCharacters
	}

	filter := make(map[string]interface{})

	for _, criteria := range Parse(query, opts.Caster) {
		switch criteria.Key {
		case "filter":
			// advanced queries
		case "skip":
			err = parseSkip(&res, criteria.Value)
		case "limit":
			err = parseLimit(&res, criteria.Value, opts)
		case "sort":
			parseSort(&res, criteria.Value)
		case "fields":
			parseFields(&res, criteria.Value)
		default:
			Convert(criteria, filter)
		}

		if err != nil {
			return res, err
		}
	}

	res.Filter = filter

	return res, err
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
