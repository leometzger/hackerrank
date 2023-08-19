package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/leometzger/hackerrank/datastructures"
)

func main() {
	var testCases int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	fmt.Fscanf(reader, "%d\n", &testCases)

	for i := 0; i < testCases; i++ {
		var llistSize int

		llistOne := &datastructures.SinglyLinkedList{}
		fmt.Fscanf(reader, "%d\n", &llistSize)
		for j := 0; j < llistSize; j++ {
			var value int32
			fmt.Fscanf(reader, "%d\n", &value)
			llistOne.InsertNodeIntoSinglyLinkedList(value)
		}

		llistTwo := &datastructures.SinglyLinkedList{}
		fmt.Fscanf(reader, "%d\n", &llistSize)
		for j := 0; j < llistSize; j++ {
			var value int32
			fmt.Fscanf(reader, "%d\n", &value)
			llistTwo.InsertNodeIntoSinglyLinkedList(value)
		}

		result := datastructures.MergeTwoLinkedLists(llistOne.Head, llistTwo.Head)
		for result != nil {
			fmt.Fprintf(writer, "%d ", result.Data)
			result = result.Next
		}
		fmt.Fprintf(writer, "\n")
	}

	writer.Flush()
}
