# 二进制集合操作

## 子集枚举

### LC982: 按位与为零的三元组

???+ note "问题描述"
    整数数组 $nums(0≤nums[i]<2^{16})$ 长度为 $n(1≤n≤1000)$ ，返回其中按位与为零的三元组数目。<br>
    即满足 $0 ≤ i,j,k < n$ 且 $nums[i]\ \&\ nums[j]\ \&\ nums[k] == 0$ 的三元组 $(i, j, k)$ 的数目。

    在 [LeetCode主站](https://leetcode.com/problems/triples-with-bitwise-and-equal-to-zero "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/triples-with-bitwise-and-equal-to-zero "困难：2085") 上查看该题。

??? info "解题思路"
    最直接的方法就是暴力枚举三元组，时间复杂度：$O(n^3)$ ，肯定会超时。<br>
    简单优化一下，令 $x = nums[i], y = nums[j], z = nums[k], a = x\ \&\ y$，预处理每个 $a$ 的数量再枚举 $a$ 和 $z$ 。

    === "Go"
        ```go
        --8<-- "binary-set/subset/go/lc982_1.go"
        ```
    === "Java"
        ```java
        --8<-- "binary-set/subset/java/lc982_1/Solution.java"
        ```
    
    时间复杂度：$O(n^2+2^{16}*n)$ ，仍然会超时，需要进一步优化。

    把二进制看做由非0位组成的集合，那么按位与为0，就相当于两个数的二进制集合没有交集。<br>
    令 $c = x \oplus$ 0xffff，则 $c$ 的所有子集与 $x$ 都不存在交集。<br>
    因此，只需要枚举 $c$ 的所有子集即可。

    !!! tip "子集枚举技巧"
        如何高效枚举 $c$ 的子集 $s$ 呢？<br>
        通过 $s = (s-1)\ \&\ c$ 可以按递减顺序高效枚举 $c$ 的子集。<br>
        注：这个二进制子集枚举技巧经常用于子集状压DP中。

    === "Go"
        ```go
        --8<-- "binary-set/subset/go/lc982_2.go"
        ```
    === "Java"
        ```java
        --8<-- "binary-set/subset/java/lc982_2/Solution.java"
        ```
    
    也可以先预处理每个数字的补集的子集的出现次数，再累加 $cnt[x\&y]$ 。

    === "Go"
        ```go
        --8<-- "binary-set/subset/go/lc982_3.go"
        ```
    === "Java"
        ```java
        --8<-- "binary-set/subset/java/lc982_3/Solution.java"
        ```

    进一步优化：根据 $nums$ 数组确定 $cnt$ 的最小长度。

    === "Go"
        ```go
        --8<-- "binary-set/subset/go/lc982_4.go"
        ```
    === "Java"
        ```java
        --8<-- "binary-set/subset/java/lc982_4/Solution.java"
        ```
