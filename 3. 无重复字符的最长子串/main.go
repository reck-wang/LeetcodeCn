package main

// use slide window and hash table
func lengthOfLongestSubstring(s string) int {
	idxMap := make(map[byte]int)
	left, maxLen := 0, 0

	for i := 0; i < len(s); i++ {
		if idx, has := idxMap[s[i]]; has {
			left = max(left, idx+1)
		}
		idxMap[s[i]] = i
		maxLen = max(maxLen, i-left+1)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
