package main

import "bytes"

func findLexSmallestString(s string, a int, b int) string {
	n := len(s)
	c1 := 10 / gcd(10, a) // 累加c1次后复原
	c2 := n / gcd(n, b)   // 轮转c2次后复原
	ans := []byte(s)
	tmp := []byte(s + s) // 方便截取轮转后的字符串
	rotate := func() {
		for i := 0; i < c2; i++ {
			l := ((n-i*b)%n + n) % n // 轮转i次后的起始位置
			if bytes.Compare(ans, tmp[l:l+n]) == 1 {
				copy(ans, tmp[l:l+n])
			}
		}
	}
	for i := 0; i < c1; i++ {
		// 累加奇数索引位置
		for j := 1; j < n<<1; j += 2 {
			tmp[j] = (tmp[j]&15+byte(a))%10 + '0'
		}
		if b&1 == 0 {
			rotate() // 执行轮转操作，取字典序最小的字符串
		} else {
			for j := 0; j < c1; j++ {
				// 累加偶数索引位置
				for k := 0; k < n<<1; k += 2 {
					tmp[k] = (tmp[k]&15+byte(a))%10 + '0'
				}
				rotate() // 执行轮转操作，取字典序最小的字符串
			}
		}
	}
	return string(ans)
}

func gcd(x, y int) int {
	for y > 0 {
		x, y = y, x%y
	}
	return x
}
