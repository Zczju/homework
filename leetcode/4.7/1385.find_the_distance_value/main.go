package main

import (
	"fmt"
	"time"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	i, counter := 0, 0
	for i < len(arr1) {
		if isDistance(arr1[i], arr2, d) {
			counter++
		}
		i++
	}

	return counter
}

func isDistance(nums1 int, arr2 []int, d int) bool {
	j := 0
	for j < len(arr2) {
		if nums1-arr2[j] > 0 {
			if nums1-arr2[j] > d {
				j++
			} else {
				return false
			}
		} else {
			if arr2[j]-nums1 > d {
				j++
			} else {
				return false
			}
		}

	}
	return true
}

func main() {
	arr1 := []int{2, 1, 100, 3}
	arr2 := []int{-5, -2, 10, -3, 7}
	start := time.Now()
	t := findTheDistanceValue(arr1, arr2, 6)
	finish := time.Now()
	fmt.Println(t)
	fmt.Println(finish.Sub(start))
}
