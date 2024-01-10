package main

func generateParenthesis(n int) []string {
	var ans []string
	var cur []byte

	var backTrack func(int, int)
	backTrack = func(left int, right int) {
		if left+right == n*2 {
			ans = append(ans, string(cur))
			return
		}

		if left < n {
			cur = append(cur, '(')
			backTrack(left+1, right)
			cur = cur[:len(cur)-1]
		}

		if right < left {
			cur = append(cur, ')')
			backTrack(left, right+1)
			cur = cur[:len(cur)-1]
		}
	}

	backTrack(0, 0)

	return ans
}
