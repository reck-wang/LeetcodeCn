package main

import "strconv"

var (
	ans      []string
	segments []int
	str      string
)

func restoreIpAddresses(s string) []string {
	segments = make([]int, 4)
	ans = []string{}
	str = s
	dfs(0, 0)
	return ans
}

func dfs(segNo, idx int) {
	if segNo == 4 {
		if idx == len(str) {
			ans = append(ans, getIpAddress())
		}
		return
	}

	if idx == len(str) {
		return
	}

	if str[idx] == '0' {
		segments[segNo] = 0
		dfs(segNo+1, idx+1)
		return
	}

	temp := 0
	for i := idx; i < len(str); i++ {
		temp = temp*10 + int(str[i]-'0')
		if temp > 0 && temp <= 0xff {
			segments[segNo] = temp
			dfs(segNo+1, i+1)
		} else {
			break
		}
	}
}

func getIpAddress() string {
	res := ""
	for i, v := range segments {
		res += strconv.Itoa(v)
		if i != 3 {
			res += "."
		}
	}
	return res
}
