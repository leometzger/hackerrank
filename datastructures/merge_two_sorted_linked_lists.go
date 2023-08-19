package datastructures

func MergeTwoLinkedLists(headOne *SinglyLinkedListNode, headTwo *SinglyLinkedListNode) *SinglyLinkedListNode {
	llist := &SinglyLinkedList{}

	for headOne != nil || headTwo != nil {
		if headOne == nil {
			llist.InsertNodeIntoSinglyLinkedList(headTwo.Data)
			headTwo = headTwo.Next
		} else if headTwo == nil {
			llist.InsertNodeIntoSinglyLinkedList(headOne.Data)
			headOne = headOne.Next
		} else if headOne.Data <= headTwo.Data {
			llist.InsertNodeIntoSinglyLinkedList(headOne.Data)
			headOne = headOne.Next
		} else {
			llist.InsertNodeIntoSinglyLinkedList(headTwo.Data)
			headTwo = headTwo.Next
		}
	}

	return llist.Head
}
