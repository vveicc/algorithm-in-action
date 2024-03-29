# 拓扑排序

## LC2127. 参加会议的最多员工数

???+ note "问题描述"
    一个公司准备组织一场会议，邀请名单上有 `2≤n≤1e5` 位员工。员工编号为 `0` 到 `n - 1` 。<br>
    公司准备了一张 **圆形** 的桌子，可以坐下 **任意数目** 的员工。<br>
    每位员工都有一位 **喜欢** 的员工，每位员工 **当且仅当** 他被安排在喜欢员工的旁边，他才会参加会议。<br>
    每位员工喜欢的员工 **不会** 是他自己。<br>
    给你一个下标从 `0` 开始的整数数组 `favorite` ，其中 `favorite[i]` 表示第 `i` 位员工喜欢的员工。<br>
    请你返回参加会议的 **最多员工数目** 。

    在 [LeetCode主站](https://leetcode.com/problems/maximum-employees-to-be-invited-to-a-meeting "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/maximum-employees-to-be-invited-to-a-meeting "困难：2449") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "topo/go/lc2127.go"
        ```

## LC2603. 收集树中金币

???+ note "问题描述"
    有一个 `1≤n≤3e4` 个节点的无向无根树，节点编号从 `0` 到 `n - 1` 。<br>
    给你一个长度为 `n` 的数组 `coins` ，其中 `coins[i]` 等于 `0` 或 `1` ，`1` 表示节点 `i` 处有一个金币。<br>
    给你一个长度为 `n - 1` 的二维整数数组 `edges` ，其中 `edges[i]` = [$a_i$, $b_i$] 表示树中节点 $a_i$ 和 $b_i$ 之间有一条边。<br>
    一开始，你需要选择树中任意一个节点出发。你可以执行下述操作任意次：

    - 收集与当前节点距离 `≤2` 的所有金币，或者
    - 移动到树中一个相邻节点。

    你需要收集树中所有的金币，并回到出发节点，请你返回最少经过的边数。如果你多次经过一条边，每一次经过都会给答案加一。

    在 [LeetCode主站](https://leetcode.com/problems/collect-coins-in-a-tree "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/collect-coins-in-a-tree "困难：2712") 查看该题。

??? info "解题思路"
    **方法一：[换根 DP](../dp/tree/root-changing.md#lc2603-收集树中金币)**

    **方法二：拓扑排序**

    === "Go"
        ```go
        --8<-- "topo/go/lc2603.go"
        ```
