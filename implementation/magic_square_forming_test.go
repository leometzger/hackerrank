package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type FormingMagicSquareTestCase struct {
	matrix [][]int32
	result int32
}

func TestFormingMagicSquares(t *testing.T) {
	assert := assert.New(t)

	tests := []FormingMagicSquareTestCase{
		{
			result: 7,
			matrix: [][]int32{
				{5, 3, 4},
				{1, 5, 8},
				{6, 4, 2},
			},
		},
		{
			result: 1,
			matrix: [][]int32{
				{4, 9, 2},
				{3, 5, 7},
				{8, 1, 5},
			},
		},
		{
			result: 4,
			matrix: [][]int32{
				{4, 8, 2},
				{4, 5, 7},
				{6, 1, 6},
			},
		},
	}

	for _, test := range tests {
		assert.Equal(test.result, FormingMagicSquare(test.matrix))
	}
}
