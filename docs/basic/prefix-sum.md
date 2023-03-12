# 前缀和 & 差分

## 前缀和

### LC1124: 表现良好的最长时间段

???+ note "问题描述"
    长度为 $n(1≤n≤10^4)$ 的工作时间表 $hours$ ，记录着某位员工每天的工作小时数。
    我们认为当员工一天中的工作小时数大于 8 小时的时候，那么这一天就是「劳累的一天」。
    所谓「表现良好的时间段」，意味在这段时间内，「劳累的天数」严格大于「不劳累的天数」。
    请你返回「表现良好时间段」的最大长度。

    在 [LeetCode主站](https://leetcode.com/problems/longest-well-performing-interval "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/longest-well-performing-interval "中等：1908") 上查看该题。

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

### CF1552D: Array Differentiation

???+ note "问题描述"
    输入 $t(1≤t≤20)$ 表示 $t$ 组数据，每组数据输入 $n(1≤n≤10)$ 和长为 $n$ 的数组 $a(-1e5≤a[i]≤1e5)$。<br>
    如果存在长为 $n$ 的数组 $b$ ，对于任意 $i$ ，都存在 $j$ 和 $k$，使得 $a[i]=b[j]-b[k]$ ，则输出 `YES`，否则输出 `NO`。<br>
    注意 $j$ 可以等于 $k$。

    在 [Codeforces](https://codeforces.com/problemset/problem/1552/D "1800")
    或 [洛谷](https://www.luogu.com.cn/problem/CF1552D "普及+/提高") 上查看该题。

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
