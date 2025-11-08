package main

import (
	"fmt"
)

func append(arr []int, val int) []int {
	for i := 0; i < len(arr); i++ {
		arr[i] += val
	}

	return arr
}

func main() {
	fmt.Println("测试数组+1")

	arr := []int{1, 2, 3}
	fmt.Println("arr:", arr)
	arr = append(arr, 4)
	fmt.Println("arr:", arr)
}
