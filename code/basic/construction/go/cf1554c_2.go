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

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		if Fscan(in, &n, &m); n > m {
			Fprintln(out, 0)
		} else {
			x := n ^ m
			if i := 0; n&x == 0 {
				for y := n | x; y>>i&1 == 1; i++ {
				}
				Fprintln(out, (x>>i|1)<<i) // 通过第四种情况修改x
			} else {
				for y := n & x; y>>i != 0; i++ {
				}
				Fprintln(out, x>>i<<i) // 通过第三种情况修改x
			}
		}
	}
}
