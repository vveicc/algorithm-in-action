# 双指针

## 滑动窗口

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

## 快慢指针

## 子序列匹配

### LC524: 通过删除字母匹配到字典里最长单词

???+ note "问题描述"
    给你一个长度为 $m([1,1e3])$ 的字符串 $s$ 。<br>
    给你一个长度为 $n([1,1e3])$ 的字符串数组 $dictionary$ ，每个字符串的长度为 $[1,1e3]$ 。<br>
    找出并返回 $dictionary$ 中最长的字符串，该字符串可以通过删除 $s$ 中的某些字符得到。<br>
    如果答案不止一个，返回长度最长且字母序最小的字符串。如果答案不存在，则返回空字符串。

    在 [LeetCode主站](https://leetcode.com/problems/longest-word-in-dictionary-through-deleting "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/longest-word-in-dictionary-through-deleting "中等") 上查看该题。

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
