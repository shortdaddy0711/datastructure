package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) Prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) PrintListData() {
	currNode := l.head
	for i := l.length; i > 0; i-- {
		fmt.Printf("%d ", currNode.data)
		currNode = currNode.next
	}
	fmt.Printf("\n")
}

func (l *linkedList) DeleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	prevToDelete := l.head
	for prevToDelete.next.data != value {
		if prevToDelete.next.next == nil {
			fmt.Printf("no match with %v\n", value)
			return
		}
		prevToDelete = prevToDelete.next
	}
	prevToDelete.next = prevToDelete.next.next
	fmt.Printf("%v has removed\n", value)
	l.length--
}

func main()  {
	myList := linkedList{}
	node1 := &node{data: 2}
	node2 := &node{data: 4}
	node3 := &node{data: 8}
	node4 := &node{data: 16}
	node5 := &node{data: 32}
	myList.Prepend(node1)
	myList.Prepend(node2)
	myList.Prepend(node3)
	myList.Prepend(node4)
	myList.Prepend(node5)
	myList.PrintListData()
	myList.DeleteWithValue(100)
	myList.DeleteWithValue(8)
	myList.PrintListData()
}