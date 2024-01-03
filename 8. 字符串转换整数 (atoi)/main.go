package main

import "math"

var table = map[string][]string{
	"start":     {"start", "singed", "in_number", "end"},
	"singed":    {"end", "end", "in_number", "end"},
	"in_number": {"end", "end", "in_number", "end"},
	"end":       {"end", "end", "end", "end"},
}

func getCol(c byte) int {
	if c == ' ' {
		return 0
	} else if c == '+' || c == '-' {
		return 1
	} else if c >= '0' && c <= '9' {
		return 2
	} else {
		return 3
	}
}

func myAtoi(s string) int {
	state := "start"

	number := 0
	signed := 1
	for i := 0; i < len(s); i++ {
		state = table[state][getCol(s[i])]
		if state == "in_number" {
			number = number*10 + int(s[i]-'0')
			if signed == 1 {
				number = min(math.MaxInt32, number)
			} else {
				number = min(-math.MinInt32, number)
			}
		} else if state == "singed" {
			if s[i] == '+' {
				signed = 1
			} else {
				signed = -1
			}
		}
	}

	return signed * number
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
