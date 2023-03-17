# 贪心

## LC1605: 给定行和列的和求可行矩阵

???+ note "问题描述"
    给你两个长度分别为 `m` 和 `n` 的非负整数数组 `rowSum` 和 `colSum` ，其中：
    
    * 数组长度 `1≤m,n≤500` 且两个数组值的范围为 `[0,1e8]`；
    * `rowSum[i]` 是二维矩阵中第 `i` 行元素的和；
    * `colSum[j]` 是二维矩阵中第 `j` 列元素的和。

    请返回大小为 `m x n` 的任意非负整数矩阵，且该矩阵满足 `rowSum` 和 `colSum` 的要求。<br>
    题目保证存在至少一个可行矩阵。

    在 [LeetCode主站](https://leetcode.com/problems/find-valid-matrix-given-row-and-column-sums "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/find-valid-matrix-given-row-and-column-sums "中等：1868") 查看该题。

??? info "解题思路"
    贪心的将 `matrix[i][j]` 的值构造为满足行列和不超过 `rowSum` 和 `colSum` 的最大值。

    === "Go"
        ```go
        --8<-- "greedy/go/lc1605.go"
        ```

## CF1054D: Changing Array

???+ note "问题描述"
    第一行输入正整数 $n(1≤n≤2e5)\ k(1≤k≤30)$ ，第二行输入长为 $n$ 的数组 $a(0≤a[i]≤2^k-1)$ 。<br>
    设 $mask = (1<<k)-1$ ，每次操作你可以把任意 $a[i]$ 修改为 $a[i] \oplus mask$ ，你可以操作任意次（包括 $0$ 次）。<br>
    修改后，最多有多少个 $a$ 的非空连续子数组，其异或和不等于 $0$ ？

    在 [Codeforces](https://codeforces.com/problemset/problem/1054/D "1900")
    或 [洛谷](https://www.luogu.com.cn/problem/CF1054D "省选/NOI-") 查看该题。

??? info "解题思路"
    [这里](https://www.luogu.com.cn/blog/endlesscheng/solution-cf1054d)有一篇很详细的题解。

    === "Go"
        ```go
        --8<-- "greedy/go/cf1054d.go"
        ```
    === "Java"
        ```java
        --8<-- "greedy/java/cf1054d/Main.java"
        ```
