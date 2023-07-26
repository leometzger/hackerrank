// https://www.hackerrank.com/challenges/insert-a-node-at-a-specific-position-in-a-linked-list/problem?isFullScreen=true

package datastructures

func insertNodeAtPosition(head *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
	var pointer *SinglyLinkedListNode = head
	count := 1
	node := &SinglyLinkedListNode{data: data}

	if position == 0 {
		node.next = head
		return node
	}

	for pointer.next != nil {
		if count == int(position) {
			previousNext := pointer.next
			pointer.next = node
			node.next = previousNext
			break
		}
		pointer = pointer.next
		count++
	}

	if pointer.next == nil {
		pointer.next = node
	}

	return head
}
