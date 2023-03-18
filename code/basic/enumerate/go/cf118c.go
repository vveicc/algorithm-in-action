package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	s := make([]byte, n)
	Fscan(in, &n, &k, &s)
	var indexes [10][]int
	for i, c := range s {
		indexes[c&15] = append(indexes[c&15], i)
	}
	var mt []byte
	var mc int = 1e5
	for i, pos := range indexes {
		// 替换出k个b
		b := byte(i + '0')
		cnt, cost := len(pos), 0
		t := make([]byte, n)
		copy(t, s)
		for d := 1; cnt < k; d++ {
			// i+d => i 字典序降低，从前往后替换
			if i+d < 10 {
				p := indexes[i+d]
				for j, m := 0, len(p); j < m && cnt < k; j++ {
					t[p[j]] = b
					cost += d
					cnt++
				}
			}
			// i-d => i 字典序升高，从后往前替换
			if i-d >= 0 {
				p := indexes[i-d]
				for j := len(p) - 1; j >= 0 && cnt < k; j-- {
					t[p[j]] = b
					cost += d
					cnt++
				}
			}
		}
		if cost < mc {
			mc, mt = cost, t
		} else if cost == mc && bytes.Compare(mt, t) == 1 {
			mt = t
		}
	}
	Fprintf(out, "%d\n%s", mc, mt)
}
