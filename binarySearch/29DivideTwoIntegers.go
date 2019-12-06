package binarySearch

import "fmt"

/*
Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.

Input: dividend = 7, divisor = -3
Output: -2
*/
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	isNegative := false
	if dividend < 0 {
		dividend -= dividend + dividend
		isNegative = !isNegative
	}
	if divisor < 0 {
		divisor -= divisor + divisor
		isNegative = !isNegative
	}
	// begin binary search
	multor := 1
	curNum := divisor
	for {
		if curNum > dividend {
			break
		} // else
		curNum += curNum
		multor += multor
	}
	// too big, go back linear
	for {
		if curNum <= dividend {
			break
		}
		curNum -= divisor
		multor -= 1
	}
	res := multor
	if isNegative {
		res -= res + res
	}
	return res
}

func Test29() {
	dividends := []int{7, 0, 8, -7, -8}
	divisors := []int{-3, -2, -4, 3, -9}
	//ans := []int{-2, 0, -2, -2, 0}
	length := len(dividends)
	for i := 0; i < length; i++ {
		res := divide(dividends[i], divisors[i])
		fmt.Println(res)
	}
}
