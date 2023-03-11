# 背包 DP

## 0-1 背包

## 分组背包

### CF148E: Porcelain

???+ note "问题描述"
    第一行输入 $n([1,100])$ 和 $m([1,1e4])$ ，接下来 $n$ 行输入 $n$ 个双端队列（dq）。<br>
    对于每个 dq，先输入 $k([1,100])$ 表示 dq 的大小，然后输入 dq 中的 $k$ 个数，范围在 $[1,100]$。所有 $k$ 之和 $≥m$。<br>
    你需要从这 $n$ 个 dq 中取出 $m$ 个数，输出这 $m$ 个数的和的最大值。

    在 [Codeforces](https://codeforces.com/problemset/problem/148/E "1900")
    或 [洛谷](https://www.luogu.com.cn/problem/CF148E "提高+/省选-") 查看该题。

??? info "代码实现"
    === "Go"
        ```go
        --8<-- "dp/knapsack/group/go/cf148e.go"
        ```
    === "Java"
        ```java
        --8<-- "dp/knapsack/group/java/cf148e/Main.java"
        ```
