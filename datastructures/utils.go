package datastructures

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func readLine(reader *bufio.Reader) string {
	str, err := reader.ReadString('\n')

	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(str, "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getResultFromFile(filePath string) []int32 {
	file, err := os.Open(filePath)
	checkError(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	var results []int32

	for {
		item := readLine(reader)
		if item == "" {
			break
		}

		intItem, err := strconv.ParseInt(item, 10, 64)
		checkError(err)
		results = append(results, int32(intItem))
	}
	return results
}
