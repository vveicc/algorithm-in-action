package main

func minimumObstacles(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	mn := m * n
	// dp[i][j]表示到达grid[i][j]需要移除的障碍物的最小数目，即最小代价
	dp := make([][]int, m)
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		vis[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			dp[i][j] = mn
		}
	}
	dp[0][0] = 0

	// 0-1 BFS，(0, 0)的最小代价为0，入队首
	q := deque{{0}}
	for q.size() != 0 {
		if x, y := q.remove(0); !vis[x][y] {
			vis[x][y] = true
			for _, delta := range dirs {
				nx, ny := x+delta[0], y+delta[1]
				if 0 <= nx && nx < m && 0 <= ny && ny < n {
					cost := grid[nx][ny]
					if dis := dp[x][y] + cost; dis < dp[nx][ny] {
						// 更新到达(nx, ny)的最小代价
						dp[nx][ny] = dis
						// 空格代价为0入队首，障碍代价为1入队尾，优先搜索代价小的坐标
						q.add(cost, nx, ny)
					}
				}
			}
		}
	}
	return dp[m-1][n-1]
}

var dirs = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

const shift, mask = 20, (1 << 20) - 1

// 两个slice头对头实现deque
type deque [2][]int

// size 队列中的元素个数
func (q *deque) size() int {
	return len(q[0]) + len(q[1])
}

// add 向队列中添加坐标(x, y)。i == 0在队首添加；i == 1在队尾添加
func (q *deque) add(i, x, y int) {
	q[i] = append(q[i], (x<<shift)|y)
}

// remove 从队列中移除一个坐标(x, y)并返回。i == 0从队首移除；i == 1从队尾移除
func (q *deque) remove(i int) (x, y int) {
	var xy int
	if tail := len(q[i]) - 1; tail > -1 {
		xy, q[i] = q[i][tail], q[i][:tail]
	} else {
		// 0^1 = 1, 1^1 = 0
		xy, q[i^1] = q[i^1][0], q[i^1][1:]
	}
	return xy >> shift, xy & mask
}
