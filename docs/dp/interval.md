# 区间 DP

## LC516: 最长回文子序列

???+ note "问题描述"
    给你一个仅由小写英文字母组成的字符串 `s(1≤|s|≤1e3)` ，计算其最长的回文子序列的长度。

    在 [LeetCode主站](https://leetcode.com/problems/longest-palindromic-subsequence "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/longest-palindromic-subsequence "中等") 查看该题。

??? info "解题思路"
    **方法一：最长公共子序列**

    记 `t = reverse(s)` ，则问题转化为计算 `s` 和 `t` 的最长公共子序列的长度。<br>
    定义 `f[i][j]` 表示 `s[:i]` 和 `t[:j]` 的最长公共子序列的长度，则状态转移如下：

    * 如果 `s[i] == t[j]` ，则 `f[i+1][j+1] = f[i][j] + 1` ；
    * 如果 `s[i] != t[j]` ，则 `f[i+1][j+1] = max(f[i][j+1], f[i+1][j])` 。

    实现时可以使用滚动数组优化掉第一个维度。

    === "Go"
        ```go
        --8<-- "interval/go/lc516_1.go"
        ```

    **方法二：区间 DP**

    定义 `f[i][j]` 表示 `s[i:j+1]` 的最长回文子序列的长度，则状态转移如下：

    * 如果 `s[i] == s[j]` ，则 `f[i][j] = f[i+1][j-1] + 2` ；
    * 如果 `s[i] != s[j]` ，则 `f[i][j] = max(f[i+1][j], f[i][j-1])` 。

    实现时可以使用滚动数组优化掉第一个维度。

    === "Go"
        ```go
        --8<-- "interval/go/lc516_2.go"
        ```

## LC1039: 三角剖分的最低得分

???+ note "问题描述"
    给定一个凸 `n(3≤n≤50)` 边形和一个长度为 `n` 的整数数组 `values` 。<br>
    其中 `1≤values[i]≤100` 是凸 `n` 边形第 `i` 个顶点的值（从某一顶点开始，按照顺时针顺序）。<br>
    将凸 `n` 边形剖分为 `n - 2` 个三角形，每个三角形的值是该三角形三个顶点值的乘积。<br>
    凸 `n` 边形三角剖分的分数是进行三角剖分后所有 `n - 2` 个三角形的值之和。<br>
    返回凸 `n` 边形进行三角剖分后可以得到的最低分。

    在 [LeetCode主站](https://leetcode.com/problems/minimum-score-triangulation-of-polygon "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/minimum-score-triangulation-of-polygon "中等：2130") 查看该题。

??? info "解题思路"
    **区间 DP**

    定义 `f[i][j]` 表示由顶点 `i, ..., j` 构成的多边形的最低得分，则状态转移如下：

    $$f[i][j] = \min\limits_{k=i+1}^{j-1}\{f[i][k]+f[k][j]+values[i]\cdot values[j]\cdot values[k]\}$$

    其中 `i, k, j` 三个顶点构成三角形。

    === "Go"
        ```go
        --8<-- "interval/go/lc1039.go"
        ```
