# 前缀和 & 差分

## 前缀和

### LC1124. 表现良好的最长时间段

???+ note "问题描述"
    长度为 $n(1≤n≤10^4)$ 的工作时间表 $hours$ ，记录着某位员工每天的工作小时数。<br>
    我们认为当员工一天中的工作小时数大于 8 小时的时候，那么这一天就是「劳累的一天」。<br>
    所谓「表现良好的时间段」，意味在这段时间内，「劳累的天数」严格大于「不劳累的天数」。<br>
    请你返回「表现良好时间段」的最大长度。

    在 [LeetCode主站](https://leetcode.com/problems/longest-well-performing-interval "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/longest-well-performing-interval "中等：1908") 查看该题。

??? info "解题思路"
    **方法一：前缀和 + 哈希表**

    === "Go"
        ```go
        --8<-- "prefix-sum/ps/go/lc1124_1.go"
        ```
    
    **方法二：前缀和 + 单调栈**

    进阶：如果数组中的值不只有 $1$ 和 $-1$，那么如何计算和至少为 $k$ 的最长子数组？

    === "Go"
        ```go
        --8<-- "prefix-sum/ps/go/lc1124_2.go"
        ```

### LC2488. 统计中位数为 K 的子数组 { data-toc-label='LC2488. 统计中位数为 K 的...' }

???+ note "问题描述"
    给你一个长度为 `n(1≤n≤1e5)` 的数组 `nums` ，该数组由从 `1` 到 `n` 的不同整数组成。<br>
    另给你一个正整数 `k(1≤k≤n)` 。统计并返回 `nums` 中的中位数等于 `k` 的非空子数组的数目。

    在 [LeetCode主站](https://leetcode.com/problems/count-subarrays-with-median-k "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/count-subarrays-with-median-k "困难：1999") 查看该题。

??? info "解题思路"
    **前缀和 + 哈希表**

    实现时可以使用数组代替哈希表，空间换时间。

    === "Go"
        ```go
        --8<-- "prefix-sum/ps/go/lc2488.go"
        ```

### LC1590. 使数组和能被 P 整除

???+ note "问题描述"
    给你一个长度为 $n(1≤n≤10^5)$ 的正整数数组 $nums(1≤nums[i]≤10^9)$ 。<br>
    请你移除一个最短的子数组（可以为空），使得剩余元素的和能被 $p(1≤p≤10^9)$ 整除。<br>
    不允许将整个数组都移除。<br>
    请你返回你需要移除的最短子数组的长度，如果无法满足题目要求，返回 $-1$ 。

    在 [LeetCode主站](https://leetcode.com/problems/make-sum-divisible-by-p "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/make-sum-divisible-by-p "中等：2039") 查看该题。

??? info "解题思路"
    **前缀和 + 哈希表**

    === "Go"
        ```go
        --8<-- "prefix-sum/ps/go/lc1590.go"
        ```

### CF1552D. Array Differentiation { data-toc-label='CF1552D. Array Differentiat...' }

???+ note "问题描述"
    输入 $t(1≤t≤20)$ 表示 $t$ 组数据，每组数据输入 $n(1≤n≤10)$ 和长为 $n$ 的数组 $a(-1e5≤a[i]≤1e5)$。<br>
    如果存在长为 $n$ 的数组 $b$ ，对于任意 $i$ ，都存在 $j$ 和 $k$，使得 $a[i]=b[j]-b[k]$ ，则输出 `YES`，否则输出 `NO`。<br>
    注意 $j$ 可以等于 $k$。

    在 [Codeforces](https://codeforces.com/problemset/problem/1552/D "1800")
    或 [洛谷](https://www.luogu.com.cn/problem/CF1552D "普及+/提高") 查看该题。

??? info "解题思路"
    数组 $a$ 中元素的顺序不影响答案。<br>
    数组 $a$ 中元素的正负不影响答案，每个元素都可以任意取反。<br>
    首先将数组 $b$ 构造为数组 $a$ 前 $n-1$ 个元素的前缀和，则有：
    
    * 当 $i < n-1$ 时，$a_i = b_{i+1} - b_i$ ；
    * 当 $i = n-1$ 时，如果存在区间 $[l, r]$ ，满足 $\sum\limits_{i=l}^r a_i = a_n$ ，则 $a_n = b_{r+1} - b_l$ 。
    
    因此，只要存在区间 $[l, r]$ ，满足 $\sum\limits_{i=l}^r a_i = a_n$ 即可。<br>
    因为数组 $a$ 中的元素可以任意交换顺序、任意取反，所以上述条件等价于数组 $a$ 存在两个元素和相等的子集。

    === "Go"
        ```go
        --8<-- "prefix-sum/ps/go/cf1552d.go"
        ```
    === "Java"
        ```java
        --8<-- "prefix-sum/ps/java/cf1552d/Main.java"
        ```

## 差分

### 一维差分

对于一个给定的数列 $A$ ，其差分数列 $B$ 根据如下公式计算：

$$B_i = A_i - A_{i-1}$$

其中 $1≤i≤n$ 且 $B_0 = A_0$ 。

### 二维差分

对于一个给定的矩阵 $A$ ，其差分矩阵 $B$ 根据如下公式计算：

$$B_{i,j} = A_{i,j} - A_{i-1,j} - A_{i,j-1} + A_{i-1,j-1}$$

其中 $1≤i,j≤n$ 。

[这里](https://zhuanlan.zhihu.com/p/439268614)有一篇很详细的教程。

#### LG3397. 地毯

???+ note "问题描述"
    第一行输入两个正整数 $n,m$ $(1≤n,m≤1e3)$。表示在 $n \times n$ 的格子上有 $m$ 个地毯。<br>
    接下来 $m$ 行，每行输入四个整数：$x_1,y_1,x_2,y_2$ 。<br>
    表示一块地毯，左上角是 $(x_1,y_1)$ ，右下角是 $(x_2,y_2)$ 。坐标从 $1$ 开始。<br>
    请输出一个 $n \times n$ 的矩阵 $grid$，其中 $grid[i][j]$ 表示这个格子被多少个地毯覆盖。

    在 [洛谷](https://www.luogu.com.cn/problem/P3397 "普及-") 查看该题。

??? info "解题思路"
    **方法一：按行差分**

    === "Go"
        ```go
        --8<-- "prefix-sum/diff/go/lg3397_1.go"
        ```
    
    **方法二：二维差分**

    === "Go"
        ```go
        --8<-- "prefix-sum/diff/go/lg3397_2.go"
        ```

#### LC2132. 用邮票贴满网格图

???+ note "问题描述"
    给你一个 `m x n` 的二进制矩阵 `grid` ，每个格子要么为 `0` （空）要么为 `1`（被占据）。<br>
    给你邮票的尺寸为 `stampHeight x stampWidth` 。`1≤m,n,stampHeight,stampWidth≤1e5` 且 `1≤m*n≤2e5` 。<br>
    我们想将邮票贴进二进制矩阵中，且满足以下 **限制** 和 **要求** ：

    1. 覆盖所有 **空** 格子。
    2. 不覆盖任何 **被占据** 的格子。
    3. 我们可以放入任意数目的邮票。
    4. 邮票可以相互有 **重叠** 部分。
    5. 邮票不允许 **旋转** 。
    6. 邮票必须完全在矩阵 **内** 。

    如果在满足上述要求的前提下，可以放入邮票，请返回 `true` ，否则返回 `false` 。

    在 [LeetCode主站](https://leetcode.com/problems/stamping-the-grid "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/stamping-the-grid "困难：2364") 查看该题。

??? info "解题思路"
    **方法一：二维前缀和 + 二维差分**

    === "Go"
        ```go
        --8<-- "prefix-sum/diff/go/lc2132.go"
        ```
