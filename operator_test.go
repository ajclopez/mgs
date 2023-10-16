package mgs

import "testing"

type operatorTest struct {
	Input    string
	Expected SearchOperator
}

type operatorStringTest struct {
	Input    SearchOperator
	Expected string
}

var operatorTests = []operatorTest{
	{"=", EQUAL},
	{"!=", NOT_EQUAL},
	{">", GREATER_THAN},
	{">=", GREATER_THAN_EQUAL},
	{"<", LESS_THAN},
	{"<=", LESS_THAN_EQUAL},
	{"!", EXISTS},
}

var operatorStringTests = []operatorStringTest{
	{EQUAL, "EQUAL"},
	{NOT_EQUAL, "NOT_EQUAL"},
	{GREATER_THAN, "GREATER_THAN"},
	{GREATER_THAN_EQUAL, "GREATER_THAN_EQUAL"},
	{LESS_THAN, "LESS_THAN"},
	{LESS_THAN_EQUAL, "LESS_THAN_EQUAL"},
	{EXISTS, "EXISTS"},
}

func TestShouldGetOperationFromString(t *testing.T) {
	for _, test := range operatorTests {
		if result := test.Expected.GetOperation(test.Input); result != test.Expected {
			t.Errorf("Result %s not equal to expected %s", result, test.Expected)
		}
	}
}

func TestShouldReturnStringFromSearchOperator(t *testing.T) {
	for _, test := range operatorStringTests {
		if result := test.Input.String(); result != test.Expected {
			t.Errorf("Result %s not equal to expected %s", result, test.Expected)
		}
	}
}
