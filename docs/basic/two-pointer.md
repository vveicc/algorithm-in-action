# 双指针

## 滑动窗口

### LC2379: 得到 K 个黑块的最少涂色次数

???+ note "问题描述"
    给你一个长度为 $n(1≤n≤100)$ 的字符串 blocks，仅包含字符 'W' 和 'B'，分别表示白色和黑色。<br>
    给你一个整数 $k(1≤k≤n)$ ，表示想要连续黑色块的数目。<br>
    每一次操作中，你可以选择一个白色块将它涂成黑色块。<br>
    请你返回至少出现一次连续 $k$ 个黑色块的最少操作次数。

    在 [LeetCode主站](https://leetcode.com/problems/minimum-recolors-to-get-k-consecutive-black-blocks "Easy")
    或 [力扣中文社区](https://leetcode.cn/problems/minimum-recolors-to-get-k-consecutive-black-blocks "简单：1360") 上查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "two-pointer/sliding-window/go/lc2379.go"
        ```

### LC713: 乘积小于 K 的子数组

???+ note "问题描述"
    给你一个长度为 $n(1≤n≤3e4)$ 的整数数组 $nums(1≤nums[i]≤1e3)$ 和一个整数 $k(0≤k≤1e6)$ 。<br>
    请你返回子数组内所有元素的乘积严格小于 $k$ 的连续子数组的数目。

    在 [LeetCode主站](https://leetcode.com/problems/subarray-product-less-than-k "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/subarray-product-less-than-k "中等") 上查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "two-pointer/sliding-window/go/lc713.go"
        ```
