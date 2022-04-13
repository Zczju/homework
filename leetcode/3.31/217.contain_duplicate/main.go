package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := map[int]int{}
	for i, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = i
	}
	return false
}

func containDuplicateFromLeetCode(nums []int) bool {
	m := map[int]struct{}{} // 使用空结构体，节省空间
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = struct{}{}
	}
	return false
}
func main() {
	arr := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(arr))
}
