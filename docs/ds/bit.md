# 树状数组

## CF652D: Nested Segments

???+ note "问题描述"
    第一行输入整数 $n (1 ≤ n ≤ 2e5)$ ，接下来 $n$ 行输入 $n$ 个闭区间，区间左右端点范围在 $[-1e9,1e9]$，所有区间端点互不相同。<br>
    对于每个区间，输出它包含多少个其它的区间。

    在 [Codeforces](https://codeforces.com/problemset/problem/652/D "1800")
    或 [洛谷](https://www.luogu.com.cn/problem/CF652D "普及+/提高") 上查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "bit/go/cf652d.go"
        ```
    === "Java"
        ```java
        --8<-- "bit/java/cf652d/Main.java"
        ```
