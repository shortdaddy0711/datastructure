package main

import "fmt"

func indexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

// SysmetricDiff -
func SysmetricDiff(a []int, b []int) []int {
	var result []int
	for i := 0; i < len(a); i++ {
		if indexOf(a[i], b) < 0 && indexOf(a[i], result) < 0 {
			result = append(result, a[i])
		}
	}
	return result
}

func main() {
	d := [][]int{{1, 2, 3, 3}, {5, 2, 1, 4}}
	err := SysmetricDiff(d[0], d[1])
	if err != nil {
		fmt.Printf("%v \n", err)
	}
}
