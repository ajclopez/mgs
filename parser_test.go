package mgs

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var booleanTests = []struct {
	Input    string
	Caster   *CastType
	Expected bool
}{
	{"true", nil, true},
	{"True", nil, true},
	{"TRUE", nil, true},
	{"TrUe", nil, true},
	{"false", nil, false},
	{"False", nil, false},
	{"FALSE", nil, false},
	{"FaLsE", nil, false},
}

func TestShouldReturnCriteriaForParseWithQuery(t *testing.T) {
	opts := &FindOptions{}
	var value CastType
	criteria := Parse("name=John", opts.Caster)
	expected := []SearchCriteria{
		{
			Prefix:    false,
			Key:       "name",
			Operation: EQUAL,
			Value:     "John",
			Caster:    &value,
		},
	}

	if len(criteria) == 0 {
		t.Errorf("Error creating criteria array")
	}

	assert.Equal(t, expected, criteria)
}

func TestShouldBoolTypeForParseValue(t *testing.T) {
	for _, test := range booleanTests {
		assert.Equal(t, test.Expected, ParseValue(test.Input, nil))
	}
}

func TestShouldIntegerTypeForParseValue(t *testing.T) {

	assert.Equal(t, int64(10), ParseValue("10", nil))
}

func TestShouldFloatTypeForParseValue(t *testing.T) {

	assert.Equal(t, float64(3.1415), ParseValue("3.1415", nil))
}

func TestShouldDateTimeTypeForParseValue(t *testing.T) {
	expected := time.Date(2023, 10, 9, 20, 00, 00, 999000000, time.UTC)

	assert.Equal(t, expected, ParseValue("2023-10-09T20:00:00.999Z", nil))
}

func TestShouldStringListTypeForParseValue(t *testing.T) {
	expected := []interface{}{"apple", "banana", "grape"}

	assert.Equal(t, expected, ParseValue("apple,banana,grape", nil))
}

func TestShouldIntegerListTypeForParseValue(t *testing.T) {
	expected := []interface{}{int64(1), int64(2), int64(3)}

	assert.Equal(t, expected, ParseValue("1,2,3", nil))
}

func TestShouldRegexTypeForParseValue(t *testing.T) {

	assert.Equal(t, "John", ParseValue("John", nil))
}

func TestShouldStringTypeForParseValue(t *testing.T) {
	expected := Regex{
		Pattern: "@gmail\\.com$",
	}

	assert.Equal(t, expected, ParseValue("/@gmail\\.com$/", nil))
}

func TestShouldReturnUnsignedIntForParseIntValueToUnsignedIntWhenIsValidNumberString(t *testing.T) {
	result, err := parseIntValueToInt("1")

	if err != nil {
		t.Errorf("Error format string to uint")
	}

	assert.Equal(t, int64(1), result)
}

func TestShouldReturnErrorForParseIntValueToUnsignedIntWhenIsInvalidNumberString(t *testing.T) {
	result, err := parseIntValueToInt("a")

	if err == nil {
		t.Errorf("Error format string to uint")
	}

	assert.Equal(t, int64(0), result)
}

func TestShouldIntegerTypeForParseValueWithCastBoolean(t *testing.T) {

	cast := BOOLEAN

	assert.Equal(t, true, ParseValue("true", &cast))
}

func TestShouldIntegerTypeForParseValueWithCastDate(t *testing.T) {

	cast := DATE
	expected := time.Date(2023, 10, 9, 20, 00, 00, 999000000, time.UTC)

	assert.Equal(t, expected, ParseValue("2023-10-09T20:00:00.999Z", &cast))
}

func TestShouldIntegerTypeForParseValueWithIntCastNumber(t *testing.T) {

	cast := NUMBER
	assert.Equal(t, int64(10), ParseValue("10", &cast))
}

func TestShouldIntegerTypeForParseValueWithFloatCastNumber(t *testing.T) {

	cast := NUMBER
	assert.Equal(t, float64(10.5), ParseValue("10.5", &cast))
}

func TestShouldIntegerTypeForParseValueWithCastRegex(t *testing.T) {

	cast := PATTERN
	expected := Regex{
		Pattern: "@gmail\\.com$",
	}

	assert.Equal(t, expected, ParseValue("/@gmail\\.com$/", &cast))
}
