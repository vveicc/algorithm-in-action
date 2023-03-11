# 二分

## 二分法

### CF1379C: Choosing flowers

???+ note "问题描述"
    第一行输入 $t(1≤t≤1e4)$ 表示 $t$ 组由空行分隔的数据。所有 $t$ 组数据的 $m$ 之和 $≤1e5$。<br>
    对于每组数据：<br>
    第一行输入 $n(1≤n≤1e9)$ 和 $m(1≤m≤1e5)$ 表示有 $m$ 种物品，每种物品有无限个，你需要选择 $n$ 个。<br>
    然后输入 $m$ 行，每行两个数字 $a[i]$ 和 $b[i]$，范围在 $[0,1e9]$。<br>
    如果第 $i$ 种物品选 $x(x>0)$ 个，收益为 $a[i]+(x-1)*b[i]$。<br>
    输出最大收益。
    
    在 [Codeforces](https://codeforces.com/problemset/problem/1379/C "2000")
    或 [洛谷](https://www.luogu.com.cn/problem/CF1379C "提高+/省选-") 上查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "binary/binary/go/cf1379c.go"
        ```
    === "Java"
        ```java
        --8<-- "binary/binary/java/cf1379c/Main.java"
        ```

## 三分法
