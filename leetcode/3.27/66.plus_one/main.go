package main

import "fmt"

// some conditions :
//      [1,2,3]
//      [1,2,3,9,9]
//      [9]
// 从后到前，找到第一个不为9的元素，将其加1，并将后续所有元素置0
// 如果所有元素均为9，全部置0，构造一个长度+1的数组，并将首位置为1
func plusOne(digits []int) []int {
	for i := 1; i <= len(digits); i++ {
		// 末位如果不是9， 加1并退出
		if digits[len(digits)-i] != 9 {
			digits[len(digits)-i]++
			return digits
		} else {
			// 末位如果是9，变为0
			// 如果是已经加到首位，首位变为0并进位
			// 如果未加到首位，退出本次循环，并在下次循环中加1退出
			digits[len(digits)-i] = 0
			if len(digits)-i == 0 {
				digits = append([]int{1}, digits...)
				return digits
			} else {
				continue
			}
		}
	}
	return digits
}

func plusOneFromLeetCode(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			// 倒序查找，找到第一个不为9的元素，将其加一
			for j := i + 1; j < n; j++ {
				digits[j] = 0
				// 将该元素后面的所有元素置为0
			}
			return digits
		}
	}

	// digits 中所有的元素均为 9
	digits = make([]int, n+1)
	// make一个长度+1的新数组，令其首位为1
	digits[0] = 1
	return digits
}

func main() {
	arr := []int{9, 9, 9, 9, 9, 9}
	arrChanged := plusOne(arr)
	// arrChanged := plusOneFromLeetCode(arr)
	fmt.Println(arrChanged)

}
