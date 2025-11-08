package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

func main() {
	fmt.Println("测试数组去重")

	arr := []int{1, 2, 2, 3, 4, 4, 5}
	fmt.Println("arr:", arr)

	newLength := removeDuplicates(arr)
	fmt.Println("after removeDuplicates arr:", arr[:newLength])
}
