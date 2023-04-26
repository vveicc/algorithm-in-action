# 树形 DP

## LC124. 二叉树中的最大路径和

???+ note "问题描述"
    二叉树中的 **路径** 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 **至多出现一次** 。该路径 **至少包含一个** 节点，且不一定经过根节点。<br>
    **路径和** 是路径中各节点值的总和。给你一个二叉树的根节点 `root` ，返回其 **最大路径和** 。

    在 [LeetCode主站](https://leetcode.com/problems/binary-tree-maximum-path-sum "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/binary-tree-maximum-path-sum "困难") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "tree/go/lc124.go"
        ```

## LC2538. 最大开销

???+ note "问题描述"
    给你一个整数 `1≤n≤1e5` 表示有 n 个节点的无向无根树，节点编号为 `0` 到 `n-1` 。<br>
    给你一个长度为 `n-1` 的二维整数数组 `edges` ，其中 $edges[i] = [a_i, b_i]$ 表示树中节点 $a_i$ 和 $b_i$ 之间有一条边。<br>
    每个节点都有一个价值。给你一个长度为 `n` 的整数数组 `price` ，其中 `1≤price[i]≤1e5` 是第 `i` 个节点的价值。

    一条路径的 **价值和** 是这条路径上所有节点的价值之和。<br>
    你可以选择树中任意一个节点作为根节点 `root` 。选择 `root` 为根的 **开销** 是以 `root` 为起点的所有路径中，**价值和** 最大的一条路径与最小的一条路径的差值。<br>
    请你返回所有节点作为根节点的选择中，**最大** 的 **开销** 为多少。

    在 [LeetCode主站](https://leetcode.com/problems/difference-between-maximum-and-minimum-price-sum "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/difference-between-maximum-and-minimum-price-sum "困难：2398") 查看该题。

??? info "解题思路"
    **方法一：[换根 DP](./root-changing.md#lc2538-最大开销)**

    **方法二：树形 DP**

    问题转化为去掉一个叶子（度为 1）节点的[最大路径和](#lc124-二叉树中的最大路径和)。

    === "Go"
        ```go
        --8<-- "tree/go/lc2538.go"
        ```
