package main

func countDigitOne(n int) int {
	return countDigit(n, 1)
}

func countDigit(n, d int) (ans int) {
	// pow10k表示10^k
	for pow10k := 1; n >= pow10k; pow10k *= 10 {
		// 累加pow10k数位出现数字d的个数
		// pow10k == 1, 10, 100, ... 分别累加 个位, 十位, 百位, ... 出现数字d的个数
		ans += n/(10*pow10k)*pow10k + min(max(n%(10*pow10k)-d*pow10k+1, 0), pow10k)
	}
	return
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
