package backtracking

//Given a collection of distinct integers, return all possible permutations.
//
//Example:
//
//Input: [1,2,3]
//Output:
//[
//[1,2,3],
//[1,3,2],
//[2,1,3],
//[2,3,1],
//[3,1,2],
//[3,2,1]
//]

func permute(nums []int) [][]int {
	result := make([]int, 0)
	length := len(nums)
	isUsed := make([]int, length)
	return innerPermute(nums, isUsed, result)
}

func innerPermute(nums []int, isUsed []int, curArray []int) [][]int {
	//return innerPermute(nums, isUsed, result)
	result := make([][]int, 0)
	for idx, val := range nums {
		if isUsed[idx] == 0 {
			isUsed[idx] = 1
			tmpArray := make([]int, len(curArray))
			copy(tmpArray, curArray)
			tmpArray = append(tmpArray, val)
			partRes := innerPermute(nums, isUsed, tmpArray)
			result = myAppend(result, partRes)
			// clear
			isUsed[idx] = 0
		}
	}
	if len(result) == 0 {
		result = append(result, curArray)
	}
	return result
}

//func myAppend(result [][]int, part [][]int) [][]int {
//    for _, val := range part {
//        result = append(result, val)
//    }
//    return result
//}
//
//func myPrint(result [][]int) {
//    for _, arr := range result {
//        for _, val := range arr {
//            fmt.Print(val)
//        }
//        fmt.Print("\n")
//    }
//}

func Test46() {
	//in1 := []int{1}
	//res := permute(in1)
	//myPrint(res)
	testCase := [][]int{{1}, {1, 2}, {1, 2, 3}}
	for _, v := range testCase {
		res := permute(v)
		myPrint(res)
	}
}
