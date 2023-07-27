package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createDoublyLinkedList(length int, desc bool) *DoublyLinkedList {
	dllist := &DoublyLinkedList{}

	for i := 0; i < length; i++ {
		dllist.insertNodeIntoDoublyLinkedList(int32(i))
	}

	return dllist
}

func TestReverseDoublyLinkedList(t *testing.T) {
	dllist := createDoublyLinkedList(5, false)

	head := ReverseDoublyLinkedList(dllist.head)

	assert.Equal(t, int32(4), head.data)
	assert.Equal(t, int32(3), head.next.data)
	assert.Equal(t, int32(2), head.next.next.data)
	assert.Equal(t, int32(1), head.next.next.next.data)
	assert.Equal(t, int32(0), head.next.next.next.next.data)
}

func TestReverseDoublyLinkedListWithNilHead(t *testing.T) {
	head := ReverseDoublyLinkedList(nil)
	assert.Nil(t, head)
}
