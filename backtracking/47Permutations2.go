package backtracking

import "fmt"

//Given a collection of numbers that might contain duplicates, return all possible unique permutations.
//
//Example:
//
//Input: [1,1,2]
//Output:
//[
//[1,1,2],
//[1,2,1],
//[2,1,1]
//]

func permuteUnique(nums []int) [][]int {
	length := len(nums)
	isUsed := make([]int, length)
	result := make([]int, 0)
	return innerPermuteUnique(nums, isUsed, result)
}

func innerPermuteUnique(nums []int, isUsed []int, curArray []int) [][]int {
	haveUsed := make(map[int]bool)
	result := make([][]int, 0)
	for idx, val := range nums {
		if isUsed[idx] == 0 {
			if _, ok := haveUsed[val]; !ok {
				//fmt.Print(ok)
				haveUsed[val] = true
				isUsed[idx] = 1
				tmpArray := make([]int, len(curArray))
				copy(tmpArray, curArray)
				tmpArray = append(tmpArray, val)
				partRes := innerPermuteUnique(nums, isUsed, tmpArray)
				result = myAppend(result, partRes)
				// clear
				isUsed[idx] = 0
			}
		}
	}
	if len(result) == 0 {
		result = append(result, curArray)
	}
	return result
}

func myAppend(result [][]int, part [][]int) [][]int {
	for _, val := range part {
		result = append(result, val)
	}
	return result
}

func myPrint(result [][]int) {
	for _, arr := range result {
		for _, val := range arr {
			fmt.Print(val)
		}
		fmt.Print("\n")
	}
}

func Test47() {
	fmt.Println("Test 47")
	//in1 := []int{1}
	//res := permute(in1)
	//myPrint(res)
	//testCase := [][]int{{1}, {1,1},{1,1,2}}
	testCase := [][]int{{1, 1}, {1, 1, 2}}
	//myPrint(testCase)
	for _, v := range testCase {
		res := permuteUnique(v)
		myPrint(res)
	}
}
