package main

func consecutiveNumbersSum(n int) (ans int) {
	// k*(k+1) <= n<<1保证k个连续正整数的最小和不超过n
	for k := 1; k*(k+1) <= n<<1; k++ {
		// 判断正整数n是否可以表示成k个连续正整数之和
		if k&1 == 0 {
			// 判断是否可以表示成偶数个连续正整数之和
			if n%k != 0 && (n<<1)%k == 0 {
				ans++
			}
		} else {
			// 判断是否可以表示成奇数个连续正整数之和
			if n%k == 0 {
				ans++
			}
		}
	}
	return
}
