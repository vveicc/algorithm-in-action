package main

import (
	"bufio"
	. "fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &s)
		r := k
		if m := bits.Len32(uint32(n - k + 1)); m < k {
			r = m
		}
		l := k - r
		mask := 1<<r - 1
		has := make([]bool, 1<<r)
		var l1, rt int
		for _, b := range s[:l] {
			l1 += int(b & 1)
		}
		for _, b := range s[l:k] {
			rt = rt<<1 | int(b&1)
		}
		has[rt] = l1 == l
		for i, b := range s[k:] {
			l1 += int(s[i+l]&1) - int(s[i]&1)
			rt = (rt<<1 | int(b&1)) & mask
			if l1 == l {
				has[rt] = true
			}
		}
		rt = mask
		for ; rt != -1 && has[rt]; rt-- {
		}
		if rt == -1 {
			Fprintln(out, "NO")
		} else {
			Fprintf(out, "YES\n%0*b\n", k, rt^mask)
		}
	}
}
