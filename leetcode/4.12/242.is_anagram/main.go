package main

import (
	"fmt"
	"sort"
)

// Method1: 哈希表
func isAnagram(s string, t string) bool {
	// 一开始的判断可以节省逻辑，
	// 如果长度不相等，一定不是字母异位词
	if len(s) != len(t) {
		return false
	}
	abMap := [26]int{}
	for _, ch := range s {
		// 这里必须用ch-'a'的原因是：
		// abMap是数组，该数组index必须从0开始，到25结束，不能越界
		abMap[ch-'a']++
	}
	for _, ch := range t {
		abMap[ch-'a']--
	}
	// 这步还可以优化，少遍历一遍
	// 在减的过程中直接判断是否有小于0的项
	for _, item := range abMap {
		if item != 0 {
			return false
		}
	}
	return true
	// 时间复杂度 O(N)
	// 空间复杂度 O(S) S长度为26

}

func isAnagramOptimized(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := map[rune]int{}
	for _, ch := range s {
		// 此处是map，可以直接用ch来做key
		cnt[ch]++
	}
	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}
	return true
}

// Method2: 排序，之后比较字符串是否相等
func isAnagramMtd2(s string, t string) bool {
	s1 := []byte(s)
	s2 := []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
	// 时间复杂度 O(NlogN) 高于线性低于平方
	// 空间复杂度 O(logN)
}

func main() {
	s := "abcdeee"
	t := "eaebecd"
	fmt.Println(isAnagram(s, t))
	fmt.Println(isAnagramMtd2(s, t))
}
