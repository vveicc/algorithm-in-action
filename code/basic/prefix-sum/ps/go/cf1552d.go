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

	var t, n int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := false
		set := map[int]bool{0: true}
		for i := (1 << n) - 1; i > 0; i-- {
			sum := 0
			for j, x := range a {
				if i>>j&1 == 1 {
					sum += x
				}
			}
			if set[sum] {
				ans = true
				break
			} else {
				set[sum] = true
			}
		}
		if ans {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}
