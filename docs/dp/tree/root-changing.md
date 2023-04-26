# 换根 DP

## CF219D. Choosing Capital for Treeland { data-toc-label='CF219D. Choosing Capital for...' }

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

## CF337D. Book of Evil

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

## LC2538. 最大开销

???+ note "问题描述"
    给你一个整数 `1≤n≤1e5` 表示有 n 个节点的无向无根树，节点编号为 `0` 到 `n-1` 。<br>
    给你一个长度为 `n-1` 的二维整数数组 `edges` ，其中 $edges[i] = [a_i, b_i]$ 表示树中节点 $a_i$ 和 $b_i$ 之间有一条边。<br>
    每个节点都有一个价值。给你一个长度为 `n` 的整数数组 `price` ，其中 `1≤price[i]≤1e5` 是第 `i` 个节点的价值。

    一条路径的 **价值和** 是这条路径上所有节点的价值之和。<br>
    你可以选择树中任意一个节点作为根节点 `root` 。选择 `root` 为根的 **开销** 是以 `root` 为起点的所有路径中，**价值和** 最大的一条路径与最小的一条路径的差值。<br>
    请你返回所有节点作为根节点的选择中，**最大** 的 **开销** 为多少。

    在 [LeetCode主站](https://leetcode.com/problems/difference-between-maximum-and-minimum-price-sum "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/difference-between-maximum-and-minimum-price-sum "困难：2398") 查看该题。

??? info "解题思路"
    **方法一：[树形 DP](./index.md#lc2538-最大开销)**

    **方法二：换根 DP**

    === "Go"
        ```go
        --8<-- "tree/root-changing/go/lc2538.go"
        ```

## LC2603. 收集树中金币

???+ note "问题描述"
    有一个 `1≤n≤3e4` 个节点的无向无根树，节点编号从 `0` 到 `n - 1` 。<br>
    给你一个长度为 `n` 的数组 `coins` ，其中 `coins[i]` 等于 `0` 或 `1` ，`1` 表示节点 `i` 处有一个金币。<br>
    给你一个长度为 `n - 1` 的二维整数数组 `edges` ，其中 `edges[i]` = [$a_i$, $b_i$] 表示树中节点 $a_i$ 和 $b_i$ 之间有一条边。<br>
    一开始，你需要选择树中任意一个节点出发。你可以执行下述操作任意次：

    - 收集与当前节点距离 `≤2` 的所有金币，或者
    - 移动到树中一个相邻节点。

    你需要收集树中所有的金币，并回到出发节点，请你返回最少经过的边数。如果你多次经过一条边，每一次经过都会给答案加一。

    在 [LeetCode主站](https://leetcode.com/problems/collect-coins-in-a-tree "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/collect-coins-in-a-tree "困难：2712") 查看该题。

??? info "解题思路"
    **方法一：[拓扑排序](../../graph/topo.md#lc2603-收集树中金币)**

    **方法二：换根 DP**

    换根 DP 的思路与 [CF337D. Book of Evil](#cf337d-book-of-evil) 类似。

    === "Go"
        ```go
        --8<-- "tree/root-changing/go/lc2603.go"
        ```
