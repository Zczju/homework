package main

import (
	"fmt"
	"strings"
)

func myAtoi(s string) int {
	if s == "" {
		return 0
	}
	i, n := 0, len(s)
	// 忽略前导空格
	// 这里也可以用Trim函数
	for i < n-1 {
		if s[i] == ' ' {
			i++
		} else {
			break
		}
	}
	// 检查非空格的第一位
	s0 := s[i:]
	if s[i] == '-' || s[i] == '+' {
		s = s[i+1:]
	} else if s[i] >= '0' && s[i] <= '9' {
		s = s[i:]
	} else {
		return 0
	}

	// 处理后的s为数字开头的字符串
	// 遍历s，并转化数字
	// 停止条件为：ch不是数字 或者 num超过最大长度
	num := 0
	for _, ch := range []byte(s) {
		ch -= '0'
		if ch > 9 || ch < 0 {
			break
		}
		// 这里的左移操作方法应该学习一下，从数字头部开始垒数字
		num = num*10 + int(ch)
		if num >= 1<<31 {
			num = 1 << 31
			break
		}
	}

	// 有负号的加符号
	if s0[0] == '-' {
		num = -num
	} else {
		// 另一边的边界
		if num >= 1<<31-1 {
			num = 1<<31 - 1
		}
	}
	return num
}

func main() {
	// 测试用例应包括 :
	// 正数: 正常，溢出
	// 负数: 正常，溢出
	// 其他字符
	s := "-91283472332"
	fmt.Println(myAtoi(s))
}

const (
	MaxInt = 1<<31 - 1
	MinInt = -(1 << 31)
)

func othersAtoi(str string) int {
	v := 0
	i := 0
	negtive := false // 记录正负号

	// trim blank prefix 去掉空格前缀
	for ; i < len(str); i++ {
		if ' ' != str[i] {
			break
		}
	}
	if i > len(str)-1 {
		return v
	}

	// first non-blank char 第一个非空字符
	if '+' == str[i] {
		i++
	} else if '-' == str[i] {
		negtive = true
		i++
	} else if str[i] < '0' || str[i] > '9' {
		return v
	}

	// integer
	for ; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			break
		}
		v = 10*v + int(str[i]-'0')
		if negtive {
			if -v <= MinInt {
				return MinInt
			}
			continue
		}
		if v < 0 || v > MaxInt {
			return MaxInt
		}
	}

	if negtive {
		v = -v
	}
	return v
}

func myAtoiUseTrim(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	negative := false
	abs := s
	switch s[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		abs = s
	case '-':
		negative = true
		abs = s[1:]
	case '+':
		abs = s[1:]
	default:
		return 0
	}

	absNum := 0
	for _, ch := range abs {
		ch -= '0'
		if ch > 9 || ch < 0 {
			break
		}
		absNum = absNum*10 + int(ch)
		if absNum >= 1<<31 {
			absNum = 1 << 31
			break
		}
	}
	if negative {
		absNum = -absNum
	} else {
		if absNum >= 1<<31-1 {
			absNum = 1<<31 - 1
		}
	}

	return absNum
}
