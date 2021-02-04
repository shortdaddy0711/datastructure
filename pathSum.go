package main

import (
	"encoding/json"
	"fmt"
)

type binaryTree struct {
	val   int
	left  *binaryTree
	right *binaryTree
}

// func targetMatch(root binaryTree, target int) bool {

// }

func main() {
	root := []byte(`{
		val: 5,
		left: {
			val: 4,
			left: {
				val:   11,
				left:  {val: 7, left: null, right: null},
				right: {val: 2, left: null, right: null},
			},
			right: null,
		},
		right: {
			val:   8,
			left:  {val: 13, left: null, right: null},
			right: {val: 4, left: null, right: {val: 1, left: null, right: null}},
		}}`)
	var tree binaryTree
	err := json.Unmarshal(root, &tree)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", tree)
}

// const pathSum = (root, target) => {
//   if (!root) {
//     return false;
//   }
//   let isTargetExist = false;
//   function checkLeaf(node, sum) {
//     if (!node) {
//       return;
//     } else if (node.right === null && node.left === null) {
//       if (target === sum + node.val) {
//         isTargetExist = true;
//       }
//       return;
//     } else {
//       checkLeaf(node.left, sum + node.val);
//       checkLeaf(node.right, sum + node.val);
//     }
//   }
//   checkLeaf(root, 0);
//   return isTargetExist;
// };
