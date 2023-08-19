// https://www.hackerrank.com/challenges/get-the-value-of-the-node-at-a-specific-position-from-the-tail/problem?isFullScreen=true

package datastructures

func getNode(head *SinglyLinkedListNode, positionFromTail int32) int32 {
	length := 0
	pointer := head

	for pointer != nil {
		length++
		pointer = pointer.Next
	}

	positionFromHead := length - (int(positionFromTail))
	pointer = head
	for i := 1; i < positionFromHead; i++ {
		pointer = pointer.Next
	}

	return pointer.Data
}
