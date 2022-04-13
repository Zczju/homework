package main

import (
	"fmt"
	"strings"
)

// 空字符串为有效的回文串
func isPalindrome(s string) bool {
	// 利用strings.Map函数提取字符中的字母和数字，并将大写换为小写
	s = strings.Map(func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return r + 32
		case r >= 'a' && r <= 'z':
			return r
		case r >= '0' && r <= '9':
			return r
		}
		return -1
	}, s)
	// 看字符串两边的元素是否相等
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true

}

func main() {
	fmt.Println(isPalindrome("0P"))
	fmt.Println(isPalindromeMtd2(",.,！@#"))
}

func isPalindromeMtd2(s string) bool {
	// 法二：双指针  注意边界问题
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		for left <= len(s)-1 && !isAlnum(s[left]) { // 条件也可以用 left < right
			left++
		}
		for right >= 0 && !isAlnum(s[right]) { // 同上
			right--
		}
		// 如果字符串s中没有字母和数字，right将会减到0，left将会加到len(s)-1
		// 所以要判断一下
		if left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}

// 可以用这个函数来判断是否是字母和数字
func isAlnum(ch byte) bool {
	// 加括号！
	return (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func isPalindromeUseExtraRoom(s string) bool {
	// 准备另一个字符串，只将字母和数字填入其中
	var sgood string
	for i := 0; i < len(s); i++ {
		if isalnum(s[i]) {
			sgood += string(s[i])
		}
	}

	n := len(sgood)
	sgood = strings.ToLower(sgood)
	for i := 0; i < n/2; i++ {
		if sgood[i] != sgood[n-1-i] {
			return false
		}
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
