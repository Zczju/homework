package main

import "fmt"

func searchRange(nums []int, target int) []int {
	return []int{searchFirst(nums, target), searchEnd(nums, target)}
}

func searchFirst(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}

	}
	return l
}

func searchEnd(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l+1)>>1
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid
		}

	}
	return l
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(nums, 8))
	fmt.Println([]byte("a"))
	fmt.Println([]byte("b"))
	fmt.Println([]byte("c"))
	fmt.Println([]byte("d"))
	fmt.Println([]byte("e"))
	fmt.Println([]byte("z"))
	str := "leetcode"
	fmt.Println([]byte(str))
}
