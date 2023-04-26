# 贪心

## LC1053. 交换一次的先前排列

???+ note "问题描述"
    给你一个长度为 `n(1≤n≤1e4)` 的正整数数组 `arr` ，其中 `1≤arr[i]≤1e4` 且可能存在重复的元素。<br>
    请你返回可在 **一次交换**（交换两数字 `arr[i]` 和 `arr[j]` 的位置）后得到的、按字典序排列小于 `arr` 的最大排列。<br>
    如果无法这么操作，就请返回原数组。

    在 [LeetCode主站](https://leetcode.com/problems/previous-permutation-with-one-swap "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/previous-permutation-with-one-swap "中等：1633") 查看该题。

??? info "解题思路"
    不失一般性的，假设 `i < j` ，则需要满足 `arr[i] > arr[j]` 才能使得交换后的字典序小于原数组。<br>
    同时，为了使交换后的数组是最大排列，需要让 `i` 尽可能大，`j` 尽可能小，且 `arr[j]` 是 `i` 右侧小于 `arr[i]` 的最大元素。

    === "Go"
        ```go
        --8<-- "greedy/go/lc1053.go"
        ```

## LC1605. 给定行列和求可行矩阵

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

## LC2551. 将珠子放入背包中

???+ note "问题描述"
    给你一个长度为 `1≤n≤1e5` 的整数数组 `weights` ，其中 `1≤weights[i]≤1e9` 是第 `i` 个珠子的重量。<br>
    同时给你一个整数 `k(1≤k≤n)` ，请你按照如下规则将所有的珠子放进 `k` 个背包：

    - 没有背包是空的。
    - 如果第 `i` 个珠子和第 `j` 个珠子在同一个背包里，那么 `[i, j]` 区间的所有珠子都必须在这个背包中。
    - 如果一个背包有 `[i, j]` 区间的所有珠子，那么这个背包的价格是 `weights[i] + weights[j]` 。

    一个珠子分配方案的 **分数** 是所有 `k` 个背包的价格之和。<br>
    请你返回所有分配方案中，**最大分数** 与 **最小分数** 的 **差值** 为多少。

    在 [LeetCode主站](https://leetcode.com/problems/put-marbles-in-bags "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/put-marbles-in-bags "困难：2042") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "greedy/go/lc2551.go"
        ```

## LC2561. 重排水果

???+ note "问题描述"
    你有两个果篮，每个果篮中有 `1≤n≤1e5` 个水果。<br>
    给你两个整数数组 `basket1` 和 `basket2` ，表示两个果篮中每个水果的成本，范围：`[1, 1e9]` 。<br>
    你希望两个果篮相等。为此，可以根据需要多次执行下述操作：

    - 选中两个下标 `i` 和 `j` ，并交换 `basket1` 中的第 `i` 个水果和 `basket2` 中的第 `j` 个水果。
    - 交换的成本是 `min(basket1[i], basket2[j])` 。

    根据果篮中水果的成本进行排序，如果排序后结果完全相同，则认为两个果篮相等。<br>
    返回使两个果篮相等的最小交换成本，如果无法使两个果篮相等，则返回 `-1` 。

    在 [LeetCode主站](https://leetcode.com/problems/rearranging-fruits "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/rearranging-fruits "困难：2222") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "greedy/go/lc2561.go"
        ```

## CF1054D. Changing Array

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
