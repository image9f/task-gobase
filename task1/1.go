
package main

import(
	"fmt"
)

func findSingleNumber(nums []int) int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
		// fmt.Println(num, counts[num])
	}

	for num, count := range counts {
		if count == 1 {
			return num
		}
	}

	return -1
}

func main() {
	nums := []int{2, 2, 4, 5, 6, 4, 5}
	fmt.Println("The single number is:", findSingleNumber(nums))
}