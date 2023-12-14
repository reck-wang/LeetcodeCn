package main

// dp
//func trap(height []int) int {
//	n := len(height)
//
//	leftMax := make([]int, n)
//	leftMax[0] = height[0]
//	for i := 1; i < n; i++ {
//		leftMax[i] = max(leftMax[i-1], height[i])
//	}
//
//	rightMax := make([]int, n)
//	rightMax[n-1] = height[n-1]
//	for i := n - 2; i >= 0; i-- {
//		rightMax[i] = max(rightMax[i+1], height[i])
//	}
//
//	var ans int
//	for i := 0; i < n; i++ {
//		ans += min(leftMax[i], rightMax[i]) - height[i]
//	}
//	return ans
//}

// 双指针
//func trap(height []int) int {
//	n := len(height)
//	l, r := 0, n-1
//	leftMax, rightMax := 0, 0
//	var ans int
//	for l < r {
//		leftMax, rightMax = max(leftMax, height[l]), max(rightMax, height[r])
//		if leftMax < rightMax {
//			ans += leftMax - height[l]
//			l++
//		} else {
//			ans += rightMax - height[r]
//			r--
//		}
//	}
//
//	return ans
//}

// 单调栈
func trap(height []int) int {
	var ans int
	var stack = make([]int, 0)

	for i := 0; i < len(height); i++ {
		for len(stack) != 0 && height[stack[len(stack)-1]] < height[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}

			left := stack[len(stack)-1]
			width := i - left - 1
			high := min(height[left], height[i]) - height[top]
			ans += width * high
		}
		stack = append(stack, i)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
