# 区间修改与懒惰标记

## [[T]](#区间修改与懒惰标记) LG3372. 线段树 1

???+ note "问题描述"
    第一行输入两个正整数 `n m` `(1≤n,m≤1e5)` ，表示数组长度和操作总个数。<br>
    第二行输入长为 `n` 的数组 `a` ，下标从 1 开始。<br>
    接下来 `m` 行，每行输入 3 或 4 个整数，表示一个操作，具体如下：

    1. `1 x y k` ：将数组 `a` 区间 `[x, y]` 内的每个数加上 `k` ；
    2. `2 x y` ：输出数组 `a` 区间 `[x, y]` 内每个数的和，每次输出一行。

    题目保证任意时刻数组 `a` 中所有元素的绝对值之和 `≤1e18` 。

    在 [洛谷](https://www.luogu.com.cn/problem/P3372 "普及/提高-") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "seg/lazy/go/lg3372.go"
        ```

## [[T]](#区间修改与懒惰标记) LG3373. 线段树 2

???+ note "问题描述"
    第一行输入两个正整数 `n m p` `(1≤n,m≤1e5)` ，表示数组长度、操作总个数和模数。<br>
    第二行输入长为 `n` 的数组 `a` ，下标从 1 开始。<br>
    接下来 `m` 行，每行输入 3 或 4 个整数，表示一个操作，具体如下：

    1. `1 x y k` ：将数组 `a` 区间 `[x, y]` 内的每个数乘上 `k` ；
    2. `2 x y k` ：将数组 `a` 区间 `[x, y]` 内的每个数加上 `k` ；
    3. `3 x y` ：输出数组 `a` 区间 `[x, y]` 内每个数的和对 `p` 取模的结果，每次输出一行。

    除样例数据 `p = 38` 外，其余 `p = 571373` 。

    在 [洛谷](https://www.luogu.com.cn/problem/P3373 "普及+/提高") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "seg/lazy/go/lg3373.go"
        ```

## LC2569. 最小无法得到的或值

???+ note "问题描述"
    给你两个长度为 `1≤n≤1e5` 下标从 0 开始的数组 `nums1` 和 `nums2` ，`0≤nums1[i]≤1` `0≤nums2[i]≤1e9` 。<br>
    另外给你一个长度为 `1≤m≤1e5` 的二维数组 `queries` 表示一些操作。总共有 3 种类型的操作：

    1. `queries[i] = [1, x, y]` ：将数组 `nums1` 区间 `[x, y]` 内的所有 0 反转成 1 ，1 反转成 0 ；
    2. `queries[i] = [2, p, 0]` ：对于所有下标 `0≤i<n` ，令 `nums2[i] = nums2[i] + nums1[i] * p` ；
    3. `queries[i] = [3, 0, 0]` ：求 `nums2` 中所有元素的和。

    请你返回一个数组，包含所有第三种操作类型的答案。

    在 [LeetCode主站](https://leetcode.com/problems/handling-sum-queries-after-update "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/handling-sum-queries-after-update "困难：2398") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "seg/lazy/go/lc2569.go"
        ```

## LC2547. 拆分数组的最小代价

???+ note "问题描述"
    给你一个长度为 `1≤n≤1e3` 的整数数组 `nums(0≤nums[i]<n)` 和一个整数 `1≤k≤1e9`。<br>
    将数组拆分成一些非空子数组。拆分的 **代价** 是每个子数组中的 **重要性** 之和。

    令 `trimmed(subarray)` 作为子数组的一个特征，其中所有仅出现一次的数字将会被移除。

    - 例如，`trimmed([3,1,2,4,3,4]) = [3,4,3,4]` 。

    子数组的 **重要性** 定义为 `k + trimmed(subarray).length` 。<br>
    找出并返回拆分 `nums` 的所有可行方案中的最小代价。

    在 [LeetCode主站](https://leetcode.com/problems/minimum-cost-to-split-an-array "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/minimum-cost-to-split-an-array "困难：2020") 查看该题。

??? info "解题思路"
    ??? tip "方法一：动态规划"
        定义 $f_{i}$ 表示拆分 `nums[:i]` 的最小代价，则状态转移如下：

        $$f_{i+1} = \min_{j=0}^i\{f_j + trimmed(nums[j:i+1]).length + k\}$$
        
        === "Go"
            ```go
            --8<-- "seg/lazy/go/lc2547_1.go"
            ```

    ??? tip "方法二：动态规划"
        记 $unique_{j,i}$ 表示 `nums[j:i+1]` 中仅出现一次的数字个数。<br>
        则 $trimmed(nums[j:i+1]).length = i + 1 - j - unique_{j,i}$ ，方法一中的状态转移转化为：

        $$f_{i+1} = i + 1 + k + \min_{j=0}^i\{f_j - j - unique_{j,i}\}$$

        定义 $f'_i = f_i - i$ ，则状态转移进一步转化为：

        $$f'_{i+1} = k + \min_{j=0}^i\{f'_j - unique_{j,i}\}$$

        最终答案为 $f_n = f'_{n} + n$ 。

        === "Go"
            ```go
            --8<-- "seg/lazy/go/lc2547_2.go"
            ```

    ??? tip "进阶：使用线段树优化"
        上述两种方法的时间复杂度都是：$O(n^2)$ ，如果数据范围扩大为 $1≤n≤1e6$ 。<br>
        采用上述方法将会超出时间限制，可以通过线段树进行优化，时间复杂度：$O(n log n)$ 。

        基于方法二，使用线段树维护 $j$ 在 $[0, i]$ 区间内的 $f'_j - unique_{j,i}$ 的最小值。<br>
        其中 $unique_{j,i}$ 与 $i$ 有关，从左至右计算至 $x = nums[i]$ 时：<br>
        记 $x = nums[i]$ 上一次出现的位置为 $pre_x$ ，上上一次出现的位置为 $pre2_x$ ，则：

        - $j$ 在 $[pre_x+1, i]$ 区间的 $unique_{j,i}$ 都要加一；
        - 如果 $x$ 不是首次出现，$j$ 在 $[pre2_x+1, pre_x]$ 区间的 $unique_{j,i}$ 都要减一，相当于把之前的加一撤销。

        通过线段树区间更新即可实现。

        === "Go"
            ```go
            --8<-- "seg/lazy/go/lc2547_3.go"
            ```
