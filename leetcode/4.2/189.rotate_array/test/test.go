package main

import "fmt"

func rotateTest(nums []int, k int) {
	n := len(nums)
	k %= n
	reverseTest(nums)
	reverseTest(nums[:k])
	reverseTest(nums[k:])

}

func reverseTest(nums []int) []int {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
	return nums
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	// rotate(nums, 3)
	// rotateMtd2(nums, 3)
	rotateTest(nums, 3)
	fmt.Println(nums)
}
