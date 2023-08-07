package main

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSort(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSort(nums []int, l int, r int, idx int) int {
	q := randomPartition(nums, l, r)
	if q == idx {
		return nums[q]
	} else if q < idx {
		return quickSort(nums, q+1, r, idx)
	}

	return quickSort(nums, l, q-1, idx)
}

func randomPartition(nums []int, l int, r int) int {
	ri := rand.Int()%(r-l+1) + l
	nums[ri], nums[r] = nums[r], nums[ri]

	x := nums[r]
	left := l

	for i := l; i < r; i++ {
		if nums[i] <= x {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}

	nums[left], nums[r] = nums[r], nums[left]

	return left
}
