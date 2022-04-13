package main

import (
	"fmt"
)

// rune可以互相运算！并用作数组的index (是int32)
func firstUniqCharMtd1(s string) int {
	// 方法一：遍历两遍s
	cnt := [26]int{}
	// 容量为26的数组，代表从a到z的26个字母
	for _, ch := range s {
		// 'a'为 97, ch-'a'表示ch是a到z的第几个字母
		// 遍历到哪个字母，就将其加一，用来记录出现次数
		// fmt.Println(reflect.TypeOf(ch)) --> int32
		cnt[ch-'a']++
	}
	// 再遍历一遍，获取第一个出现次数为1的字母的下标
	for i, ch := range s {
		if cnt[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	s := "leetcode"
	fmt.Println(firstUniqCharTest(s))
}

func firstUniqCharMtd2(s string) int {
	n := len(s)
	cnt := [26]int{}
	for i := range cnt {
		cnt[i] = n
	}
	// 将cnt中元素值全置为n

	for i, ch := range s {
		ch -= 'a'
		if cnt[ch] == n {
			cnt[ch] = i
		} else {
			cnt[ch] = n + 1
		}
	}
	// 遍历一遍字符串，将只出现一次的元素值置为其下标i
	// 同时将出现多次的元素值置为n+1
	// 注意此时：没出现过的元素值为n

	// "cccsba"
	ans := n
	for _, p := range cnt {
		// 首先注意: cnt中元素的顺序是从a到z
		// 比如上边的字符串:s,b,a都只出现了一次，他们对应值(即下标)分别为3,4,5；而c出现了多次，c对应值为6+1+1+1
		// 每找到一个下标更小的值，都将其赋给ans
		// 这样遍历结束后，ans是只出现一次的元素中，下标最小的那一个
		if p < ans {
			ans = p
		}
	}
	if ans < n {
		return ans
	}
	// 遍历一遍cnt，找出只出现一次的元素的最小下标
	// 如果s中全是重复元素，则ans在遍历中不会被赋值，为n
	// 此时应返回-1
	return -1
}

func firstUniqCharTest(s string) int {
	// 举一反三做法，与方法二理论相同
	cnt := [26]int{}
	for i := range cnt {
		cnt[i] = -1
	}
	// cnt中元素值全为-1

	for i, ch := range s {
		ch -= 'a'
		if cnt[ch] == -1 {
			cnt[ch] = i
		} else {
			cnt[ch] = -2
		}
	}
	// 遍历一遍字符串，将只出现一次的元素值置为其下标i (这里出了问题，下标i一定不能为默认值，所以默认值不能设为0)
	// 故修改默认值为-1
	// 同时将出现多次的元素值置为-2
	// 注意此时：没出现过的元素值为-1

	ans := len(s)
	for _, p := range cnt {
		// 找到大于等于0的值里面最小的一个
		if p >= 0 {
			if p < ans {
				ans = p
			}
		}
	}
	if ans != len(s) {
		return ans
	}
	// 遍历一遍cnt，找出只出现一次的元素的最小下标,即ans
	// 如果s中全是重复元素，则ans在遍历中不会被赋值，为n
	// 此时应返回-1
	return -1
}
