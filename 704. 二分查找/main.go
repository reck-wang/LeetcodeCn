package main

func search(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid
		} else {
			return mid
		}
	}

	return -1
}
