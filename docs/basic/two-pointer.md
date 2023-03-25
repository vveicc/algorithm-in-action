# 双指针

## 滑动窗口

### LC713: 乘积小于 K 的子数组

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

### LC1574: 删除最短的子数组使剩余数组有序

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

## 快慢指针

### LC141: 环形链表

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

### LC524: 通过删除字母匹配到字典里最长单词

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

## 利用序列有序性

### LC167: 两数之和 II - 输入有序数组

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
