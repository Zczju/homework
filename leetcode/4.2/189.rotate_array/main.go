package main

import (
	"fmt"
)

// 这两个方法都需要传址

func rotate(nums []int, k int) {
	// 数组未改变，值传递
	n := len(nums) - 1
	arr := nums[(n - k + 1):]
	nums = nums[:(n - k + 1)]
	nums = append(arr, nums...)
}

func rotateMtd2(nums []int, k int) {
	for i := 0; i < k; i++ {
		num := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		nums = append([]int{num}, nums...)
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	// rotate(nums, 3)
	// rotateMtd2(nums, 3)
	rotateFromLeetCode(nums, 3)
	fmt.Println(nums)
}

func rotateFromLeetCode(nums []int, k int) {
	n := len(nums)
	// if k == n {
	//     return
	// }
	// if k > n {
	// }
	// k可能大于数组长度，取余
	k %= n
	// 小的数对大的数取余，结果还是小的数
	// 若k比n大，则需要忽略无效的变换 (因为移动n个位置，数组不变)
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
	// 先反转整个数组，然后分别反转前k个和后k个

}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}

}
