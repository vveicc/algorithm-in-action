# 枚举

## LC829: 连续整数求和

???+ note "问题描述"
    给定一个正整数 $n([1,1e9])$ ，返回连续正整数满足所有数字之和为 $n$ 的组数 。 

    在 [LeetCode主站](https://leetcode.com/problems/consecutive-numbers-sum "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/consecutive-numbers-sum "困难：1694") 上查看该题。

??? info "解题思路"
    枚举和为 $n$ 的连续正整数的长度 $k$ 。

    === "Go"
        ```go
        --8<-- "enumerate/go/lc829.go"
        ```
