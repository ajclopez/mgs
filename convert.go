package mgs

import (
	"reflect"
)

// Convert converts the criteria value to a MongoDB query
func Convert(criteria SearchCriteria) map[string]interface{} {

	value := ParseValue(criteria.Value, criteria.Caster)
	filter := make(map[string]interface{})

	switch criteria.Operation {
	case EQUAL:
		key := reflect.ValueOf(value).Kind()
		if key == reflect.Slice {
			filter[criteria.Key] = buildMongoQuery("$in", value)
		} else if key == reflect.Struct {
			filter[criteria.Key] = buildRegexOperation(value)
		} else {
			filter[criteria.Key] = value
		}
	case NOT_EQUAL:
		key := reflect.ValueOf(value).Kind()
		if key == reflect.Slice {
			filter[criteria.Key] = buildMongoQuery("$nin", value)
		} else if key == reflect.Struct {
			filter[criteria.Key] = buildMongoQuery("$not", buildRegexOperation(value))
		} else {
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
