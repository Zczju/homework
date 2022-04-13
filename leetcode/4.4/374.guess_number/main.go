package main

import (
	"fmt"
	"sort"
)

func guessNumber(n int) int {
	left, right, mid := 1, n, 0
	for left <= right {
		mid = left + (right-left)/2
		t := guess(mid)
		if t == 0 {
			return mid
		} else if t == 1 {
			right = mid - 1
		} else if t == -1 {
			left = mid + 1
		}
	}
	return 0
}

func guessNumberFromLeetCode(n int) int {
	return sort.Search(n, func(i int) bool {
		return guess(i) >= 0
	})
}

func guess(num int) int {
	pick := 6
	if num == pick {
		return 0
	} else if num > pick {
		return 1
	} else {
		return -1
	}
}

func main() {
	fmt.Println(guessNumberFromLeetCode(10))
}
