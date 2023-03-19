# 回文字符串

## LC1616: 分割两个字符串得到回文串

???+ note "问题描述"
    给你两个长度均为 `n(1≤n≤1e5)` 的字符串 `a` 和 `b` 。<br>
    请判断是否存在下标 `0≤i≤n` 使得 `a[:i]+b[i:]` 或 `b[:i]+a[i:]` 构成回文字符串。

    在 [LeetCode主站](https://leetcode.com/problems/split-two-strings-to-make-palindrome "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/split-two-strings-to-make-palindrome "中等：1868") 查看该题。

??? info "解题思路"
    **中心扩展法**

    === "Go"
        ```go
        --8<-- "palindrome/go/lc1616.go"
        ```
