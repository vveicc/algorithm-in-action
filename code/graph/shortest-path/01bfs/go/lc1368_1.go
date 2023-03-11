package main

func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	mn := m * n
	// dp[x][y]表示从起点到位置(x, y)的最小代价
	dp := make([][]int, m)
	vis := make([][]bool, m)
	for i := range dp {
		dp[i] = make([]int, n)
		vis[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = mn
		}
	}
	dp[0][0] = 0

	q := deque{{0}}
	for q.size() != 0 {
		if x, y := q.remove(0); !vis[x][y] {
			vis[x][y] = true
			for i, d := range dirs {
				nx, ny := x+d[0], y+d[1]
				if 0 <= nx && nx < m && 0 <= ny && ny < n {
					cost := 0
					if i+1 != grid[x][y] {
						cost = 1
					}
					if dis := dp[x][y] + cost; dis < dp[nx][ny] {
						dp[nx][ny] = dis
						q.add(cost, nx, ny)
					}
				}
			}
		}
	}
	return dp[m-1][n-1]
}

var dirs = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

const shift, mask = 8, (1 << 8) - 1

// 两个 slice 头对头实现双端队列
type deque [2][]int

// size 队列中的元素个数
func (q *deque) size() int { return len(q[0]) + len(q[1]) }

// add 向队列中添加坐标(x, y)。i == 0在队首添加；i == 1在队尾添加
func (q *deque) add(i, x, y int) { q[i] = append(q[i], x<<shift|y) }

// remove 从队列中移除一个坐标(x, y)并返回。i == 0从队首移除；i == 1从队尾移除
func (q *deque) remove(i int) (x, y int) {
	var xy int
	if tail := len(q[i]) - 1; tail == -1 {
		xy, q[i^1] = q[i^1][0], q[i^1][1:]
	} else {
		xy, q[i] = q[i][tail], q[i][:tail]
	}
	return xy >> shift, xy & mask
}
