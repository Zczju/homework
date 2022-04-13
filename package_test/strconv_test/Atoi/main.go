package main

import (
	"fmt"
	"strconv"
)

func main() {
	// strconv.Atoi用于将字符串string转化为整数int
	// func Atoi(s string) (i int, err error)
	// 输入s string，输出一个int和一个可能发生的error
	s1 := "01234"
	if integer, err := strconv.Atoi(s1); err == nil {
		fmt.Println(integer)
	}

	// 当输入的字符串不可转为整数类型，输出0和 err: strconv.Atoi: parsing "hello_world": invalid syntax
	s2 := "hello_world"
	if notInteger, err := strconv.Atoi(s2); err != nil {
		fmt.Println(notInteger)
		fmt.Println(err)
	}

	// 即使字符串中有一部分为整数也不行 strconv.Atoi: parsing "123hello_world": invalid syntax
	s3 := "123hello_world"
	if notInteger, err := strconv.Atoi(s3); err != nil {
		fmt.Println(notInteger)
		fmt.Println(err)
	}

	// 即使是浮点数也不行 strconv.Atoi: parsing "3.1415926": invalid syntax
	s4 := "3.1415926"
	if notIntegerButFloat, err := strconv.Atoi(s4); err != nil {
		fmt.Println(notIntegerButFloat)
		fmt.Println(err)
	}

	// 当输入的整数太大时
	// output : 9223372036854775807 (64位的最大整数)
	// strconv.Atoi: parsing "999999999999999999999": value out of range
	s5 := "9223372036854775808"
	if bigInteger, err := strconv.Atoi(s5); err != nil {
		fmt.Println(bigInteger)
		fmt.Println(err)
	}
}
