# 树形 DP

## 换根 DP

### CF219D. Choosing Capital for Treeland { data-toc-label='CF219D. Choosing Capital fo...' }

???+ note "问题描述"
    输入 $n(2≤n≤2e5)$ 和 $n-1$ 条边 $v \quad w$，表示一条 $v \to w$ 的有向边（节点编号从 $1$ 开始）。<br>
    保证输入构成一棵树。<br>
    你可以把边反向，即 $v \to w$ 改成 $w \to v$。<br>
    定义 $f(x)$ 表示以 $x$ 为根时，要让 $x$ 能够到达任意点，需要反向的边的数量。<br>
    第一行输出 $min(f(x))$，第二行升序输出所有等于 $min(f(x))$ 的节点编号。

    在 [Codeforces](https://codeforces.com/problemset/problem/219/D "1700")
    或 [洛谷](https://www.luogu.com.cn/problem/CF219D "普及+/提高") 查看该题。

??? info "解题思路"
    先通过 DFS 计算出以 $1$ 为根的反向边数量。<br>
    然后进行换根 DP，假设以 $x$ 为根的反向边数量为 $n$ ，考虑与 $x$ 相连的节点 $y$ ：
    
    1. 如果 $x \to y$ ，则以 $y$ 为根的反向边数量为 $n+1$ ；
    2. 如果 $x \gets y$ ，则以 $y$ 为根的反向边数量为 $n-1$ 。

    === "Go"
        ```go
        --8<-- "tree/root-changing/go/cf219d.go"
        ```
    === "Java"
        ```java
        --8<-- "tree/root-changing/java/cf219d/Main.java"
        ```

### CF337D. Book of Evil

???+ note "问题描述"
    输入 $n\ m(1≤m≤n≤1e5)\ d(0≤d≤n-1)$ 表示一棵 $n$ 个节点的树，其中 $m$ 个节点有怪物，这些怪物是由一个传送门生成的，传送门与任意怪物的距离不超过 $d$。<br>
    然后输入 $m$ 个互不相同的数，表示怪物所在节点编号（从 $1$ 开始）。<br>
    然后输入 $n-1$ 行，每行两个节点编号，表示树的边。<br>
    输出可能存在传送门的节点的个数。注意传送门只有一个。

    在 [Codeforces](https://codeforces.com/problemset/problem/337/D "2000")
    或 [洛谷](https://www.luogu.com.cn/problem/CF337D "提高+/省选-") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "tree/root-changing/go/cf337d.go"
        ```
    === "Java"
        ```java
        --8<-- "tree/root-changing/java/cf337d/Main.java"
        ```
