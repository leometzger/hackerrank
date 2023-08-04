package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/leometzger/hackerrank/datastructures"
)

func main() {
	var n int32
	commands := []string{}
	params := [][]int32{}

	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()
	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	fistLine := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(fistLine[0], 10, 64)
	checkError(err)
	n = int32(nTemp)

	qTemp, err := strconv.ParseInt(fistLine[1], 10, 64)
	checkError(err)
	q := int32(qTemp)

	for i := 0; i < int(q); i++ {
		line := strings.Split(strings.TrimSpace(readLine(reader)), " ")
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

	fmt.Println("Res", n, len(commands), len(params))

	result := datastructures.MergingCommunities(n, commands, params)

	fmt.Println("Res", n, len(commands), len(params), result)

	for _, res := range result {
		fmt.Fprintf(writer, "%d\n", res)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
