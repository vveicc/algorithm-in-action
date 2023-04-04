# 区间 DP

## LC516. 最长回文子序列

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

## LC1039. 三角剖分的最低得分

???+ note "问题描述"
    给定一个凸 `n(3≤n≤50)` 边形和一个长度为 `n` 的整数数组 `values` 。<br>
    其中 `1≤values[i]≤100` 是凸 `n` 边形第 `i` 个顶点的值（从某一顶点开始，按照顺时针顺序）。<br>
    将凸 `n` 边形剖分为 `n - 2` 个三角形，每个三角形的值是该三角形三个顶点值的乘积。<br>
    凸 `n` 边形三角剖分的分数是进行三角剖分后所有 `n - 2` 个三角形的值之和。<br>
    返回凸 `n` 边形进行三角剖分后可以得到的最低分。

    在 [LeetCode主站](https://leetcode.com/problems/minimum-score-triangulation-of-polygon "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/minimum-score-triangulation-of-polygon "中等：2130") 查看该题。

??? info "解题思路"
    定义 `f[i][j]` 表示由顶点 `i, ..., j` 构成的多边形的最低得分，则状态转移如下：

    $$f[i][j] = \min\limits_{k=i+1}^{j-1}\{f[i][k]+f[k][j]+values[i]\cdot values[j]\cdot values[k]\}$$

    其中 `i, k, j` 三个顶点构成三角形。

    === "Go"
        ```go
        --8<-- "interval/go/lc1039.go"
        ```

## LG1880. [NOI1995] 石子合并

???+ note "问题描述"
    第一行输入正整数 $n(1≤n≤100)$ 表示有 $n$ 堆石子形成一个环。<br>
    第二行输入长为 $n$ 的数组 $a$ ，其中 $0≤a_i≤20$ 表示第 $i$ 堆石子的个数。<br>
    现要将石子有次序地合并成一堆，规定每次只能选相邻的 $2$ 堆合并成新的一堆，并将新的一堆的石子数，记为该次合并的得分。<br>
    请计算将 $n$ 堆石子合并成 $1$ 堆的最小得分和最大得分。第一行输出最小得分，第二行输出最大得分。

    在 [洛谷](https://www.luogu.com.cn/problem/P1880 "普及+/提高") 查看该题。

??? info "解题思路"
    复制一倍，断环成链。

    === "Go"
        ```go
        --8<-- "interval/go/lg1880.go"
        ```
    === "Java"
        ```java
        --8<-- "interval/java/lg1880/Main.java"
        ```

## LC1000. 合并石头的最低成本

???+ note "问题描述"
    有 `n(1≤n≤30)` 堆石头排成一排，第 `i` 堆中有 `1≤stones[i]≤100` 块石头。<br>
    每次移动需要将连续的 `k(2≤k≤30)` 堆石头合并为一堆，而这次移动的成本为这 `k` 堆石头的总块数。<br>
    找出把所有石头合并成一堆的最低成本。如果不可能，返回 `-1` 。

    在 [LeetCode主站](https://leetcode.com/problems/minimum-cost-to-merge-stones "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/minimum-cost-to-merge-stones "困难：2423") 查看该题。

??? info "解题思路"
    定义 $f[i][j]$ 表示将 $[i, j]$ 区间内的石头合并至不能再合并的最低成本。
    
    首先枚举 $i≤x<j$ 分别将 $[i, x]$ 和 $[x+1, j]$ 区间内的石头合并至不能再合并：
    
    $$f[i][j] = \min\limits_{x=i+q\times (k-1)}^{j-1}\{f[i][x]+f[x+1][j]\}$$
    
    如果 `(j-i)%(k-1) == 0` ，则 $[i, j]$ 区间内的石头需要合并成一堆，最低成本为：
    
    $$f[i][j] = f[i][j] + \sum\limits_{p=i}^j stones[p]$$
    
    其中子数组和可以通过前缀和优化。

    === "Go"
        ```go
        --8<-- "interval/go/lc1000.go"
        ```
