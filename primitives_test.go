package mgs

import (
	"errors"
	"testing"
)

// PrimitivesImpl converts string to ObjectID.
type PrimitivesImpl struct{}

func (p PrimitivesImpl) ObjectID(val string) (interface{}, error) {
	if val == "" {
		return nil, errors.New("empty string is not a valid ObjectID")
	}
	// Mocking a valid ObjectID for demonstration purposes
	return val + "_ObjectID", nil
}

func TestPrimitivesObjectID(t *testing.T) {
	p := PrimitivesImpl{}

	tests := []struct {
		name      string
		input     string
		want      interface{}
		expectErr bool
	}{
		{
			name:      "Valid input",
			input:     "12345",
			want:      "12345_ObjectID",
			expectErr: false,
		},
		{
			name:      "Empty input",
			input:     "",
			want:      nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := p.ObjectID(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("ObjectID() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if got != tt.want {
				t.Errorf("ObjectID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
