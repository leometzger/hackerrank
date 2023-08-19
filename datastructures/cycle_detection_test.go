package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCycleDetectionWithEmptyLList(t *testing.T) {
	llist := createInsertSpecificLinkedList(0, false)

	assert.False(t, DetectCycle(llist))
}

func TestCycleDetectionWithCycle(t *testing.T) {
	llist := createInsertSpecificLinkedList(3, false)
	llist.Head.Next.Next = llist.Head

	assert.True(t, DetectCycle(llist))
}

func TestCycleDetectionWithoutCycle(t *testing.T) {
	llist := createInsertSpecificLinkedList(10, false)

	assert.False(t, DetectCycle(llist))
}
