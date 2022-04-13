package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "A man, a plan, a canal: Panama"
	s = strings.Map(func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return r + 32
		case r >= 'a' && r <= 'z':
			return r
		}
		return -1
		// If mapping returns a negative value, the character is dropped from the string with no replacement.
		// 如果mapping函数返回一个负值，这个字符被抛出string，并没有任何替代
	}, s)
	fmt.Println(s)
}
