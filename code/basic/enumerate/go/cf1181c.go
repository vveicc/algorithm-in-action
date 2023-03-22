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

	var n, m, ans int
	Fscan(in, &n, &m)
	a := make([]string, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	lec := make([]int, n)
	for j := 0; j < m; j++ { // 枚举右边界
		// lec[i]表示第i行第j列往左连续相同字符的个数
		for i, row := range a {
			if j != 0 && row[j] == row[j-1] {
				lec[i]++
			} else {
				lec[i] = 1
			}
		}

	next:
		for i := 0; i < n; {
			i0 := i      // 中间h行的第一行
			mn := lec[i] // 左侧最短同色长度
			// 处理中间h行
			for i++; i < n && a[i][j] == a[i0][j]; i++ {
				mn = min(mn, lec[i])
			}
			// 此时i指向后h行的第一行
			if h := i - i0; i0 < h || i+h > n {
				continue // 前面或后面不够h行
			} else {
				i0--
				// 处理前h行
				for k := i0; k > i0-h; k-- {
					if a[k][j] != a[i0][j] {
						continue next // 前h行不同色
					}
					mn = min(mn, lec[k])
				}
				// 处理后h行
				for k := i; k < i+h; k++ {
					if a[k][j] != a[i][j] {
						continue next // 后h行不同色
					}
					mn = min(mn, lec[k])
				}
				ans += mn
			}
		}
	}

	Fprintln(out, ans)
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
