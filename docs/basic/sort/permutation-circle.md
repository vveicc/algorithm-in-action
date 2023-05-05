# 置换环

**置换环** 用于计算通过 **两两交换** 将排列变为有序的最少交换次数。<br>
元素互不相同的数组可以通过离散化转为排列。

## LC2471. 逐层排序所需的最少操作数目 { data-toc-label='LC2471. 逐层排序所需的最少...' }

???+ note "问题描述"
    给你一个 **值互不相同** 的二叉树的根节点 `root` ，节点个数范围：`[1, 1e5]` ，节点值范围：`[1, 1e5]` 。<br>
    在一步操作中，你可以选择 **同一层** 上任意两个节点，交换这两个节点的值。<br>
    返回每一层按 **严格递增顺序** 排序所需的最少操作数目。

    在 [LeetCode主站](https://leetcode.com/problems/minimum-number-of-operations-to-sort-a-binary-tree-by-level "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/minimum-number-of-operations-to-sort-a-binary-tree-by-level "中等：1635") 查看该题。

??? info "解题思路"
    **方法一：DFS + 离散化 + 置换环**

    === "Go"
        ```go
        --8<-- "sort/permutation-circle/go/lc2471_1.go"
        ```

    **方法二：BFS + 离散化 + 置换环**

    === "Go"
        ```go
        --8<-- "sort/permutation-circle/go/lc2471_2.go"
        ```
