package mgs

import (
	"reflect"
)

// Convert converts the criteria value to a MongoDB query
func Convert(criteria SearchCriteria, qh *QueryHandler) map[string]interface{} {

	value := ParseValue(criteria.Value, qh, criteria.Caster)
	filter := make(map[string]interface{})

	switch criteria.Operation {
	case EQUAL:
		typeOf := reflect.ValueOf(value)
		switch {
		case typeOf.Kind() == reflect.Slice:
			filter[criteria.Key] = buildMongoQuery("$in", value)
		case reflect.TypeOf(value).String() == "mgs.Regex":
			filter[criteria.Key] = buildRegexOperation(value)
		default:
			filter[criteria.Key] = value
		}
	case NOT_EQUAL:
		typeOf := reflect.ValueOf(value)
		switch {
		case typeOf.Kind() == reflect.Slice:
			filter[criteria.Key] = buildMongoQuery("$nin", value)
		case reflect.TypeOf(value).String() == "mgs.Regex":
			filter[criteria.Key] = buildMongoQuery("$not", buildRegexOperation(value))
		default:
			filter[criteria.Key] = buildMongoQuery("$ne", value)
		}
	case GREATER_THAN:
		filter[criteria.Key] = buildMongoQuery("$gt", value)
	case GREATER_THAN_EQUAL:
		filter[criteria.Key] = buildMongoQuery("$gte", value)
	case LESS_THAN:
		filter[criteria.Key] = buildMongoQuery("$lt", value)
	case LESS_THAN_EQUAL:
		filter[criteria.Key] = buildMongoQuery("$lte", value)
	case EXISTS:
		filter[criteria.Key] = buildMongoQuery("$exists", !criteria.Prefix)
	}

	return filter
}

func buildMongoQuery(operator string, value interface{}) map[string]interface{} {
	query := make(map[string]interface{})
	query[operator] = value

	return query
}

func buildRegexOperation(value interface{}) map[string]interface{} {
	regex := make(map[string]interface{})
	regex["$regex"] = value.(Regex).Pattern
	regex["$options"] = value.(Regex).Option

	return regex
}
