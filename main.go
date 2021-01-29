package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	currNode := l.head
	for i := l.length; i > 0; i-- {
		fmt.Printf("%d \n", currNode.data)
		currNode = currNode.next
	}
}

func (l *linkedList) deleteWithValue(value int) {
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
			fmt.Println("no match")
			return
		}
		prevToDelete = prevToDelete.next
	}
	prevToDelete.next = prevToDelete.next.next
	l.length--
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 2}
	node2 := &node{data: 4}
	node3 := &node{data: 8}
	node4 := &node{data: 16}
	node5 := &node{data: 32}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.printListData()
	myList.deleteWithValue(100)
	myList.deleteWithValue(8)
	myList.printListData()
}
