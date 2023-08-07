package main

import (
	"math/rand"
	"time"
)

// 方法一: 单路快排(面对数据中存在大量重复数据的用例时会超时)
func sortArray(nums []int) []int {
	rand.Seed(time.Now().UnixNano())
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, l int, r int) {
	if l >= r {
		return
	}

	q := partition(nums, l, r)
	quickSort(nums, l, q-1)
	quickSort(nums, q+1, r)
}

func partition(nums []int, l int, r int) int {
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
