package main

var dirs = [4][2]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	visit := make([][]bool, m)
	for i := range visit {
		visit[i] = make([]bool, n)
	}
	var total = m * n
	var dirIdx = 0
	var ans []int

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if total == 0 {
			return
		}

		total--
		ans = append(ans, matrix[i][j])
		visit[i][j] = true

		nextI, nextJ := i+dirs[dirIdx][0], j+dirs[dirIdx][1]
		if nextI < 0 || nextI >= m || nextJ < 0 || nextJ >= n || visit[nextI][nextJ] {
			dirIdx = (dirIdx + 1) % 4
			nextI, nextJ = i+dirs[dirIdx][0], j+dirs[dirIdx][1]
		}

		dfs(nextI, nextJ)
	}
	dfs(0, 0)
	return ans
}
