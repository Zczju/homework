package main

import "fmt"

func main() {
	arr := []int{2, 2, 4, 4, 5, 5, 7}
	theSingleNumber := singleNumber(arr)
	fmt.Println(theSingleNumber)

}

// 使用位运算的性质：异或√、同或
// 两个相同的数字 num^num = 0 ; 0^num = num ; 不论怎么运算，同一组数进行异或的结果相同
// 将数组中所有数字异或一遍，最后剩下的数字就是只出现了一次的数字
func singleNumber(nums []int) int {
	eos := 0 // 还可以省略这个变量，节省空间复杂度
	for _, num := range nums {
		eos ^= num
	}
	return eos
}
