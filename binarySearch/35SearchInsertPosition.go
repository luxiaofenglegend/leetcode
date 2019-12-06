package binarySearch

func searchInsert(nums []int, target int) int {
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
