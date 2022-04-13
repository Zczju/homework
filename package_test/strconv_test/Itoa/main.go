package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// func Itoa(i int) string
	integer := 1234
	str := strconv.Itoa(integer)
	fmt.Println(str)
	arr := []byte(str)
	fmt.Println(arr)
	fmt.Println(string(arr))
	fmt.Println(str[0])
	fmt.Println(str[1])
	fmt.Println('a')
	fmt.Println('1')

	integer2 := -1234
	str2 := strconv.Itoa(integer2)
	fmt.Println(str2)
	fmt.Println([]byte(str2))

	// constant -99999999999999999999999 overflows int
	//integer3 := -99999999999999999999999
	//str3 := strconv.Itoa(integer3)
	//fmt.Println(str3)
	//fmt.Println([]byte(str3))

	fmt.Println(reflect.TypeOf('a'))
	fmt.Println(reflect.TypeOf(arr[0]))

	sgood := ""
	sgood += string(str2[0])
	fmt.Println(-9223372036854775808)
	i := 9223372036854775807 // math.MaxInt 整数的最大值 再+1会自动转为 -9223372036854775808
	fmt.Println(i + 1)

}
