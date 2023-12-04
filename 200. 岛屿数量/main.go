package main

type UnionFind struct {
	Count  int
	Parent []int
	Rank   []int
}

func NewUnionFind(grid [][]byte) *UnionFind {
	m := len(grid)
	n := len(grid[0])

	u := &UnionFind{
		Parent: make([]int, n*m),
		Rank:   make([]int, n*m),
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				u.Parent[i*n+j] = i*n + j
				u.Count++
			}
			u.Rank = append(u.Rank, 0)
		}
	}

	return u
}

func (u *UnionFind) union(i, j int) {
	root1 := u.find(i)
	root2 := u.find(j)

	if root1 != root2 {
		if u.Rank[root1] < u.Rank[root2] {
			u.Parent[root1] = root2
		} else if u.Rank[root1] > u.Rank[root2] {
			u.Parent[root2] = root1
		} else {
			u.Parent[root2] = root1
			u.Rank[root1]++
		}
		u.Count--
	}
}

func (u *UnionFind) find(i int) int {
	if u.Parent[i] != i {
		u.Parent[i] = u.find(u.Parent[i])
	}
	return u.Parent[i]
}

var dirs = [...][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	u := NewUnionFind(grid)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				for _, dir := range dirs {
					nextI, nextJ := i+dir[0], j+dir[1]
					if nextI >= 0 && nextI < m && nextJ >= 0 && nextJ < n && grid[nextI][nextJ] == '1' {
						u.union(i*n+j, nextI*n+nextJ)
					}
				}
			}
		}
	}

	return u.Count
}
