package main

import "github.com/JKAravind/dsa-go/lists"

func main() {
	linkedlist := lists.NewLinkedList()
	linkedlist.Insert(1)
	linkedlist.Insert(2)
	linkedlist.Insert(3)
	linkedlist.Print()
}
