package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type HurdleRaceTestCase struct {
	k      int32
	height []int32
	result int32
}

func TestHurdleRace(t *testing.T) {
	assert := assert.New(t)
	tests := []HurdleRaceTestCase{
		{
			int32(10),
			[]int32{2, 3, 4, 5},
			int32(0),
		},
		{
			int32(2),
			[]int32{2, 3, 4, 5},
			int32(3),
		},
		{
			int32(12),
			[]int32{100, 28, 4, 10},
			int32(88),
		},
	}

	for _, test := range tests {
		assert.Equal(HurdleRace(test.k, test.height), test.result)
	}
}
