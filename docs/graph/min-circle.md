# 最小环

## LC2608. 图中的最短环

???+ note "问题描述"
    现有一个含 `2≤n≤1e3` 个顶点的无向图，节点编号从 `0` 到 `n - 1` 。<br>
    图中的边由二维整数数组 `edges` 表示，其中 `edges[i]` = [$u_i$, $v_i$] 表示节点 $u_i$ 和 $v_i$ 之间存在一条边。<br>
    每对节点最多通过一条边连接，并且不存在与自身相连的节点。<br>
    返回图中 **最短** 环的长度。如果不存在环，则返回 `-1` 。

    在 [LeetCode主站](https://leetcode.com/problems/shortest-cycle-in-a-graph "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/shortest-cycle-in-a-graph "困难：1904") 查看该题。

??? info "解题思路"
    **方法一：删边 + 最短路**

    === "Go"
        ```go
        --8<-- "min-circle/go/lc2608_1.go"
        ```

    **方法二：枚举起点 + BFS**

    === "Go"
        ```go
        --8<-- "min-circle/go/lc2608_2.go"
        ```
