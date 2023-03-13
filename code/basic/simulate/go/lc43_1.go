package main

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	} else {
		ans := "0"
		m, n := len(num1), len(num2)
		//   num2
		// x num1
		// 取出num1中的每一位分别与num2相乘，并累加结果
		for i := 0; i < m; i++ {
			if a := num1[m-i-1] & 15; a != 0 {
				var x byte
				s := make([]byte, n+1)
				for j := 0; j < n; j++ {
					x += a * (num2[n-j-1] & 15)
					s[n-j] = '0' + x%10
					x /= 10 // 进位
				}
				if x > 0 {
					s[0] = '0' + x
				} else {
					s = s[1:]
				}
				for j := 0; j < i; j++ {
					s = append(s, '0')
				}
				// 累加结果
				ans = add(ans, string(s))
			}
		}
		return ans
	}
}

func add(num1 string, num2 string) string {
	if m, n := len(num1), len(num2); m > n {
		return add(num2, num1)
	} else {
		var x byte
		ans := make([]byte, n+1)
		for i := 0; i < n; i++ {
			if i < m {
				x += num1[m-i-1] & 15
			}
			x += num2[n-i-1] & 15
			ans[n-i] = '0' + x%10
			x /= 10 // 进位
		}
		if x > 0 {
			ans[0] = '0' + x
			return string(ans)
		} else {
			return string(ans[1:])
		}
	}
}
