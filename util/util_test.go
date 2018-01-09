package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParseInt(t *testing.T) {
	tests := []struct {
		Value    string
		Expected int
	}{
		{"0", 0},
		{"1", 1},
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"5", 5},
		{"6", 6},
		{"7", 7},
		{"8", 8},
		{"9", 9},
		{"10", 10},
		{"11", 11},
		{"12", 12},
		{"13", 13},
		{"14", 14},
		{"15", 15},
		{"16", 16},
		{"17", 17},
		{"18", 18},
		{"19", 19},
		{"20", 20},
		{"21", 21},
		{"22", 22},
		{"23", 23},
		{"24", 24},
		{"25", 25},
		{"26", 26},
		{"27", 27},
		{"28", 28},
		{"29", 29},
		{"30", 30},
		{"31", 31},
	}

	for _, test := range tests {
		valueParsed, err := ParseInt(test.Value)

		assert.NoError(t, err)
		assert.Equal(t, valueParsed, test.Expected)
	}

	valueError, err := ParseInt("^F4")

	assert.Error(t, err)
	assert.Equal(t, valueError, 0)
}
