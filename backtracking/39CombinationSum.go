package backtracking

import (
	"sort"
)
import "leetcode/utils"

/*
Example 1:

Input: candidates = [2,3,6,7], target = 7,
A solution set is:
[
  [7],
  [2,2,3]
]

Example 2:

Input: candidates = [2,3,5], target = 8,
A solution set is:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]

*/
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	xsize := len(candidates)
	ysize := target + 1
	getEmptyMat := func(x int, y int) [][][]int {
		result := make([][][]int, xsize)
		for i, _ := range result {
			result[i] = make([][]int, y)
			for j, _ := range result[i] {
				result[i][j] = make([]int, 0)
			}
		}
		return result
	}
	//
	result := getEmptyMat(xsize, ysize)
	// init the first line
	idx := 0
	for j := 0; j < ysize; j++ {
		xval := candidates[idx]
		if j%xval == 0 {
			result[idx][j] = []int{j / xval}
		}
	}
	// iterate remains
	idx++
	for ; idx < xsize; idx++ {
		xval := candidates[idx]
		for j := 0; j < ysize; j++ {
			if len(result[idx-1][j]) != 0 {
				result[idx][j] = append(result[idx][j], 0)
			}
			counter := 1
			for {
				remains := j - counter*xval
				if remains < 0 {
					break
				}
				if len(result[idx-1][remains]) != 0 {
					result[idx][j] = append(result[idx][j], counter)
				}
				counter++
			}
		}
	}
	ut := utils.PrintUtil{}
	ut.Print3DimInt(result)
	realRes := getRealResult(result, candidates, xsize-1, ysize-1)
	return realRes
}

func getRealResult(tag [][][]int, candicates []int, xIdx int, yIdx int) [][]int {
	if xIdx < 0 || yIdx < 0 {
		return make([][]int, 0)
	}
	if len(tag[xIdx][yIdx]) == 0 {
		return make([][]int, 0)
	}
	//else if len(tag[xIdx][yIdx]) == 1 {
	//   length := tag[xIdx][yIdx][0]
	//   arr := make([]int, length)
	//   for i := 0; i < length; i++ {
	//       arr[i] = candicates[xIdx]
	//   }
	//   tmpRes := make([][]int, 1)
	//   tmpRes[0] = arr
	//    ut := utils.PrintUtil{}
	//    ut.Print2DInt(tmpRes)
	//   return tmpRes
	//}
	res := make([][]int, 0)
	nums := tag[xIdx][yIdx]
	//ut := utils.PrintUtil{}
	//ut.PrintArrInt(nums)
	xVal := candicates[xIdx]
	for _, num := range nums {
		//fmt.Println(num, xIdx, yIdx)
		remains := yIdx - num*xVal
		tmp := getRealResult(tag, candicates, xIdx-1, remains)
		if len(tmp) == 0 {
			tmp = make([][]int, 1)
		}
		for idx, _ := range tmp {
			for i := 0; i < num; i++ {
				tmp[idx] = append(tmp[idx], xVal)
			}
		}
		res = myAppend39(res, tmp)
	}
	//fmt.Println()
	return res
}

func myAppend39(res [][]int, tmp [][]int) [][]int {
	for _, t := range tmp {
		res = append(res, t)
	}
	return res
}

func Test39() {
	candidates := []int{2, 3, 5}
	target := 8
	mat := combinationSum(candidates, target)
	ut := utils.PrintUtil{}
	ut.Print2DInt(mat)

	candidates = []int{2, 3, 6, 7}
	target = 7
	mat = combinationSum(candidates, target)
	ut.Print2DInt(mat)

}
