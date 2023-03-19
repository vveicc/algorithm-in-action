package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	items := make([][3]int, n)
	for i := range items {
		Fscan(in, &items[i][0], &items[i][1], &items[i][2])
	}

	// 滚动数组优化
	cur := make([]int, m+1)
	pre := make([]int, m+1)
	// 单调递减索引队列
	q := make([]int, m+1)
	for _, item := range items {
		v, w, c := item[0], item[1], item[2]
		if c == -1 || c == 1 { // 0-1 背包
			for j := m; j >= v; j-- {
				cur[j] = max(cur[j], cur[j-v]+w)
			}
		} else if cv := c * v; c == 0 || cv >= m { // 完全背包
			for j := v; j <= m; j++ {
				cur[j] = max(cur[j], cur[j-v]+w)
			}
		} else { // 多重背包，使用单调队列优化
			// 滚动数组
			pre, cur = cur, pre
			// 枚举余数
			for j := 0; j < v; j++ {
				head, tail := 0, -1
				// 枚举余数j的方案：j+v, j+2*v, j+3*v, ...
				for k := j; k <= m; k += v {
					// 初始化为上次计算的结果
					cur[k] = pre[k]
					// 将不在窗口内的索引从队首出队
					for ; head <= tail && k-q[head] > cv; head++ {
					}
					// 使用队首的最大值更新结果
					if head <= tail {
						// 使用k-q[head]的容量装入当前物品
						cur[k] = max(cur[k], pre[q[head]]+(k-q[head])/v*w)
					}
					// 将队尾的价值小于当前方案的索引移除，保证单调递减
					for ; head <= tail && pre[q[tail]]+(k-q[tail])/v*w <= pre[k]; tail-- {
					}
					// 将当前方案的索引入队尾
					tail++
					q[tail] = k
				}
			}
		}
	}

	Fprintln(out, cur[m])
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
