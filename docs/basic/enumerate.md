# 枚举

## LC1615: 最大网络秩

???+ note "问题描述"
    给你 `n(2≤n≤100)` 座城市，编号从 `0` 开始。<br>
    给你一个数组 `roads` ，其中 $roads[i] = [a_i, b_i]$ 都表示在城市 $a_i$ 和 $b_i$ 之间有一条双向道路。<br>
    题目保证每对城市之间最多只有一条道路相连。<br>
    任意两座不同城市构成的城市对的网络秩定义为：与这两座城市直接相连的道路总数。<br>
    如果存在一条道路直接连接这两座城市，则这条道路只计算一次。<br>
    请返回所有不同城市对中的最大网络秩。

    在 [LeetCode主站](https://leetcode.com/problems/maximal-network-rank "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/maximal-network-rank "中等：1522") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "enumerate/go/lc1615.go"
        ```

## LC829: 连续整数求和

???+ note "问题描述"
    给定一个正整数 `n(1≤n≤1e9)` ，返回连续正整数满足所有数字之和为 `n` 的组数 。

    在 [LeetCode主站](https://leetcode.com/problems/consecutive-numbers-sum "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/consecutive-numbers-sum "困难：1694") 查看该题。

??? info "解题思路"
    枚举和为 `n` 的连续正整数的长度 `k` 。

    === "Go"
        ```go
        --8<-- "enumerate/go/lc829.go"
        ```
