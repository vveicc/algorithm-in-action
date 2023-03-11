# 树形 DP

## 换根 DP

### CF219D: Choosing Capital for Treeland

???+ note "问题描述"
    输入 $n(2≤n≤2e5)$ 和 $n-1$ 条边 $v \quad w$，表示一条 $v \to w$ 的有向边（节点编号从 $1$ 开始）。<br>
    保证输入构成一棵树。<br>
    你可以把边反向，即 $v \to w$ 改成 $w \to v$。<br>
    定义 $f(x)$ 表示以 $x$ 为根时，要让 $x$ 能够到达任意点，需要反向的边的数量。<br>
    第一行输出 $min(f(x))$，第二行升序输出所有等于 $min(f(x))$ 的节点编号。

    在 [Codeforces](https://codeforces.com/problemset/problem/219/D) 或 [洛谷](https://www.luogu.com.cn/problem/CF219D) 查看该题。

??? info "代码实现"
    === "Go"
        ```go
        --8<-- "dp/tree/root-changing/go/cf219d.go"
        ```
    === "Java"
        ```java
        --8<-- "dp/tree/root-changing/java/cf219d/Main.java"
        ```
