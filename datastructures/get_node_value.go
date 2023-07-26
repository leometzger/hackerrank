// https://www.hackerrank.com/challenges/get-the-value-of-the-node-at-a-specific-position-from-the-tail/problem?isFullScreen=true

package datastructures

func getNode(node *SinglyLinkedListNode, positionFromTail int32) int32 {
	length := 0
	n := node.next

	for n.next != nil {
		length++
		n = n.next
	}

	n = node.next
	for i := 0; i < length-int(positionFromTail)+1; i++ {
		n = n.next
	}

	return n.data
}
