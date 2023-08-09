package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ContactsTestCase struct {
	queries [][]string
	result  []int32
}

func TestContacts(t *testing.T) {
	tests := []ContactsTestCase{
		{
			queries: [][]string{
				{"add", "ed"},
				{"add", "eddie"},
				{"add", "edward"},
				{"find", "ed"},
				{"add", "edwina"},
				{"find", "edw"},
				{"find", "a"},
			},
			result: []int32{3, 2, 0},
		},
		{
			queries: [][]string{
				{"add", "hack"},
				{"add", "hackerrank"},
				{"find", "hac"},
				{"find", "hak"},
			},
			result: []int32{2, 0},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, contacts(test.queries))
	}
}
