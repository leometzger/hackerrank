package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type LibraryFineTestCase struct {
	d1     int32
	m1     int32
	y1     int32
	d2     int32
	m2     int32
	y2     int32
	result int32
}

func TestLibraryFine(t *testing.T) {
	tests := []LibraryFineTestCase{
		{
			d1:     9,
			m1:     6,
			y1:     2015,
			d2:     6,
			m2:     6,
			y2:     2015,
			result: 45,
		},
		{
			d1:     9,
			m1:     6,
			y1:     2017,
			d2:     6,
			m2:     6,
			y2:     2015,
			result: 10000,
		},
		{
			d1:     9,
			m1:     8,
			y1:     2015,
			d2:     6,
			m2:     6,
			y2:     2015,
			result: 1000,
		},
		{
			d1:     12,
			m1:     1,
			y1:     2014,
			d2:     1,
			m2:     1,
			y2:     2015,
			result: 0,
		},
		{
			d1:     2,
			m1:     7,
			y1:     2014,
			d2:     1,
			m2:     1,
			y2:     2015,
			result: 0,
		},
		{
			d1:     28,
			m1:     2,
			y1:     2015,
			d2:     15,
			m2:     4,
			y2:     2015,
			result: 0,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, libraryFine(test.d1, test.m1, test.y1, test.d2, test.m2, test.y2))
	}
}
