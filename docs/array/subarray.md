# 子数组运算模板

**适用场景** 对子数组中的元素进行某种运算，计算

1. 数组中所有子数组的运算结果有多少种；
2. 运算结果等于指定值的子数组有多少个；
3. 运算结果等于指定值的子数组的最小长度或最大长度。

进行的运算可以是以下几类：

1. 按位与、按位或；
2. 最大公约数GCD（Greatest Common Divisor）；
3. 最小公倍数LCM（Least Common Multiple）。

**模板原理** 枚举子数组右端点（或左端点），每个运算结果对应一段连续的左端点（或右端点），计算时只需要记录并处理每个连续段的信息。

## LC2447. GCD 为 K 的子数组数目

???+ note "问题描述"
    给你一个长度为 `1≤n≤1e3` 的整数数组 `nums` 和一个整数 `k` ，`1≤nums[i],k≤1e9` 。<br>
    请你统计并返回 `nums` 的子数组中元素的最大公因数等于 `k` 的子数组数目。

    在 [LeetCode主站](https://leetcode.com/problems/number-of-subarrays-with-gcd-equal-to-k "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/number-of-subarrays-with-gcd-equal-to-k "中等：1603") 查看该题。

??? info "解题思路"
    **方法一：暴力枚举**

    === "Go"
        ```go
        --8<-- "subarray/go/lc2447_1.go"
        ```

    **方法二：子数组运算模板**

    === "Go"
        ```go
        --8<-- "subarray/go/lc2447_2.go"
        ```

## LC2470. LCM 为 K 的子数组数目

???+ note "问题描述"
    给你一个长度为 `1≤n≤1e3` 的整数数组 `nums` 和一个整数 `k` ，`1≤nums[i],k≤1e3` 。<br>
    请你统计并返回 `nums` 的子数组中元素的最小公倍数等于 `k` 的子数组数目。

    在 [LeetCode主站](https://leetcode.com/problems/number-of-subarrays-with-lcm-equal-to-k "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/number-of-subarrays-with-lcm-equal-to-k "中等：1560") 查看该题。

??? info "解题思路"
    **方法一：暴力枚举**

    === "Go"
        ```go
        --8<-- "subarray/go/lc2470_1.go"
        ```

    **方法二：子数组运算模板**

    === "Go"
        ```go
        --8<-- "subarray/go/lc2470_2.go"
        ```

## LC898. 子数组按位或的结果种数

???+ note "问题描述"
    给你一个长度为 `1≤n≤5e4` 的整数数组 `arr` ，`0≤arr[i]≤1e9` 。<br>
    请你统计并返回 `arr` 的所有子数组中元素的按位或的结果种数。多次出现的结果仅计算一次。

    在 [LeetCode主站](https://leetcode.com/problems/bitwise-ors-of-subarrays "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/bitwise-ors-of-subarrays "中等：2133") 查看该题。

??? info "解题思路"
    **方法一：子数组运算模板**

    === "Go"
        ```go
        --8<-- "subarray/go/lc898.go"
        ```

## LC2411. 按位或最大的子数组的最小长度 { data-toc-label='LC2411. 按位或最大的子数组的...' }

???+ note "问题描述"
    给你一个长度为 `1≤n≤1e5` 的整数数组 `nums` ，`0≤nums[i]≤1e9` 。<br>
    请你返回一个大小为 `n` 的整数数组 `answer`，其中 `answer[i]` 是开始位置为 `i` ，按位或运算结果最大，且长度最短的子数组的长度。

    在 [LeetCode主站](https://leetcode.com/problems/smallest-subarrays-with-maximum-bitwise-or "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/smallest-subarrays-with-maximum-bitwise-or "中等：1938") 查看该题。

??? info "解题思路"
    **方法一：利用或运算的性质**

    === "Go"
        ```go
        --8<-- "subarray/go/lc2411_1.go"
        ```

    **方法二：子数组运算模板**

    === "Go"
        ```go
        --8<-- "subarray/go/lc2411_2.go"
        ```

## LC1521. 最接近目标值的按位与结果 { data-toc-label='LC1521. 最接近目标值的按位与...' }

???+ note "问题描述"
    给你一个长度为 `1≤n≤1e5` 的整数数组 `arr` 和一个整数 `0≤target≤1e7`，`1≤arr[i]≤1e6` 。<br>
    请你返回 `arr` 的所有非空子数组对应的 `|x - target|` 的最小值，其中 `x` 是子数组元素的按位与结果。

    在 [LeetCode主站](https://leetcode.com/problems/find-a-value-of-a-mysterious-function-closest-to-target "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/find-a-value-of-a-mysterious-function-closest-to-target "困难：2384") 查看该题。

??? info "解题思路"
    **方法一：子数组运算模板**

    === "Go"
        ```go
        --8<-- "subarray/go/lc1521.go"
        ```
