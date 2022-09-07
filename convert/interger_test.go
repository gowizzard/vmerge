package convert_test

import (
	"github.com/gowizzard/vmerge/convert"
	"reflect"
	"testing"
)

// TestInteger is to test the convert integer function
// We convert a number it to an integer and check the result
func TestInteger(t *testing.T) {

	tests := []struct {
		number   string
		expected int
	}{
		{
			number:   "56",
			expected: 56,
		},
		{
			number:   "25",
			expected: 25,
		},
		{
			number:   "43",
			expected: 43,
		},
		{
			number:   "165",
			expected: 165,
		},
		{
			number:   "465",
			expected: 465,
		},
	}

	for _, value := range tests {

		integer := convert.Integer(value.number)

		if !reflect.DeepEqual(value.expected, integer) {
			t.Fatalf("expected: %d, got %d", value.expected, integer)
		}

	}

}
