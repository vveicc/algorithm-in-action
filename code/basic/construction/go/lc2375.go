package main

func smallestNumber(pattern string) string {
	n := len(pattern)
	bytes := make([]byte, n+1)
	for i := range bytes {
		bytes[i] = byte(i) + '1' // 要求字典序最小，先构造成升序排列
	}
	for i := 0; i < n; i++ {
		if pattern[i] == 'D' {
			l := i
			for i++; i < n && pattern[i] == 'D'; i++ {
			}
			// 翻转数组bytes中的连续'D'段
			for r := i; l < r; r-- {
				bytes[l], bytes[r] = bytes[r], bytes[l]
				l++
			}
		}
	}
	return string(bytes)
}
