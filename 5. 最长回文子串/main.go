package main

// 中心扩展法
func longestPalindrome(s string) string {
	maxLen, start, end := 0, 0, 0
	n := len(s)
	for i := 0; i < n; i++ {
		var start1, end1 = i, i
		for ; start1 >= 0 && end1 < n && s[start1] == s[end1]; start1, end1 = start1-1, end1+1 {
		}
		start1, end1 = start1+1, end1-1
		len1 := end1 - start1
		if len1 > maxLen {
			maxLen = len1
			start = start1
			end = end1
		}

		var start2, end2 = i, i + 1
		for ; start2 >= 0 && end2 < n && s[start2] == s[end2]; start2, end2 = start2-1, end2+1 {
		}
		start2, end2 = start2+1, end2-1
		len2 := end2 - start2
		if len2 > maxLen {
			maxLen = len2
			start = start2
			end = end2
		}
	}

	return s[start : end+1]
}
