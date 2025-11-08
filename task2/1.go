package main

import "fmt"

func add_test(val *int) {

	*val += 10
}

func mul_test(val []int) {
	for i := range val {
		val[i] *= 2
	}
}

func main() {

	fmt.Println("指针操作测试")

	var a int = 10
	b := []int{1, 2, 3, 4, 5}

	fmt.Println("整数操作，Before ", a)
	add_test(&a)
	fmt.Println("After ", a)

	fmt.Println("slice操作，Before ", b)
	mul_test(b)
	fmt.Println("After ", b)

}
