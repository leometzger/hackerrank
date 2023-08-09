package datastructures

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

func TestContactsFiles(t *testing.T) {
	input := "./tests/input03-trie.txt"
	output := "./tests/output03-trie.txt"
	queries := readParamsFromFile(input)
	result := getResultFromFile(output)

	assert.Equal(t, result, contacts(queries))
}

func readParamsFromFile(filePath string) [][]string {
	file, err := os.Open(filePath)
	checkError(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	nOperations := strings.TrimSpace(readLine(reader))
	operations, err := strconv.ParseInt(nOperations, 10, 64)
	checkError(err)

	var queries [][]string
	for i := 0; i < int(operations); i++ {
		query := strings.Split(strings.TrimSpace(readLine(reader)), " ")
		if len(query) < 2 {
			continue
		}

		queries = append(queries, query)
	}
	return queries
}
