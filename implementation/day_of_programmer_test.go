package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type DayOfProgrammerTestCase struct {
	year   int32
	result string
}

func TestDayOfProgrammer(t *testing.T) {
	assert := assert.New(t)
	tests := []DayOfProgrammerTestCase{
		{int32(2017), "13.09.2017"},
		{int32(2016), "12.09.2016"},
		{int32(1800), "12.09.1800"},
		{int32(1918), "26.09.1918"},
		{int32(1917), "13.09.1917"},
		{int32(1772), "12.09.1772"},
	}

	for _, test := range tests {
		assert.Equal(test.result, DayOfProgrammer(test.year))
	}
}
