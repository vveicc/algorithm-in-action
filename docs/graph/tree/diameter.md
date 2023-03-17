# 树的直径

## LC1617: 统计子树中城市之间最大距离

???+ note "问题描述"
    给你 $n(2≤n≤15)$ 个城市，编号从 $1$ 到 $n$ 。<br>
    给你一个长度为 $n-1$ 的数组 $edges$ ，其中 $edges[i]=[u_i,v_i]$ 表示城市 $u_i$ 和 $v_i$ 之间有一条双向边。<br>
    题目保证所有城市形成一棵树。<br>
    请你返回一个大小为 $n-1$ 的数组，其中第 $d$ 个元素（从 $1$ 开始）是城市间最大距离恰好等于 $d$ 的子树数目。

    在 [LeetCode主站](https://leetcode.com/problems/count-subtrees-with-max-distance-between-cities "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/count-subtrees-with-max-distance-between-cities "困难：2308") 查看该题。

??? info "解题思路"
    **方法一：枚举子树 + 树形DP**

    枚举城市子集，检查是否是有效子树，如果有效，则通过树形DP计算子树的直径。

    === "Go"
        ```go
        --8<-- "tree/diameter/go/lc1617_1.go"
        ```

    **方法二：枚举子树 + 两次DFS**

    枚举城市子集，检查是否是有效子树，如果有效，则通过两次DFS计算子树的直径。

    === "Go"
        ```go
        --8<-- "tree/diameter/go/lc1617_2.go"
        ```

    **方法三：枚举直径端点**

    枚举直径两端点，计算以这两个端点作为直径的子树个数，注意去重处理。

    === "Go"
        ```go
        --8<-- "tree/diameter/go/lc1617_3.go"
        ```
