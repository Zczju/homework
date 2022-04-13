package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	// 这里有一个问题，就是map中的key是独一无二的
	// 所以如果nums中有重复元素，就需要考虑脚标的问题
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		if v, ok := m[target-nums[i]]; ok && i != v {
			return []int{i, v}
		}
	}

	return nil
}

func twoSumMtd2(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		// 先检验，再加入字典
		if v, ok := m[target-nums[i]]; ok {
			return []int{i, v}
		}
		m[nums[i]] = i
	}
	return nil
}

func twoSumMTD3(nums []int, target int) []int {
	// 两层for暴力解法
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumMtd4(nums []int, target int) []int {
	// 双指针
	i, j := 0, 1
	for nums[i]+nums[j] != target {
		if j == len(nums)-1 {
			i++
			j = i
		}
		j++
	}
	return []int{i, j}
}

func main() {
	arr := []int{5, 3, 3}
	fmt.Println(twoSum(arr, 6))
}
