# 后缀数组

## LC1163. 字典序最大的子串

???+ note "问题描述"
    给你一个长度为 `n(1≤n≤4e5)` 且仅含有小写英文字母的字符串 `s` ，找出它的所有子串并按字典序排列，返回排在最后的那个子串。

    在 [LeetCode主站](https://leetcode.com/problems/last-substring-in-lexicographical-order "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/last-substring-in-lexicographical-order "困难：1864") 查看该题。

??? info "解题思路"
    **方法一：双指针**

    === "Go"
        ```go
        --8<-- "suffix-array/go/lc1163_1.go"
        ```

    **方法二：Z函数（扩展KMP）**

    === "Go"
        ```go
        --8<-- "suffix-array/go/lc1163_2.go"
        ```
