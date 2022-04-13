package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left, right, mid := 0, len(nums)-1, 0
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if target > nums[mid] {
		return mid + 1
	}
	if target < nums[mid] {
		if mid == 0 {
			return 0
		}
		return mid - 1
	}

	return -1
}

func main() {
	nums := []int{1, 3}
	index := searchInsert(nums, 2)
	fmt.Println(index)
}
