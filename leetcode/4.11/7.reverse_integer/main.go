package main

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	str := strconv.Itoa(x)
	arr := []byte(str)
	if str[0] == '-' {
		s, _ := strconv.Atoi(reverseStr(arr[1:]))
		if s > 1<<31 { // 1向左移31位，是2147483648,表示32位的最大整数
			return 0
		}
		return -s

	} else {
		s, _ := strconv.Atoi(reverseStr(arr))
		if s > 1<<31-1 { // 2147483647
			return 0
		}
		return s

	}

}

func reverseStr(s []byte) string {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return string(s)
}

func main() {
	s := 1223456
	fmt.Println(reverse(s))
	fmt.Println(1 << 31)
}
