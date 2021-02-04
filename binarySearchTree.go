package main

import "fmt"

// Node -
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

var depth int = 0

// Insert -
func (n *Node) Insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search -
func (n *Node) Search(k int) bool {
	depth++
	if n == nil {
		return false
	}
	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}
	return true
}

func main() {
	tree := &Node{Key: 45}
	tree.Insert(23)
	tree.Insert(59)
	tree.Insert(324)
	tree.Insert(2345)
	tree.Insert(4)
	tree.Insert(57)
	tree.Insert(2)
	tree.Insert(4678)
	tree.Insert(90)
	tree.Insert(236)
	fmt.Printf("%+v\n", tree)
	fmt.Printf("%v\n", tree.Search(57))
	fmt.Printf("%v\n", depth)
}
