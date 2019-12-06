package divideConquer

import (
	"fmt"
	"math/rand"
)

/*
Example 1:

Input: [3,2,1,5,6,4] and k = 2
Output: 5

Example 2:

Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4

Note that it is the kth largest element in the sorted order, not the kth distinct element.
*/

func findKthLargest(nums []int, k int) int {
	return innner215(nums, k)
}

func innner215(nums []int, k int) int {
	//for _, n := range nums {
	//    fmt.Print(n, " ")
	//}
	//fmt.Println(k, " ")
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 && k == 1 {
		return nums[0]
	}
	pivot := rand.Intn(length-1) + 1
	//pivot := 0
	pVal := nums[pivot]
	nums, pIdx := partition(nums, pVal)
	rightLen := length - pIdx
	if rightLen == k {
		return nums[pIdx]
	} else if rightLen < k {
		return innner215(nums[0:pIdx], k-rightLen)
	} else {
		return innner215(nums[pIdx:length], k)
	}
}

func partition(nums []int, pVal int) ([]int, int) {
	length := len(nums)
	left := make([]int, 0, length)
	right := make([]int, 0, length)
	for _, num := range nums {
		if num < pVal {
			left = append(left, num)
		} else if num == pVal {
			posiible := rand.Float64()
			if posiible > 0.5 {
				left = append(left, num)
			} else {
				right = append(right, num)
			}
		} else {
			right = append(right, num)
		}
	}
	//left =
	return append(left, right...), len(left)
}

func Test215() {
	//arr := []int{3,2,1,5,6,4}
	//fmt.Println(findKthLargest(arr, 2))
	//nums := []int{3,2,3,1,2,4,5,5,6}
	//fmt.Println(findKthLargest(nums, 4))
	//nums1 := []int{99, 99}
	//fmt.Println(findKthLargest(nums1, 1))
	for i := 0; i < 100; i++ {
		nums1 := []int{2, 1}
		fmt.Println(findKthLargest(nums1, 2))
	}

}
