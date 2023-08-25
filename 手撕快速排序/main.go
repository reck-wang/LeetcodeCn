package main

import (
	"math/rand"
	"time"
)

func sortArray(nums []int) []int {
	rand.Seed(time.Now().UnixNano())
	quickSort(nums, 0, len(nums)-1)
	return nums
}

//func quickSort(nums []int, l int, r int) {
//	if l >= r {
//		return
//	}
//
//	q := partition(nums, l, r)
//	quickSort(nums, l, q-1)
//	quickSort(nums, q+1, r)
//}

// 方法二: 单路快排(面对数据中存在大量重复数据的用例时会超时)
//func partition(nums []int, l int, r int) int {
//  随机生成参考值索引
//	ri := rand.Int()%(r-l+1) + l
//  将参考值交换到最右边
//	nums[ri], nums[r] = nums[r], nums[ri]
//
//	x := nums[r]
//  定义小于等于参考值元素的起始索引
//	left := l
//
//	for i := l; i < r; i++ {
//		if nums[i] <= x {
//          将小于等于参考值的元素逐个移动到左边
//			nums[left], nums[i] = nums[i], nums[left]
//			left++
//		}
//	}
//
//  将参考值置回分好区的数组之中
//	nums[left], nums[r] = nums[r], nums[left]
//
//	return left
//}

// 方法二: 双路快排
//func partition(nums []int, l int, r int) int {
//	ri := rand.Intn(r-l+1) + l
//	nums[l], nums[ri] = nums[ri], nums[l]
//
//	x := nums[l]
//	i, j := l+1, r
//
//	for {
//		for i <= r && nums[i] < x {
//			i++
//		}
//
//		for j > l && nums[j] > x {
//			j--
//		}
//
//		if i > j {
//			break
//		}
//
//		nums[i], nums[j] = nums[j], nums[i]
//		i++
//		j--
//	}
//
//	nums[l], nums[j] = nums[j], nums[l]
//	return j
//}

// 方法三: 三路快排
func quickSort(nums []int, l int, r int) {
	if l >= r {
		return
	}

	ri := rand.Intn(r-l+1) + l
	nums[l], nums[ri] = nums[ri], nums[l]

	x := nums[l]
	lt, gt := l, r+1

	for i := l + 1; i < gt; {
		if nums[i] > x {
			nums[i], nums[gt-1] = nums[gt-1], nums[i]
			gt--
			continue
		}

		if nums[i] < x {
			nums[i], nums[lt+1] = nums[lt+1], nums[i]
			lt++
		}
		i++
	}

	nums[l], nums[lt] = nums[lt], nums[l]

	quickSort(nums, l, lt)
	quickSort(nums, gt, r)
}
