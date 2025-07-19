package mgs

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type result struct {
	Query Query
	Err   error
}

var tests = []struct {
	Input    string
	Expected result
}{
	{
		"skip=10",
		result{
			Query{
				Filter:     map[string][]interface{}{},
				Sort:       map[string]int(nil),
				Limit:      0,
				Skip:       10,
				Projection: map[string]int(nil),
			},
			nil,
		},
	},
	{
		"skip=number",
		result{
			Query{
				Filter:     map[string][]interface{}(nil),
				Sort:       map[string]int(nil),
				Limit:      0,
				Skip:       0,
				Projection: map[string]int(nil),
			},
			ErrValueNoMatch,
		},
	},
	{
		"limit=25",
		result{
			Query{
				Filter:     map[string][]interface{}{},
				Sort:       map[string]int(nil),
				Limit:      25,
				Skip:       0,
				Projection: map[string]int(nil),
			},
			nil,
		},
	},
	{
		"limit=number",
		result{
			Query{
				Filter:     map[string][]interface{}(nil),
				Sort:       map[string]int(nil),
				Limit:      0,
				Skip:       0,
				Projection: map[string]int(nil),
			},
			ErrValueNoMatch,
		},
	},
	{
		"sort=+date,-age,name",
		result{
			Query{
				Filter: map[string][]interface{}{},
				Sort: map[string]int{
					"date": 1,
					"age":  -1,
					"name": 1,
				},
				Limit:      0,
				Skip:       0,
				Projection: map[string]int(nil),
			},
			nil,
		},
	},
	{
		"fields=firstname,lastname,email",
		result{
			Query{
				Filter: map[string][]interface{}{},
				Sort:   map[string]int(nil),
				Limit:  0,
				Skip:   0,
				Projection: map[string]int{
					"firstname": 1,
					"lastname":  1,
					"email":     1,
				},
			},
			nil,
		},
	},
	{
		"_id=666f3a3ecf615a0f4d455411",
		result{
			Query{
				Filter: map[string][]interface{}{
					"$and": {
						map[string]interface{}{
							"_id": "ObjectID(\"666f3a3ecf615a0f4d455411\")",
						},
					},
				},
				Sort:       map[string]int(nil),
				Limit:      0,
				Skip:       0,
				Projection: map[string]int(nil),
			},
			nil,
		},
	},
	{
		"city=Madrid&age>=18",
		result{
			Query{
				Filter: map[string][]interface{}{
					"$and": {
						map[string]interface{}{
							"city": "Madrid",
						},
						map[string]interface{}{
							"age": map[string]interface{}{
								"$gte": int64(18),
							},
						},
					},
				},
				Sort:       map[string]int(nil),
				Limit:      0,
				Skip:       0,
				Projection: map[string]int(nil),
			},
			nil,
		},
	},
	{
		"filter=gender = female",
		result{
			Query{
				Filter: map[string][]interface{}{
					"$and": {
						map[string]interface{}{
							"gender": "female",
						},
					},
				},
				Sort:       map[string]int(nil),
				Limit:      0,
				Skip:       0,
				Projection: map[string]int(nil),
			},
			nil,
		},
	},
	{
		"filter=(city = Madrid or city = Santiago) and gender = female",
		result{
			Query{
				Filter: map[string][]interface{}{
					"$and": {
						map[string][]interface{}{
							"$or": {
								map[string]interface{}{
									"city": "Madrid",
								},
								map[string]interface{}{
									"city": "Santiago",
								},
							},
						},
						map[string]interface{}{
							"gender": "female",
						},
					},
				},
				Sort:       map[string]int(nil),
				Limit:      0,
				Skip:       0,
				Projection: map[string]int(nil),
			},
			nil,
		},
	},
}

// MockPrimitives it is a mock of the Primitives for testing.
type MockTypeConverter struct{}

func (mc *MockTypeConverter) Convert(oidStr string) (interface{}, error) {
	if oidStr == "666f3a3ecf615a0f4d455411" {
		return fmt.Sprintf("ObjectID(\"%s\")", oidStr), nil
	}

	return nil, errors.New("Invalid ObjectId")
}

var queryHandler = NewQueryHandler(&MockTypeConverter{})

func TestReturnDefaultQueryForMongoGoSearchWhenQueryIsEmpty(t *testing.T) {
	result, err := queryHandler.MongoGoSearch("", &FindOptions{})

	if err != nil {
		t.Errorf("Error parse url to mongo query search")
	}

	assert.Equal(t, Query{}, result)
}

func TestReturnErrorForMongoGoSearchWhenQueryInvalidCharacters(t *testing.T) {
	result, err := queryHandler.MongoGoSearch("name=Jho%%n", &FindOptions{})

	if err == nil {
		t.Errorf("Error parse url to mongo query search")
	}

	assert.Equal(t, Query{}, result)
}

func TestReturnQuery(t *testing.T) {
	for _, test := range tests {
		query, err := queryHandler.MongoGoSearch(test.Input, &FindOptions{})

		assert.Equal(t, test.Expected.Query, query)
		assert.Equal(t, test.Expected.Err, err)
	}
}

func TestReturnQueryUseFindOptionsWithCaster(t *testing.T) {

	caster := map[string]CastType{
		"mobile": STRING,
	}

	opts := FindOption()
	opts.SetCaster(caster)

	expected := result{
		Query{
			Filter: map[string][]interface{}{
				"$and": {
					map[string]interface{}{
						"mobile": "+56900000000",
					},
				},
			},
			Sort:       map[string]int(nil),
			Limit:      0,
			Skip:       0,
			Projection: map[string]int(nil),
		},
		nil,
	}

	query, err := queryHandler.MongoGoSearch("mobile=+56900000000", opts)

	assert.Equal(t, expected.Query, query)
	assert.Equal(t, expected.Err, err)
}

func TestReturnQueryUseFindOptionsWithDefaultLimit(t *testing.T) {

	opts := FindOption()
	opts.SetDefaultLimit(10)

	expected := result{
		Query{
			Filter:     map[string][]interface{}{},
			Sort:       map[string]int(nil),
			Limit:      10,
			Skip:       0,
			Projection: map[string]int(nil),
		},
		nil,
	}

	query, err := queryHandler.MongoGoSearch("limit=a", opts)

	assert.Equal(t, expected.Query, query)
	assert.Equal(t, expected.Err, err)
}

func TestReturnQueryUseFindOptionsWithMaxLimit(t *testing.T) {

	opts := FindOption()
	opts.SetMaxLimit(500)

	expected := result{
		Query{
			Filter:     map[string][]interface{}{},
			Sort:       map[string]int(nil),
			Limit:      500,
			Skip:       0,
			Projection: map[string]int(nil),
		},
		nil,
	}

	query, err := queryHandler.MongoGoSearch("limit=2000", opts)

	assert.Equal(t, expected.Query, query)
	assert.Equal(t, expected.Err, err)
}
