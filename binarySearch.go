package main

import "fmt"

func binarySearch(nums []int, t int) int {
	min := 0
	max := len(nums) - 1
	for min <= max {
		mid := int((min + max) / 2)
		if nums[mid] == t {
			return mid
		} else if nums[mid] > t {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 6, 9, 12}
	t := 6
	index := binarySearch(nums, t)
	fmt.Printf("%v\n", index)
}
