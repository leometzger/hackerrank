package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoSortedLinkedLists(t *testing.T) {
	listOne := &SinglyLinkedList{}
	listTwo := &SinglyLinkedList{}

	listOne.InsertNodeIntoSinglyLinkedList(1)
	listOne.InsertNodeIntoSinglyLinkedList(3)
	listOne.InsertNodeIntoSinglyLinkedList(5)
	listOne.InsertNodeIntoSinglyLinkedList(5)
	listTwo.InsertNodeIntoSinglyLinkedList(2)
	listTwo.InsertNodeIntoSinglyLinkedList(6)

	result := MergeTwoLinkedLists(listOne.Head, listTwo.Head)
	expectedResult := []int32{1, 2, 3, 5, 5, 6}

	for _, expected := range expectedResult {
		assert.Equal(t, expected, result.Data)
		if result.Next != nil {
			result = result.Next
		}
	}
}
