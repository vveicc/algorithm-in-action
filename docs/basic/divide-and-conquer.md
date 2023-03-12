# 递归 & 分治

## LC1096: 花括号展开 II

???+ note "问题描述"
    花括号展开的表达式是一个由花括号、逗号和小写英文字母组成的字符串。<br>
    具体的语法规则请前往力扣查看。<br>
    给你一个长度为 $n(1≤n≤60)$ 的合法表达式，返回它所表示的所有字符串组成的有序列表。

    在 [LeetCode主站](https://leetcode.com/problems/brace-expansion-ii "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/brace-expansion-ii "困难：2349") 上查看该题。

??? info "解题思路"
    **方法一：自顶向下递归**

    === "Go"
        ```go
        --8<-- "divide-and-conquer/go/lc1096_1.go"
        ```

    **方法二：自底向上展开**

    === "Go"
        ```go
        --8<-- "divide-and-conquer/go/lc1096_2.go"
        ```
