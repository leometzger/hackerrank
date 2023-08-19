package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetNodeTestCase struct {
	node             *SinglyLinkedListNode
	positionFromTail int32
	result           int32
}

func createLinkedList(length int, desc bool) *SinglyLinkedList {
	list := SinglyLinkedList{}
	if desc {
		for i := length; i > 0; i-- {
			list.InsertNodeIntoSinglyLinkedList(int32(i))
		}
	} else {
		for i := 0; i < length; i++ {
			list.InsertNodeIntoSinglyLinkedList(int32(i))
		}
	}
	return &list
}

func TestGetNodeValueSmallList(t *testing.T) {
	assert := assert.New(t)
	singleList := createLinkedList(1, true)
	smallList := createLinkedList(4, true)
	smallListDesc := createLinkedList(10, false)
	bigList := createLinkedList(10000, false)

	tests := []GetNodeTestCase{
		{
			node:             smallList.Head,
			positionFromTail: 2,
			result:           3,
		},
		{
			node:             smallListDesc.Head,
			positionFromTail: 2,
			result:           7,
		},
		{
			node:             bigList.Head,
			positionFromTail: 4,
			result:           9995,
		},
		{
			node:             singleList.Head,
			positionFromTail: 0,
			result:           1,
		},
	}

	for _, test := range tests {
		assert.Equal(test.result, getNode(test.node, test.positionFromTail))
	}
}
