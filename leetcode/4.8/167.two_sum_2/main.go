package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		if isTarget, anoIndex := searchTarget(target-numbers[i], numbers[i+1:]); isTarget {
			return []int{i + 1, anoIndex + (i + 1) + 1}
		}
	}
	return nil
}

func searchTarget(num int, nums []int) (isTarget bool, anotherIndex int) {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == num {
			return true, mid
		} else if nums[mid] > num {
			r = mid - 1
		} else if nums[mid] < num {
			l = mid + 1
		}
	}

	return false, -1
}

func main() {
	numbers := []int{-1, 0}
	fmt.Println(twoSum(numbers, -1))
}
