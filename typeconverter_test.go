package mgs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockTypeConverter struct{}

func (m *mockTypeConverter) Convert(typeName, value string) (interface{}, error) {
	switch typeName {
	case "ObjectID":
		// Mocking a valid ObjectID for demonstration purposes
		return value, nil
	default:
		return nil, errors.New("unsupported type")
	}
}

func TestTypeConverterObjectID(t *testing.T) {
	converter := &mockTypeConverter{}
	val := "656ec3473d9b0f45f5ed0321"

	result, err := converter.Convert("ObjectID", val)

	assert.NoError(t, err)
	assert.Equal(t, val, result)
}
