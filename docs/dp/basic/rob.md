# 打家劫舍

## LC198. 打家劫舍 I

???+ note "问题描述"
    给定一个长度为 `n(1≤n≤100)` 的整数数组 `nums` ，`0≤nums[i]≤400` 表示沿街第 `i` 间房屋内的存放金额。<br>
    如果两间相邻的房屋都被偷，系统会自动报警。计算不触发报警的情况下，能够偷窃到的最高金额。

    在 [LeetCode主站](https://leetcode.com/problems/house-robber "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/house-robber "中等") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "basic/rob/go/lc198.go"
        ```

## LC213. 打家劫舍 II

???+ note "问题描述"
    给定一个长度为 `n(1≤n≤100)` 的整数数组 `nums` ，`0≤nums[i]≤1000` 表示 **围成一圈** 的第 `i` 间房屋内的存放金额。<br>
    如果两间相邻的房屋都被偷，系统会自动报警。计算不触发报警的情况下，能够偷窃到的最高金额。

    在 [LeetCode主站](https://leetcode.com/problems/house-robber-ii "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/house-robber-ii "中等") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "basic/rob/go/lc213.go"
        ```

## LC337. 打家劫舍 III

???+ note "问题描述"
    所有房屋组成一颗二叉树，节点数在 `[1, 1e4]` 范围内，节点值在 `[0, 1e4]` 范围内，表示房屋内的存放金额。<br>
    只能从根节点进入该地区，如果两间直接相连的房屋都被偷，系统会自动报警。计算不触发报警的情况下，能够偷窃到的最高金额。

    在 [LeetCode主站](https://leetcode.com/problems/house-robber-iii "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/house-robber-iii "中等") 查看该题。

??? info "解题思路"
    **树形 DP**

    === "Go"
        ```go
        --8<-- "basic/rob/go/lc337.go"
        ```

## LC2560. 打家劫舍 IV

???+ note "问题描述"
    沿街有一排连续的房屋。每间房屋内都藏有一定的现金。现在有一位小偷计划从这些房屋中窃取现金。<br>
    由于相邻的房屋装有相互连通的防盗系统，所以小偷 **不会窃取相邻的房屋** 。<br>
    小偷的 **窃取能力** 定义为他在窃取过程中能从单间房屋中窃取的 **最大金额** 。<br>
    给你一个长度为 `1≤n≤1e5` 的整数数组 `nums` ，`1≤nums[i]≤1e9` 表示第 `i` 间房屋中存放的现金金额。<br>
    另给你一个整数 `k(1≤k≤(n+1)/2)` ，表示窃贼将会窃取的 **最少** 房屋数。小偷总能窃取至少 `k` 间房屋。<br>
    返回小偷的 **最小** 窃取能力。

    在 [LeetCode主站](https://leetcode.com/problems/house-robber-iv "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/house-robber-iv "中等：2081") 查看该题。

??? info "解题思路"
    **方法一：二分查找 + DP**

    === "Go"
        ```go
        --8<-- "basic/rob/go/lc2560_1.go"
        ```

    **方法二：二分查找 + 贪心**

    === "Go"
        ```go
        --8<-- "basic/rob/go/lc2560_2.go"
        ```

## LC2597. 美丽子集的数目

???+ note "问题描述"
    给你一个长度为 `1≤n≤20` 的正整数数组 `nums` 和一个正整数 `k` ，`1≤nums[i],k≤1000` 。<br>
    如果 `nums` 的子集中，任意两个整数的绝对差均不等于 `k` ，则认为该子集是一个 **美丽** 子集。<br>
    返回数组 `nums` 中 **非空** 且 **美丽** 的子集数目。

    在 [LeetCode主站](https://leetcode.com/problems/the-number-of-beautiful-subsets "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/the-number-of-beautiful-subsets "中等：2023") 查看该题。

??? info "解题思路"
    **方法一：回溯**

    === "Go"
        ```go
        --8<-- "basic/rob/go/lc2597_1.go"
        ```

    **方法二：同余分组 + 打家劫舍 + 乘法原理**

    === "Go"
        ```go
        --8<-- "basic/rob/go/lc2597_2.go"
        ```
