package main

import (
	"fmt"
)

func find_target(nums []int, target int) (int, int) {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return i, j
			}
		}
	}
	return 0, 0
}

func main() {
	fmt.Println("数组中找出目标值的测试")

	nums := []int{1, 2, 4, 6, 9}

	target := 8

	i, j := find_target(nums, target)
	fmt.Printf("find_target: (%d, %d)\n", nums[i], nums[j])

}
