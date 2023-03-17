package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, k, x, s int
	Fscan(in, &n, &k)
	ans := c2(n + 1) // 异或前缀和数对总数
	mask := 1<<k - 1
	cnt := map[int]int{s: 1}
	for ; n > 0; n-- {
		Fscan(in, &x)
		s ^= x
		cnt[min(s, s^mask)]++ // s和s^mask的个数一起统计
	}
	for _, c := range cnt {
		// 减去相同数对数目，平均分配s和s^mask的个数，使相同异或前缀和数对数目最少
		ans -= c2(c>>1) + c2((c+1)>>1)
	}
	Println(ans)
}

func c2(n int) int64 {
	return (int64(n) * int64(n-1)) >> 1 // n个数字的不同数对数目
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
