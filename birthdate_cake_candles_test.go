package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type BirthdateCakeCandlesTestCase struct {
	candles []int32
	result  int32
}

func TestBirthdateCakeCandles(t *testing.T) {
	assert := assert.New(t)

	tests := []BirthdateCakeCandlesTestCase{
		{
			candles: []int32{3, 2, 1, 3},
			result:  2,
		},
		{
			candles: []int32{6, 2, 1, 3},
			result:  1,
		},
		{
			candles: []int32{9, 9, 9, 9},
			result:  4,
		},
	}

	for _, test := range tests {
		assert.Equal(test.result, BirthdateCakeCandles(test.candles))
	}
}
