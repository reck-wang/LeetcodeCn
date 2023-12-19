package main

func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	i, j, k := n-2, n-1, n-1

	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
		}

		nums[i], nums[k] = nums[k], nums[i]
	}

	for x, y := j, n-1; x < y; x, y = x+1, y-1 {
		nums[x], nums[y] = nums[y], nums[x]
	}
}
