package main

import (
	"fmt"
	"sort"
	"time"
)

// 屎山代码解法 ： 二分
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	i, counter := 0, 0
	sort.Ints(arr2)
	for i < len(arr1) {
		if isDistance(arr1[i], arr2, d) {
			counter++
		}
		i++
	}
	return counter
}

func isDistance(nums1 int, arr2 []int, d int) bool {
	l, r := 0, len(arr2)-1
	for l <= r {
		mid := l + (r-l)>>1
		if arr2[mid] < nums1 {
			l = mid + 1
		} else if arr2[mid] > nums1 {
			r = mid - 1
		} else if arr2[mid] == nums1 {
			return false
		}
	}
	if l > len(arr2)-1 {
		if nums1-arr2[len(arr2)-1] > d {
			return true
		} else {
			return false
		}
	} else if r < 0 {
		if arr2[0]-nums1 > d {
			return true
		} else {
			return false
		}
	} else if nums1-arr2[r] > d && arr2[l]-nums1 > d {
		return true
	} else {
		return false
	}
}

func main() {
	arr1 := []int{-8, -7}
	arr2 := []int{4, 10, -4, 5, 2}
	start := time.Now()
	t := findTheDistanceValue(arr1, arr2, 55)
	finish := time.Now()
	fmt.Println(t)
	fmt.Println(finish.Sub(start))
}
