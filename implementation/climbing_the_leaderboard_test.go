package implementation

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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
	ranked, player := readInput()
	result := readOutput()

	assert.Equal(t, result, climbingLeaderboard(ranked, player))
}

func readInput() ([]int32, []int32) {
	var ranked []int32
	var players []int32

	file, err := os.Open("./tests/climbing-leaderboard-input.txt")
	checkError(err)

	reader := bufio.NewReader(file)

	_, _, err = reader.ReadLine()
	checkError(err)

	line, _, err := reader.ReadLine()
	checkError(err)
	valuesAsString := strings.Split(string(line), " ")

	for _, value := range valuesAsString {
		rank, err := strconv.ParseInt(value, 10, 32)
		checkError(err)
		ranked = append(ranked, int32(rank))
	}

	_, _, err = reader.ReadLine()
	checkError(err)

	line, _, err = reader.ReadLine()
	checkError(err)
	valuesAsString = strings.Split(string(line), " ")

	for _, value := range valuesAsString {
		player, err := strconv.ParseInt(string(value), 10, 32)
		checkError(err)
		players = append(players, int32(player))
	}

	return ranked, players
}

func readOutput() []int32 {
	var result []int32

	file, err := os.Open("./tests/climbing-leaderboard-output.txt")
	checkError(err)

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		value, err := strconv.ParseInt(string(line), 10, 32)
		checkError(err)
		result = append(result, int32(value))
	}

	return result
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
