package main

func permute(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	arr := make([]int, 0, n)
	set := make(map[int]bool, n)

	var dfs = func([]int) {}
	dfs = func(arr []int) {
		if len(arr) == n {
			temp := make([]int, n)
			copy(temp, arr)
			ans = append(ans, temp)
			return
		}

		for _, num := range nums {
			if set[num] {
				continue
			}
			arr = append(arr, num)
			set[num] = true
			dfs(arr)
			set[num] = false
			arr = arr[:len(arr)-1]
		}
	}
	dfs(arr)

	return ans
}
