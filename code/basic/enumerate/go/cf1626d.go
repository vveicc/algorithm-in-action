package main

import (
	"bufio"
	. "fmt"
	"math/bits"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		sort.Ints(a)

		ans := cost(n) + 2
		// 枚举与第一组的大小最接近的2的幂次
		for i := 1; ans != 0 && i < n; i <<= 1 {
			cx := sort.SearchInts(a, a[i]) // < x = a[i]
			ans = min(ans, i-cx+cost(n-cx)+1)
			// 枚举与第二组的大小最接近的2的幂次
			for j := 1; ans != 0 && cx+j < n; j <<= 1 {
				cy := sort.SearchInts(a, a[cx+j]) // < y = a[cx+j]
				ans = min(ans, i+j-cy+cost(n-cy))
			}
		}
		Fprintln(out, ans)
	}
}

func cost(x int) int {
	l := bits.Len32(uint32(x))
	if x == 1<<l>>1 {
		return 0
	} else {
		return 1<<l - x
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
