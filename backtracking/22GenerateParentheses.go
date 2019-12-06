package backtracking

import "fmt"

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/
func generateParenthesis(n int) []string {
	if n == 0 {
		return make([]string, 0)
	}
	strs := []string{""}
	return generate2(strs, n, n)
}

func generate2(strs []string, left int, right int) []string {
	if left == 0 && right == 0 {
		return strs
	}
	var leftResult, rightResult []string
	length := len(strs)
	if left == right {
		leftStrs := make([]string, length)
		copy(leftStrs, strs)
		for index, _ := range leftStrs {
			leftStrs[index] = leftStrs[index] + "("
		}
		//printStrArray(leftStrs, "leftStr")
		leftResult = generate2(leftStrs, left-1, right)
	} else {
		if left > 0 {
			leftStrs := make([]string, length)
			copy(leftStrs, strs)
			for index, _ := range leftStrs {
				leftStrs[index] = leftStrs[index] + "("
			}

			leftResult = generate2(leftStrs, left-1, right)
		}
		if right > 0 {
			rightStrs := make([]string, length)
			copy(rightStrs, strs)
			for index, _ := range rightStrs {
				rightStrs[index] = rightStrs[index] + ")"
			}
			//printStrArray(rightStrs, "rightStr")
			rightResult = generate2(rightStrs, left, right-1)
		}
	}
	lenLeft := len(leftResult)
	lenRight := len(rightResult)
	lenResult := lenLeft + lenRight
	result := make([]string, lenResult)
	//result[:lenLeft] = leftResult[:]
	copy(result[0:lenLeft], leftResult)
	copy(result[lenLeft:], rightResult)
	//fmt.Printf("left = %d, right = %d\n", left, right)
	//printStrArray(leftResult, "left")
	//printStrArray(rightResult, "right")
	return result
}

func printStrArray(strs []string, msg string) {
	fmt.Println(msg)
	for _, e := range strs {
		fmt.Println(e)
	}
}

func Test22() {
	for i := 1; i < 2; i++ {
		result := generateParenthesis(3)
		printStrArray(result, "res")
	}
}
