package datastructures

func ReverseDoublyLinkedList(head *DoublyLinkedListNode) *DoublyLinkedListNode {
	if head == nil {
		return head
	}

	pointer := head
	for {
		// last node of
		if pointer.next == nil {
			pointer.next = pointer.prev
			pointer.prev = nil
			return pointer
		}

		tmp := pointer.next
		pointer.next = pointer.prev
		pointer.prev = tmp
		pointer = tmp
	}
}
