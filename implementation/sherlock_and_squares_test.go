package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type SherlockAndSquaresTestCase struct {
	a      int32
	b      int32
	result int32
}

func TestSherlockAndSquares(t *testing.T) {
	tests := []SherlockAndSquaresTestCase{
		{
			a:      3,
			b:      9,
			result: 2,
		},
		{
			a:      17,
			b:      24,
			result: 0,
		},
		{
			a:      465868129,
			b:      988379794,
			result: 9855,
		},
		{

			a:      181510012,
			b:      293922871,
			result: 3672,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, squares(test.a, test.b))
	}
}
