package datastructures

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MergingCommunitiesTestCase struct {
	n        int32
	commands []string
	params   [][]int32
	result   []int32
}

func TestMergingCommunities(t *testing.T) {
	tests := []MergingCommunitiesTestCase{
		{
			n: 3,
			commands: []string{
				"Q",
				"M",
				"Q",
				"M",
				"Q",
				"Q",
			},
			params: [][]int32{
				{1},
				{1, 2},
				{2},
				{2, 3},
				{3},
				{2},
			},
			result: []int32{1, 2, 3, 3},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, MergingCommunities(test.n, test.commands, test.params))

	}
}

func TestMergingCommunitiesFiles(t *testing.T) {
	n, commands, params := getFunctionArgsFromFile("./tests/mc04.txt")
	expectedResult := getResultFromFile("./tests/mc04-result.txt")
	result := MergingCommunities(n, commands, params)

	for i := 0; i < len(result); i++ {
		passed := assert.Equal(t, expectedResult[i], result[i], fmt.Sprintf("Error on %d", i))
		if !passed {
			break
		}
	}
}

func getFunctionArgsFromFile(filePath string) (int32, []string, [][]int32) {
	commands := []string{}
	params := [][]int32{}

	file, err := os.Open(filePath)
	checkError(err)
	defer file.Close()

	var n, q int32
	reader := bufio.NewReader(file)

	_, err = fmt.Fscanf(reader, "%d %d\n", &n, &q)
	checkError(err)

	for i := 0; i < int(q); i++ {
		line := strings.Split(strings.TrimSpace(readLine(reader)), " ")
		if line[0] == "" {
			break
		}

		command := line[0]
		commands = append(commands, command)

		param1Temp, err := strconv.ParseInt(line[1], 10, 64)
		checkError(err)
		param1 := int32(param1Temp)

		if command == "M" {
			param2Temp, err := strconv.ParseInt(line[2], 10, 64)
			checkError(err)
			param2 := int32(param2Temp)
			params = append(params, []int32{param1, param2})
		} else {
			params = append(params, []int32{param1})
		}
	}
	return n, commands, params
}
