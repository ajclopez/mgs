package mgs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var converetTests = []struct {
	Criteria SearchCriteria
	Expected map[string]interface{}
}{
	{
		SearchCriteria{
			Prefix:    false,
			Key:       "name",
			Operation: EQUAL,
			Value:     "Jhon",
		},
		map[string]interface{}{"name": "Jhon"},
	},
	{
		SearchCriteria{
			Prefix:    false,
			Key:       "status",
			Operation: EQUAL,
			Value:     "QUEUED,DEQUEUED",
		},
		map[string]interface{}{"status": map[string]interface{}{"$in": []interface{}{"QUEUED", "DEQUEUED"}}},
	},
	{
		SearchCriteria{
			Prefix:    false,
			Key:       "email",
			Operation: EQUAL,
			Value:     "/@gmail\\.com$/",
		},
		map[string]interface{}{"email": map[string]interface{}{"$regex": "@gmail\\.com$", "$options": ""}},
	},
	{
		SearchCriteria{
			Prefix:    false,
			Key:       "status",
			Operation: NOT_EQUAL,
			Value:     "SENT",
		},
		map[string]interface{}{"status": map[string]interface{}{"$ne": "SENT"}},
	},
	{
		SearchCriteria{
			Prefix:    false,
			Key:       "status",
			Operation: NOT_EQUAL,
			Value:     "QUEUED,DEQUEUED",
		},
		map[string]interface{}{"status": map[string]interface{}{"$nin": []interface{}{"QUEUED", "DEQUEUED"}}},
	},
	{
		SearchCriteria{
			Prefix:    false,
			Key:       "phone",
			Operation: NOT_EQUAL,
			Value:     "/^58/",
		},
		map[string]interface{}{"phone": map[string]interface{}{"$not": map[string]interface{}{"$regex": "^58", "$options": ""}}},
	},
	{
		SearchCriteria{
			Prefix:    false,
			Key:       "price",
			Operation: GREATER_THAN,
			Value:     "5",
		},
		map[string]interface{}{"price": map[string]interface{}{"$gt": int64(5)}},
	},
	{
		SearchCriteria{
			Prefix:    false,
			Key:       "price",
			Operation: GREATER_THAN_EQUAL,
			Value:     "5",
		},
		map[string]interface{}{"price": map[string]interface{}{"$gte": int64(5)}},
	},
	{
		SearchCriteria{
			Prefix:    false,
			Key:       "price",
			Operation: LESS_THAN,
			Value:     "5",
		},
		map[string]interface{}{"price": map[string]interface{}{"$lt": int64(5)}},
	},
	{
		SearchCriteria{
			Prefix:    false,
			Key:       "price",
			Operation: LESS_THAN_EQUAL,
			Value:     "5",
		},
		map[string]interface{}{"price": map[string]interface{}{"$lte": int64(5)}},
	},
	{
		SearchCriteria{
			Prefix:    true,
			Key:       "email",
			Operation: EXISTS,
			Value:     "",
		},
		map[string]interface{}{"email": map[string]interface{}{"$exists": false}},
	},
}

func TestShouldConvertFromSearchCriteria(t *testing.T) {
	for _, test := range converetTests {
		result := Convert(test.Criteria, queryHandler)
		assert.Equal(t, test.Expected, result)
	}
}
