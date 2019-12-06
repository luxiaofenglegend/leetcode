package binarySearch

import "fmt"

/*
Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]

Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
*/

func searchRange2(nums []int, target int) []int {
	length := len(nums)
	lo := 0
	hi := length
	mid := 0
	for lo-hi < 0 {
		mid = (lo + hi) / 2
		midVal := nums[mid]
		if midVal < target {
			lo = mid + 1
		} else if midVal > target {
			hi = mid - 1
		} else {
			break
		}
	}
	if mid < 0 || mid >= length {
		return []int{-1, -1}
	}
	if nums[mid] != target {
		return []int{-1, -1}
	}
	// go left, go right
	left := mid
	right := mid
	for left-1 >= 0 && nums[left-1] == target {
		left = left - 1
	}
	for right+1 < length && nums[right+1] == target {
		right = right + 1
	}
	return []int{left, right}
}

func searchRange(nums []int, target int) []int {
	length := len(nums)
	if length == 0 {
		return []int{-1, -1}
	}
	lo := 0
	hi := length
	var mi int
	for lo < hi {
		mi = (lo + hi) / 2
		if nums[mi] > target {
			hi = mi
		} else {
			lo = mi + 1
		}
	}
	end := lo - 1
	begin := lo - 1
	if end < 0 {
		return []int{-1, -1}
	}
	if nums[end] != target {
		return []int{-1, -1}
	}
	for begin-1 >= 0 {
		if nums[begin-1] == target {
			begin -= 1
		} else {
			break
		}

	}
	return []int{begin, end}
}

func Test34() {
	arr1 := []int{3, 3, 3}
	tar1 := 3
	res1 := searchRange(arr1, tar1)
	fmt.Println(res1)
}
