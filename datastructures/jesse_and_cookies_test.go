package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type JesseAndCookiesTestCase struct {
	k      int32
	A      []int32
	result int32
}

func TestJesseAndCookies(t *testing.T) {
	tests := []JesseAndCookiesTestCase{
		{
			k:      7,
			result: 2,
			A:      []int32{1, 2, 3, 9, 10, 12},
		},
		{
			k:      700,
			result: -1,
			A:      []int32{1, 2, 3, 9, 10, 12},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, cookies(test.k, test.A))
	}
}
