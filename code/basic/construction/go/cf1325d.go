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

	var u, v int
	Fscan(in, &u, &v)
	if u > v {
		Fprintln(out, -1)
	} else if u == v {
		if u == 0 {
			Fprintln(out, 0)
		} else {
			Fprintln(out, 1)
			Fprintln(out, u)
		}
	} else {
		if d := v - u; d&1 == 1 {
			Fprintln(out, -1)
		} else if d >>= 1; d&u == 0 {
			Fprintln(out, 2)
			Fprintln(out, u|d, d)
		} else {
			Fprintln(out, 3)
			Fprintln(out, u, d, d)
		}
	}
}
