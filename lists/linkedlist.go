package lists

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil}
}

func (ll *LinkedList) Insert(data int) {
	newNode := Node{data: data, next: nil}
	if ll.head == nil {
		ll.head = &newNode
	} else {
		curr := ll.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = &newNode
	}

}

func (ll *LinkedList) Print() {
	curr := ll.head
	for curr != nil {
		fmt.Println(curr.data, "-->")
		curr = curr.next
	}
}
