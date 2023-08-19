// https://www.hackerrank.com/challenges/insert-a-node-at-a-specific-position-in-a-linked-list/problem?isFullScreen=true

package datastructures

func insertNodeAtPosition(head *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
	var pointer *SinglyLinkedListNode = head
	count := 1
	node := &SinglyLinkedListNode{Data: data}

	if position == 0 {
		node.Next = head
		return node
	}

	for pointer.Next != nil {
		if count == int(position) {
			previousNext := pointer.Next
			pointer.Next = node
			node.Next = previousNext
			break
		}
		pointer = pointer.Next
		count++
	}

	if pointer.Next == nil {
		pointer.Next = node
	}

	return head
}
