package mgs

import "testing"

func TestShouldGetOperatorPattern(t *testing.T) {
	if GetOperatorPattern() == nil {
		t.Errorf("Operator regular expression instance failed")
	}
}

func TestShouldGetRegexpPattern(t *testing.T) {
	if GetRegexPattern() == nil {
		t.Errorf("Regular expression instance failed")
	}
}
