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
		var s string
		Fscan(in, &n, &s)
		a := make([]int, n)
		b := make([]int, n)
		for i := range a {
			a[i] = n - i // 数组a的LIS最短，先构造成降序排列
			b[i] = i + 1 // 数组b的LIS最长，先构造成升序排列
		}
		reverse := func(arr []int, sign byte) {
			for i := 0; i < n-1; i++ {
				if s[i] == sign {
					l := i
					for i++; i < n-1 && s[i] == sign; i++ {
					}
					for r := i; l < r; r-- {
						arr[l], arr[r] = arr[r], arr[l]
						l++
					}
				}
			}
		}
		reverse(a, '<') // 翻转数组a中的连续'<'段
		reverse(b, '>') // 翻转数组b中的连续'>'段
		for _, x := range a {
			Fprint(out, x, " ")
		}
		Fprintln(out)
		for _, x := range b {
			Fprint(out, x, " ")
		}
		Fprintln(out)
	}
}
