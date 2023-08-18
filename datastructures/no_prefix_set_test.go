package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type NoPrefixSetTestCase struct {
	words  []string
	result string
}

func TestNoPrefixSet(t *testing.T) {
	tests := []NoPrefixSetTestCase{
		{
			words:  []string{"ab", "bc", "cd"},
			result: "",
		},
		{
			words:  []string{"abcd", "bcd", "abcde", "bcde"},
			result: "abcde",
		},
		{
			words:  []string{"aab", "defgab", "abcde", "aabcde", "bbbbbbbbbb", "jabjjjad"},
			result: "aabcde",
		},
		{
			words:  []string{"aab", "aac", "aacghgh", "aabghgh"},
			result: "aacghgh",
		},
		{
			words:  []string{"aab", "aac", "a", "aacghgh", "aabghgh"},
			result: "a",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, noPrefix(test.words))
	}
}
