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
			list.InsertNodeIntoSinglyLinkedList(int32(i))
		}
	} else {
		for i := 0; i < length; i++ {
			list.InsertNodeIntoSinglyLinkedList(int32(i))
		}
	}
	return &list
}

func TestInsertHead(t *testing.T) {
	list := createInsertSpecificLinkedList(5, false)

	result := insertNodeAtPosition(list.Head, 20, 0)

	assert.Equal(t, result.Data, int32(20))
	assert.Equal(t, result.Next.Data, int32(0))
	assert.Equal(t, result.Next.Next.Data, int32(1))
}

func TestInsertSpecificPosition(t *testing.T) {
	list := createInsertSpecificLinkedList(5, false)

	result := insertNodeAtPosition(list.Head, 20, 2)

	assert.Equal(t, int32(0), result.Data)
	assert.Equal(t, int32(1), result.Next.Data)
	assert.Equal(t, int32(20), result.Next.Next.Data)
	assert.Equal(t, int32(2), result.Next.Next.Next.Data)
	assert.Equal(t, int32(3), result.Next.Next.Next.Next.Data)
}

func TestInsertTail(t *testing.T) {
	list := createInsertSpecificLinkedList(3, false)

	result := insertNodeAtPosition(list.Head, 32, 3)
	assert.Equal(t, int32(0), result.Data)
	assert.Equal(t, int32(1), result.Next.Data)
	assert.Equal(t, int32(2), result.Next.Next.Data)
	assert.Equal(t, int32(32), result.Next.Next.Next.Data)
}
