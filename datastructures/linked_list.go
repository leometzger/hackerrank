package datastructures

import (
	"bufio"
	"fmt"
)

type SinglyLinkedListNode struct {
	Data int32
	Next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	Head *SinglyLinkedListNode
	Tail *SinglyLinkedListNode
}

func (singlyLinkedList *SinglyLinkedList) InsertNodeIntoSinglyLinkedList(nodeData int32) {
	node := &SinglyLinkedListNode{
		Next: nil,
		Data: nodeData,
	}

	if singlyLinkedList.Head == nil {
		singlyLinkedList.Head = node
	} else {
		singlyLinkedList.Tail.Next = node
	}

	singlyLinkedList.Tail = node
}

func PrintSinglyLinkedList(node *SinglyLinkedListNode, sep string, writer *bufio.Writer) {
	for node != nil {
		fmt.Fprintf(writer, "%d", node.Data)

		node = node.Next

		if node != nil {
			fmt.Fprintf(writer, sep)
		}
	}
}
