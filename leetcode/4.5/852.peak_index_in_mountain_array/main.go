package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1
	mid := 0
	for l <= r {
		mid = l + (r-l)/2
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		} else if arr[mid] < arr[mid-1] {
			r = mid - 1
		} else if arr[mid] < arr[mid+1] {
			l = mid + 1
		}
	}
	return -1
}

func main() {
	arr := []int{3, 5, 3, 2, 0}
	fmt.Println(peakIndexInMountainArray(arr))
}
