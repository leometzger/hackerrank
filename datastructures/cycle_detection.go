// https://www.hackerrank.com/challenges/detect-whether-a-linked-list-contains-a-cycle/problem?isFullScreen=true
// this one I have translated to Java 8, as in Golang it
// has not input reading code stub.

package datastructures

func DetectCycle(llist *SinglyLinkedList) bool {
	if llist.Head == nil {
		return false
	}

	cache := make(map[*SinglyLinkedListNode]bool)
	pointer := llist.Head

	for pointer.Next != nil {
		if cache[pointer] {
			return true
		}
		cache[pointer] = true
		pointer = pointer.Next
	}

	return false
}
