package mgs

import (
	"regexp"
)

const OPERATOR_PATTERN string = "(?P<Prefix>!?)(?P<Key>[^><!=]+)(?P<Operator>[><]=?|!?=|)(?P<Value>.*)"

const REGEX_PATTERN string = "^\\/(?P<Pattern>.*)\\/(?P<Option>[igm]*)$"

var operator *regexp.Regexp = nil
var pattern *regexp.Regexp = nil

func init() {
	operator = regexp.MustCompile(OPERATOR_PATTERN)
	pattern = regexp.MustCompile(REGEX_PATTERN)
}

// GetOperatorPattern gets operator RegExp pattern
func GetOperatorPattern() *regexp.Regexp {
	return operator
}

// GetRegexPattern gets RegExp pattern
func GetRegexPattern() *regexp.Regexp {
	return pattern
}
