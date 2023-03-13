package main

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	} else {
		m, n := len(num1), len(num2)
		// 当num1和num2均不为0时，其乘积的长度为m+n或m+n-1
		ans := make([]rune, m+n)
		for i, a := range num1 {
			for j, b := range num2 {
				// 逐位相乘，num1[i]*num2[j]的结果累加到ans[i+j+1]
				ans[i+j+1] += (a & 15) * (b & 15)
			}
		}
		for i := m + n - 1; i > 0; i-- {
			if ans[i] > 9 { // 进位
				ans[i-1] += ans[i] / 10
				ans[i] %= 10
			}
		}
		for ans[0] == 0 {
			ans = ans[1:] // 去除前缀0
		}
		for i := range ans {
			ans[i] += '0' // 数字转为字符
		}
		return string(ans)
	}
}
