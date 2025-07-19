package mgs

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const DATE_FORMAT = "2006-01-02T15:04:05.000Z"

// Parse parses a given url query.
func Parse(query string, caster *map[string]CastType) []SearchCriteria {
	criteria := []SearchCriteria{}
	for _, condition := range strings.Split(query, "&") {
		criteria = append(criteria, CriteriaParser(condition, caster))
	}
	return criteria
}

// CriteriaParser build criteria from a condition expression
func CriteriaParser(condition string, caster *map[string]CastType) SearchCriteria {

	pattern := GetOperatorPattern()
	match := pattern.FindStringSubmatch(condition)

	key := match[pattern.SubexpIndex("Key")]
	var value CastType
	if caster != nil {
		value = (*caster)[key]
	}

	return SearchCriteria{
		Prefix:    match[pattern.SubexpIndex("Prefix")] != "",
		Key:       key,
		Operation: SearchOperator(1).GetOperation(match[pattern.SubexpIndex("Operator")]),
		Value:     match[pattern.SubexpIndex("Value")],
		Caster:    &value,
	}
}

// ParseValue converts a string to a data type
func ParseValue(value string, qh *QueryHandler, cast *CastType) interface{} {

	if cast == nil {
		return parseValue(value, qh)
	}

	switch option := *cast; option {
	case BOOLEAN:
		if b, err := parseBool(value); err == nil {
			return b
		}
	case DATE:
		if datetime, err := parseDateTime(value); err == nil {
			return datetime
		}
	case NUMBER:
		if integer, err := parseInt(value); err == nil {
			return integer
		}

		if float, err := parseFloat(value); err == nil {
			return float
		}
	case PATTERN:
		if match := parseRegex(value); match != nil {
			return match
		}
	case STRING:
		return value
	}

	return parseValue(value, qh)
}

func parseValue(value string, qh *QueryHandler) interface{} {

	if result := parseConverter(value, qh); result != nil {
		return result
	}

	if b, err := parseBool(value); err == nil {
		return b
	}

	if integer, err := parseInt(value); err == nil {
		return integer
	}

	if float, err := parseFloat(value); err == nil {
		return float
	}

	if datetime, err := parseDateTime(value); err == nil {
		return datetime
	}

	if list := parseList(value, qh); list != nil {
		return list
	}

	if match := parseRegex(value); match != nil {
		return match
	}

	return value
}

func parseIntValueToInt(value string) (int64, error) {
	result, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, ErrValueNoMatch
	}
	return result, nil
}

func parseBool(value string) (bool, error) {
	lower := strings.ToLower(value)
	if lower == "true" || lower == "false" {
		b, err := strconv.ParseBool(lower)
		return b, err
	}
	return false, fmt.Errorf("invalid boolean value: %s", value)
}

func parseInt(value string) (int64, error) {
	i, err := strconv.ParseInt(value, 10, 64)
	return i, err
}

func parseFloat(value string) (float64, error) {
	f, err := strconv.ParseFloat(value, 64)
	return f, err
}

func parseDateTime(value string) (time.Time, error) {
	datetime, err := time.Parse(DATE_FORMAT, value)
	if err == nil {
		return datetime, err
	}
	return datetime, err
}

func parseList(value string, qh *QueryHandler) []interface{} {
	parts := strings.Split(value, ",")
	if len(parts) > 1 {
		result := make([]interface{}, 0, len(parts))
		for _, p := range parts {
			result = append(result, parseValue(p, qh))
		}
		if len(result) > 0 {
			return result
		}
	}
	return nil
}

func parseRegex(value string) interface{} {
	match := GetRegexPattern().FindStringSubmatch(value)
	if match != nil {
		return Regex{
			Pattern: match[GetRegexPattern().SubexpIndex("Pattern")],
			Option:  match[GetRegexPattern().SubexpIndex("Option")],
		}
	}
	return nil
}

func parseConverter(value string, qh *QueryHandler) interface{} {
	if result, err := qh.Converter.Convert(value); err == nil {
		return result
	}

	return nil
}
