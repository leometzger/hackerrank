// https://www.hackerrank.com/challenges/insert-a-node-at-a-specific-position-in-a-linked-list/problem?isFullScreen=true

package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type InsertSpecificPositionTestCase struct {
	head     *SinglyLinkedListNode
	data     uint32
	position uint32
}

func createInsertSpecificLinkedList(length int, desc bool) *SinglyLinkedList {
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

func TestInsertHead(t *testing.T) {
	list := createInsertSpecificLinkedList(5, false)

	result := insertNodeAtPosition(list.head, 20, 0)

	assert.Equal(t, result.data, int32(20))
	assert.Equal(t, result.next.data, int32(0))
	assert.Equal(t, result.next.next.data, int32(1))
}

func TestInsertSpecificPosition(t *testing.T) {
	list := createInsertSpecificLinkedList(5, false)

	result := insertNodeAtPosition(list.head, 20, 2)

	assert.Equal(t, int32(0), result.data)
	assert.Equal(t, int32(1), result.next.data)
	assert.Equal(t, int32(20), result.next.next.data)
	assert.Equal(t, int32(2), result.next.next.next.data)
	assert.Equal(t, int32(3), result.next.next.next.next.data)
}

func TestInsertTail(t *testing.T) {
	list := createInsertSpecificLinkedList(3, false)

	result := insertNodeAtPosition(list.head, 32, 3)
	assert.Equal(t, int32(0), result.data)
	assert.Equal(t, int32(1), result.next.data)
	assert.Equal(t, int32(2), result.next.next.data)
	assert.Equal(t, int32(32), result.next.next.next.data)
}
