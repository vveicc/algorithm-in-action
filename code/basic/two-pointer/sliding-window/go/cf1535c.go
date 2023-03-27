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

	var T int
	for Fscan(in, &T); T > 0; T-- {
		var s string
		Fscan(in, &s)
		var ans int64
		// prev[0]表示前一个与索引奇偶性相同的非'?'字符（'0'或'1'）的索引
		// prev[1]表示前一个与索引奇偶性不同的非'?'字符（'0'或'1'）的索引
		prev := []int{-1, -1}
		for i, c := range s {
			if c != '?' {
				prev[(i&1)^int(c&1)] = i
			}
			// (prev[0], i]区间内的非'?'字符（'0'或'1'）与索引奇偶性均不同
			// (prev[1], i]区间内的非'?'字符（'0'或'1'）与索引奇偶性均相同
			ans += int64(i - min(prev[0], prev[1]))
		}
		Fprintln(out, ans)
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
