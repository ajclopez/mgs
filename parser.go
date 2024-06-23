package mgs

import (
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
		if b, err := strconv.ParseBool(value); err == nil {
			return b
		}
	case DATE:
		if datetime, err := time.Parse(DATE_FORMAT, value); err == nil {
			return datetime
		}
	case NUMBER:
		if integer, err := strconv.ParseInt(value, 10, 64); err == nil {
			return integer
		}

		if float, err := strconv.ParseFloat(value, 64); err == nil {
			return float
		}
	case PATTERN:
		if match := GetRegexPattern().FindStringSubmatch(value); match != nil {
			return Regex{
				Pattern: match[GetRegexPattern().SubexpIndex("Pattern")],
				Option:  match[GetRegexPattern().SubexpIndex("Option")],
			}
		}
	case STRING:
		return value
	}

	return parseValue(value, qh)
}

func parseValue(value string, qh *QueryHandler) interface{} {

	if strings.EqualFold(value, "true") || strings.EqualFold(value, "false") {
		if b, err := strconv.ParseBool(strings.ToLower(value)); err == nil {
			return b
		}
	}

	var err error

	var integer int64
	if integer, err = strconv.ParseInt(value, 10, 64); err == nil {
		return integer
	}

	var float float64
	if float, err = strconv.ParseFloat(value, 64); err == nil {
		return float
	}

	var datetime time.Time
	if datetime, err = time.Parse(DATE_FORMAT, value); err == nil {
		return datetime
	}

	list := strings.Split(value, ",")
	if len(list) > 1 {
		var characters []interface{}
		for _, _value := range list {
			characters = append(characters, parseValue(_value, qh))
		}
		if len(characters) > 0 {
			return characters
		}
	}

	if match := GetRegexPattern().FindStringSubmatch(value); match != nil {
		return Regex{
			Pattern: match[GetRegexPattern().SubexpIndex("Pattern")],
			Option:  match[GetRegexPattern().SubexpIndex("Option")],
		}
	}

	objectId, err := qh.Primitives.ObjectID(value)
	if err == nil {
		return objectId
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
