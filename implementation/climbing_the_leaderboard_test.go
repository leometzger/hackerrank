package implementation

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ClimbingTheLeaderboard struct {
	ranked []int32
	player []int32
	result []int32
}

func TestClimbingTheLeaderboard(t *testing.T) {
	tests := []ClimbingTheLeaderboard{
		{
			ranked: []int32{100, 100, 50, 40, 40, 20, 10},
			player: []int32{5, 25, 50, 120},
			result: []int32{6, 4, 2, 1},
		},
		{
			ranked: []int32{100, 90, 90, 80, 75, 60},
			player: []int32{50, 65, 77, 90, 102},
			result: []int32{6, 5, 4, 2, 1},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, climbingLeaderboard(test.ranked, test.player))
	}
}

func TestClimbingTheLeaderboardFiles(t *testing.T) {
	ranked, player := readInput("./tests/climbing-leaderboard-input.txt")
	result := readOutput("./tests/climbing-leaderboard-output.txt")

	assert.Equal(t, result, climbingLeaderboard(ranked, player))
}

func readInput(filePath string) ([]int32, []int32) {
	file, err := os.Open(filePath)
	checkError(err)

	reader := bufio.NewReader(file)

	var n int
	_, err = fmt.Fscanf(reader, "%d\n", &n)
	checkError(err)
	ranked := scanLine(reader, n)

	var m int
	_, err = fmt.Fscanf(reader, "%d\n", &m)
	checkError(err)
	players := scanLine(reader, m)

	return ranked, players
}

func scanLine(reader io.Reader, n int) []int32 {
	items := make([]int32, n)
	temp := make([]interface{}, n)
	for i := 0; i < n; i++ {
		temp[i] = &items[i]
	}
	_, err := fmt.Fscanln(reader, temp...)
	checkError(err)
	return items
}

func readOutput(filePath string) []int32 {
	var result []int32

	file, err := os.Open(filePath)
	checkError(err)
	reader := bufio.NewReader(file)

	for {
		var res int
		_, err = fmt.Fscanf(reader, "%d\n", &res)
		if err != nil {
			break
		}
		result = append(result, int32(res))
	}

	return result
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
