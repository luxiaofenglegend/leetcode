package binarySearch

import "fmt"

/*
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
*/

func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	// first, find the smallest one's index
	normalNums, beginIndex := findBeginIndex(nums, target)
	//fmt.Printf("beginIndex:%d\n", beginIndex)
	// second, ordinary binary search
	curIndex := searchInsert33(normalNums, target)
	curIndex = (curIndex + beginIndex) % len(nums)
	if nums[curIndex] == target {
		return curIndex
	}
	return -1
}

func searchInsert33(nums []int, target int) int {
	length := len(nums)
	lo := 0
	hi := length
	var mid int
	for lo < hi {
		mid = (lo + hi) / 2
		if nums[mid] < target {
			lo = mid + 1
			// hi = mid
		} else {
			// lo = mid + 1
			hi = mid
		}
	}
	return hi
}

func findBeginIndex(nums []int, target int) ([]int, int) {
	length := len(nums)
	lo := 0
	hi := length - 1
	mid := 0
	for hi-lo+1 > 2 {
		mid = (lo + hi) / 2
		loVal := nums[lo]
		hiVal := nums[hi]
		midVal := nums[mid]
		if loVal < midVal && midVal < hiVal {
			break
		} else if midVal < hiVal {
			hi = mid
		} else if loVal < midVal {
			lo = mid
		} else {
			fmt.Println("Bad happen!")
		}
	}
	var idx int
	if nums[lo] < nums[hi] {
		idx = lo
	} else {
		idx = hi
	}
	//res := 0
	//if nums[lo] < nums[mid]
	res := nums[idx:]
	res = append(res, nums[0:idx]...)
	return res, idx
}

func Test33() {
	arr1 := []int{4, 5, 6, 7, 0, 1, 2}
	for _, item := range arr1 {
		res1 := search(arr1, item)
		fmt.Println(res1)
	}
	res1 := search(arr1, 8)
	fmt.Println(res1)

	fmt.Println(search([]int{3, 1}, 1))
	fmt.Println(search([]int{5, 1, 3}, 1))
}
