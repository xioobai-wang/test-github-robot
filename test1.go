package main

import "fmt"

func main() {
	test1()
}

func test1() {
	nums := []int{1, 2, 3, 4, 5}

	length := len(nums)
	for i := 0; i < length/2; i++ {
		nums[i], nums[length-i-1] = nums[length-i-1], nums[i]
	}
	fmt.Println(nums)
}
