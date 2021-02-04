package main

import (
	"encoding/json"
	"fmt"
)

type treeNode struct {
	Val   int       `json:"val"`
	Left  *treeNode `json:"left"`
	Right *treeNode `json:"right"`
}



func hasPathSum(root *treeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

func pathSum(root *treeNode, sum int) [][]int {
	result := [][]int{}

	




	return result
}

func main() {
	root := []byte(`{
		"val": 5,
		"left": {
			"val": 4,
			"left": {
				"val":   11,
				"left":  {"val": 7, "left": null, "right": null},
				"right": {"val": 2, "left": null, "right": null}},
			"right": null},
		"right": {
			"val":   8,
			"left":  {"val": 13, "left": null, "right": null},
			"right": {"val": 4, "left": null, "right": {"val": 1, "left": null, "right": null}}}}`)
	var tree treeNode
	err := json.Unmarshal(root, &tree)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", hasPathSum(&tree, 22))
	fmt.Printf("%#v\n", pathSum(&tree, 22))

}
