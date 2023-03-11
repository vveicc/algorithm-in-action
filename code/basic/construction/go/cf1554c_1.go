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
			i, x := 1, 0
			for m++; m>>i != 0; i++ {
			}
			for ; i >= 0; i-- {
				if n>>i&1 == 0 && m>>i&1 == 1 {
					x |= 1 << i
				} else if n>>i&1 == 1 && m>>i&1 == 0 {
					break
				}
			}
			Fprintln(out, x)
		}
	}
}
