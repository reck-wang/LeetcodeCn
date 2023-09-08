package main

func removeDuplicates(nums []int) int {
	var i, j int
	for ; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			nums[j] = nums[i]
			j++
		}
	}
	nums[j] = nums[i]
	j++
	return j
}
