# 双指针

## 滑动窗口

### LC713. 乘积小于 K 的子数组

???+ note "问题描述"
    给你一个长度为 `n(1≤n≤3e4)` 的整数数组 `nums(1≤nums[i]≤1e3)` 和一个整数 `k(0≤k≤1e6)` 。<br>
    请你返回子数组内所有元素的乘积严格小于 `k` 的连续子数组的数目。

    在 [LeetCode主站](https://leetcode.com/problems/subarray-product-less-than-k "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/subarray-product-less-than-k "中等") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "two-pointer/sliding-window/go/lc713.go"
        ```

### CF1535C. Unstable String

???+ note "问题描述"
    第一行输入 `t(1≤t≤1e4)` 表示 `t` 组数据。所有数据的字符串长度之和 `≤2e5`。<br>
    每组数据输入一个长度不超过 `2e5` 的字符串 `s` ，仅包含 `'0'` `'1'` `'?'` 三种字符。

    定义子串是 **不稳定的** 当且仅当子串中任意相邻两数均不相同，如 `101010⋯` 或 `010101⋯` 。<br>
    定义子串是 **好看的** 当且仅当可以将其中的 `?` 改为 `0` 或 `1`（每个 `?` 怎么改是独立的）使得这个子串是不稳定的。<br>
    求字符串中好看的子串个数之和。

    在 [Codeforces](https://codeforces.com/problemset/problem/1535/C "1400")
    或 [洛谷](https://www.luogu.com.cn/problem/CF1535C "普及+/提高") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "two-pointer/sliding-window/go/cf1535c.go"
        ```
    === "Java"
        ```java
        --8<-- "two-pointer/sliding-window/java/cf1535c/Main.java"
        ```

### LC1574. 删除最短的子数组使剩余数组有序 { data-toc-label='LC1574. 删除最短的子数组使...' }

???+ note "问题描述"
    给你一个长度为 `n(1≤n≤1e5)` 的整数数组 `arr(0≤arr[i]≤1e9)` 。<br>
    请你删除一个子数组（可以为空），使得 `arr` 中剩下的元素是 **非递减** 的。<br>
    请你返回满足题目要求的最短子数组的长度。

    在 [LeetCode主站](https://leetcode.com/problems/shortest-subarray-to-be-removed-to-make-array-sorted "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/shortest-subarray-to-be-removed-to-make-array-sorted "中等：1932") 查看该题。

??? info "解题思路"
    **方法一：枚举左端点，移动右端点**

    === "Go"
        ```go
        --8<-- "two-pointer/sliding-window/go/lc1574_1.go"
        ```

    **方法二：枚举右端点，移动左端点**

    === "Go"
        ```go
        --8<-- "two-pointer/sliding-window/go/lc1574_2.go"
        ```

### LC1040. 移动石子直到连续 II

???+ note "问题描述"
    在一个长度 **无限** 的数轴上，有 `3≤n≤1e4` 颗位置互不相同的石子，第 i 颗石子的位置为 `1≤stones[i]≤1e9`。<br>
    如果一颗石子的位置最小/最大，那么该石子被称作 **端点石子** 。<br>
    每个回合，你可以将一颗端点石子拿起并移动到一个未占用的位置，使得该石子不再是一颗端点石子。<br>
    值得注意的是，如果石子像 `stones = [1,2,5]` 这样，你将 **无法** 移动位于位置 5 的端点石子，因为无论将它移动到任何位置（例如 0 或 3），该石子都仍然会是端点石子。<br>
    当你无法进行任何移动时，即这些石子的位置连续时，游戏结束。请计算要使游戏结束，你可以执行的最小和最大移动次数分别是多少？以长度为 2 的数组形式返回答案：`answer = [minimum_moves, maximum_moves]` 。

    在 [LeetCode主站](https://leetcode.com/problems/moving-stones-until-consecutive-ii "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/moving-stones-until-consecutive-ii "中等：2456") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "two-pointer/sliding-window/go/lc1040.go"
        ```

## 快慢指针

### LC141. 环形链表

???+ note "问题描述"
    给你一个链表的头节点 `head` ，判断链表中是否有环。链表中的节点数目范围：`[0,1e4]`。<br>
    如果链表中存在环，则返回 `true` ；否则，返回 `false` 。

    在 [LeetCode主站](https://leetcode.com/problems/linked-list-cycle "Easy")
    或 [力扣中文社区](https://leetcode.cn/problems/linked-list-cycle "简单") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "two-pointer/fast-and-slow/go/lc141.go"
        ```

## 子序列匹配

### LC524. 通过删除字母匹配到字典里最长单词 { data-toc-label='LC524. 通过删除字母匹配到...' }

???+ note "问题描述"
    给你一个长度为 `m(1≤m≤1e3)` 的字符串 `s` 。<br>
    给你一个长度为 `n(1≤n≤1e3)` 的字符串数组 `dictionary` ，每个字符串的长度为 `[1,1e3]` 。<br>
    找出并返回 `dictionary` 中最长的字符串，该字符串可以通过删除 `s` 中的某些字符得到。<br>
    如果答案不止一个，返回长度最长且字母序最小的字符串。如果答案不存在，则返回空字符串。

    在 [LeetCode主站](https://leetcode.com/problems/longest-word-in-dictionary-through-deleting "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/longest-word-in-dictionary-through-deleting "中等") 查看该题。

??? info "解题思路"
    **方法一：双指针**

    === "Go"
        ```go
        --8<-- "two-pointer/subseq-matching/go/lc524_1.go"
        ```

    **方法二：预处理 + 双指针**

    === "Go"
        ```go
        --8<-- "two-pointer/subseq-matching/go/lc524_2.go"
        ```

### LC2565. 最少得分子序列

???+ note "问题描述"
    给你一个长度为 `1≤m≤1e5` 字符串 `s` 和一个长度为 `1≤n≤1e5` 字符串 `t` 。<br>
    你可以从字符串 `t` 中删除任意数目的字符。<br>
    如果没有从字符串 `t` 中删除字符，那么得分为 0 ，否则字符串的得分为 `right - left + 1` ，其中：

    - `left` 为删除字符中的最小下标。
    - `right` 为删除字符中的最大下标。

    请你返回使 `t` 成为 `s` 子序列的最小得分。

    在 [LeetCode主站](https://leetcode.com/problems/subsequence-with-the-minimum-score "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/subsequence-with-the-minimum-score "困难：2432") 查看该题。

??? info "解题思路"
    **方法一：双指针 + 二分查找**

    === "Go"
        ```go
        --8<-- "two-pointer/subseq-matching/go/lc2565_1.go"
        ```

    **方法二：双指针**

    === "Go"
        ```go
        --8<-- "two-pointer/subseq-matching/go/lc2565_2.go"
        ```

## 利用序列有序性

### LC167. 两数之和 II - 有序数组

???+ note "问题描述"
    给你一个长度为 `n(2≤n≤3e4)` 的整数数组 `numbers`（下标从 `1` 开始），该数组已按非递减顺序排列。<br>
    请你从数组中找出下标分别为 `i` 和 `j` 的两个数，满足两数和等于 `target` 且 `i < j`。<br>
    以长度为 `2` 的整数数组 `[i, j]` 的形式返回这两个数的下标 `i` 和 `j` 。<br>
    题目保证存在且仅存在一个有效答案。你所设计的解决方案必须只使用常量级的额外空间。

    在 [LeetCode主站](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted "中等") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "two-pointer/seq-ordering/go/lc167.go"
        ```
