package main

import "fmt"

func moveZeroes(nums []int) {
	n := len(nums)
	i := 0
	// 很特么怪的想法: 如果这个元素是0，将其删除；
	// 全部删完之后，原长度减现长度，得到删去了多少0，加在末尾即可
	for {
		if i == len(nums)-1 {
			break
		}
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}

	nowLen := len(nums)
	if nowLen != n {
		for j := 0; j <= n-nowLen; j++ {
			nums = append(nums, 0)
		}
	}

}

func main() {
	arr := []int{1, 0, 0, 3, 12, 89, 9}
	moveZeroes(arr)
	// moveZeroes2Pointers(arr)
	// moveZeroesFromLeetCode(arr)
	fmt.Println(arr)
}

func moveZeroesFromLeetCode(nums []int) {
	i, j := 0, 0
	for j < len(nums) {
		if nums[j] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
		j++
	}
}

func moveZeroes2Pointers(nums []int) {
	i, j := 0, 0
	for {
		if i == len(nums) || j == len(nums) {
			return
		}
		for nums[i] == 0 {
			i++
			// 第一个不是0的元素
			if i == len(nums) {
				return
			}
		}
		for nums[j] != 0 {
			j++
			// 第一个0
			if j == len(nums) {
				return
			}
		}
		if i > j {
			nums[j] = nums[i]
			nums[i] = 0
		} else {
			i++
		}

	}

}
