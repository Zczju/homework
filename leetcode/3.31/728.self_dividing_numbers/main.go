package main

import "fmt"

func selfDividingNumbers(left int, right int) (arr []int) {
	for i := left; i <= right; i++ {
		if isSelfDividingNumber(i) {
			arr = append(arr, i)
		}
	}

	return
}

func isSelfDividingNumber(num int) bool {
	for i := num; i > 0; i /= 10 { // 此处不能写i>=0，因为golang中1/10得到的结果是0 =>会导致所有的结果都为false
		// (golang中只有相同类型的变量才能进行运算，结果也是相同的类型)
		if d := i % 10; d == 0 || num%d != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
}
