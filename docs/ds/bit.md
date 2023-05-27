# 树状数组

## CF652D. Nested Segments

???+ note "问题描述"
    第一行输入整数 $n (1 ≤ n ≤ 2e5)$ ，接下来 $n$ 行输入 $n$ 个闭区间，区间左右端点范围在 $[-1e9,1e9]$，所有区间端点互不相同。<br>
    对于每个区间，输出它包含多少个其它的区间。

    在 [Codeforces](https://codeforces.com/problemset/problem/652/D "1800")
    或 [洛谷](https://www.luogu.com.cn/problem/CF652D "普及+/提高") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "bit/go/cf652d.go"
        ```
    === "Java"
        ```java
        --8<-- "bit/java/cf652d/Main.java"
        ```

## LC1626. 无矛盾的最佳球队

???+ note "问题描述"
    假设你是球队的经理，你想组合一支总体得分最高的球队。球队的得分是球队中所有球员的分数总和。<br>
    球队中的矛盾会限制球员的发挥，所以必须选出一支 **没有矛盾** 的球队。如果一名年龄较小球员的分数 **严格大于** 一名年龄较大的球员，则存在矛盾。同龄球员之间不会发生矛盾。<br>
    给你两个长度为 `n(1≤n≤1e3)` 的数组 `scores` 和 `ages`，其中每组 `scores[i]([1,1e6])` 和 `ages[i]([1,1e3])` 表示第 `i` 名球员的分数和年龄。请你返回 **所有可能的无矛盾球队中得分最高那支的分数** 。

    在 [LeetCode主站](https://leetcode.com/problems/best-team-with-no-conflicts "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/best-team-with-no-conflicts "中等：2027") 查看该题。

??? info "解题思路"
    **方法一：动态规划**

    === "Go"
        ```go
        --8<-- "bit/go/lc1626_1.go"
        ```

    **方法二：树状数组**

    如果数组长度 `n` 增大至 `1e5` ，方法一就可能会超时。<br>
    需要使用树状数组（或线段树）将时间复杂度降为：$O(n\log{n} + n\log{U})$。其中 $U={max}(ages)$ 。<br>
    如果 `ages[i]` 的范围较大，可以通过离散化将复杂度降为：$O(n\log{n})$ 。<br>

    === "Go"
        ```go
        --8<-- "bit/go/lc1626_2.go"
        ```

## LC2659. 将数组清空

???+ note "问题描述"
    给你一个包含若干 **互不相同** 整数的数组 nums ，你需要执行以下操作 **直到数组为空** ：

    - 如果数组中第一个元素是当前数组中的 **最小值** ，则删除它。
    - 否则，将第一个元素移动到数组的 **末尾** 。

    请你返回需要多少次操作使 nums 为空。

    在 [LeetCode主站](https://leetcode.com/problems/make-array-empty "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/make-array-empty "困难：2282") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "bit/go/lc2659.go"
        ```
