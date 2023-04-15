package main

func collectTheCoins(coins []int, edges [][]int) (ans int) {
	n := len(coins)
	ug := make([][]int, n)
	degrees := make([]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		ug[x] = append(ug[x], y)
		ug[y] = append(ug[y], x)
		degrees[x]++
		degrees[y]++
	}

	// 第一遍拓扑排序剪掉没有金币的子树
	q := make([]int, 0, n)
	for x, degree := range degrees {
		if degree == 1 && coins[x] == 0 {
			q = append(q, x)
			degrees[x] = 0 // 剪掉没有金币的叶子节点
		}
	}
	for len(q) != 0 {
		x := q[0]
		q = q[1:]
		for _, y := range ug[x] {
			if degrees[y]--; degrees[y] == 1 && coins[y] == 0 {
				q = append(q, y)
				degrees[y] = 0 // 剪掉没有金币的叶子节点
			}
		}
	}

	// 第二遍拓扑排序剪掉最外面两层
	for x, degree := range degrees {
		if degree == 1 {
			q = append(q, x)
			degrees[x] = 0 // 剪掉最外层
		}
	}

	if len(q) <= 1 {
		return // 剪枝，树中最多剩余1个节点，直接收集
	}

	// 注：仅剪掉最外面两层的写法可以更简单，此写法适用于剪掉最外面多层的情况
	var tmp []int
	for i := 1; i < 2 && len(q) > 0; i++ {
		tmp, q = q, nil
		for _, x := range tmp {
			for _, y := range ug[x] {
				if degrees[y]--; degrees[y] == 1 {
					q = append(q, y)
					degrees[y] = 0 // 剪掉第i+1层
				}
			}
		}
	}

	// 统计答案
	for _, e := range edges {
		if degrees[e[0]] > 0 && degrees[e[1]] > 0 {
			ans += 2
		}
	}
	return
}
