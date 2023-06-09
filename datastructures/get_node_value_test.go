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
			list.insertNodeIntoSinglyLinkedList(int32(i))
		}
	} else {
		for i := 0; i < length; i++ {
			list.insertNodeIntoSinglyLinkedList(int32(i))
		}
	}
	return &list
}

func TestGetNodeValueSmallList(t *testing.T) {
	assert := assert.New(t)
	smallList := createLinkedList(4, true)
	smallListDesc := createLinkedList(10, false)
	bigList := createLinkedList(10000, false)

	tests := []GetNodeTestCase{
		{
			node:             smallList.head,
			positionFromTail: 2,
			result:           2,
		},
		{
			node:             smallListDesc.head,
			positionFromTail: 2,
			result:           8,
		},
		{
			node:             bigList.head,
			positionFromTail: 4,
			result:           9996,
		},
	}

	for _, test := range tests {
		assert.Equal(test.result, getNode(test.node, test.positionFromTail))
	}
}
