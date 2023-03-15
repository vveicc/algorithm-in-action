# 背包 DP

## 0-1 背包

### [[T]](#t-acwing2-01背包问题 "模板") AcWing2: 01背包问题

???+ note "问题描述"
    第一行输入 $n$ 和 $m$ 两个整数，$1≤n,m≤1000$ 。表示有 $n$ 件物品和一个容量为 $m$ 的背包。<br>
    接下来 $n$ 行，每行输入 $v_i$ 和 $w_i$ 两个整数，表示第 i 件物品的体积是 $v_i$ ，价值是 $w_i$ 。<br>
    求解将哪些物品装入背包，可使这些物品的总体积不超过背包容量，且总价值最大。输出最大总价值。

    在 [AcWing](https://www.acwing.com/problem/content/2 "简单") 查看该题。

??? info "解题思路"
    定义 $f[i][j]$ 表示只从前 $i$ 个物品中选择的情况下，容量为 $j$ 的背包所能达到的最大总价值。<br>
    根据第 $i$ 件物品是否放入，状态转移如下：
    
    $$f[i][j] = {max}(f[i-1][j], f[i-1][j-v_i] + w_i)$$

    实现时可以优化掉第一个维度，注意循环枚举顺序。

    === "Go"
        ```go
        --8<-- "knapsack/01/go/acwing2.go"
        ```
    === "Java"
        ```java
        --8<-- "knapsack/01/java/acwing2/Main.java"
        ```

## 完全背包

## 多重背包

## 混合背包

## 二维费用背包

## 分组背包

### CF148E: Porcelain

???+ note "问题描述"
    第一行输入 $n(1≤n≤100)$ 和 $m(1≤m≤1e4)$ ，接下来 $n$ 行输入 $n$ 个双端队列（dq）。<br>
    对于每个 dq，先输入 $k(1≤k≤100)$ 表示 dq 的大小，所有 $k$ 之和 $≥m$。<br>
    然后输入 dq 中的 $k$ 个数，范围在 $[1,100]$。<br>
    你需要从这 $n$ 个 dq 中取出 $m$ 个数，输出这 $m$ 个数的和的最大值。

    在 [Codeforces](https://codeforces.com/problemset/problem/148/E "1900")
    或 [洛谷](https://www.luogu.com.cn/problem/CF148E "提高+/省选-") 查看该题。

??? info "解题思路"
    定义 $groups[i][j]$ 表示从第 $i$ 个 dq 中取出 $j+1$ 个数的最大和。<br>
    预处理每一个 dq ，枚举在队头取出的数字个数，通过前缀和计算得到 $groups[i][j]$ 。<br>
    预处理得到 $groups[i][j]$ 后，问题转化为背包容量为 $m$ 的分组背包问题。

    === "Go"
        ```go
        --8<-- "knapsack/group/go/cf148e.go"
        ```
    === "Java"
        ```java
        --8<-- "knapsack/group/java/cf148e/Main.java"
        ```

## 有依赖的背包

## 动态物品背包

## 多背包

以上都是单个背包的情况，这里整理一些涉及到多个背包的问题。

### 静态物品多背包

每个物品的费用和价值固定。

#### CF687C: The Values You Can Make

???+ note "问题描述"
    第一行输入 $n$ 和 $k$ 两个整数，$1≤n,k≤500$ 。<br>
    第二行输入长度为 $n$ 的数组 $c(1≤c[i]≤500)$ 。<br>
    从 $c$ 中选择若干个数，组成数组 $a$ ，满足 $sum(a) = k$ 。<br>
    从 $a$ 中再选择若干个数，组成数组 $b$（可以为空）。<br>
    计算 $sum(b)$ 的所有可能的值。第一行输出这些值的个数 $q$ ，第二行按升序输出这 $q$ 个数。

    在 [Codeforces](https://codeforces.com/problemset/problem/687/C "1900")
    或 [洛谷](https://www.luogu.com.cn/problem/CF687C "普及+/提高") 查看该题。

??? info "解题思路"
    如果能从 $c$ 中选择两个不相交的子集 $A$ 和 $B$，使得 $\sum A = x$ 且 $\sum B = k-x$ 。则 $x$ 就是一个可能值。<br>
    问题转化为 **有两个背包的 0-1 背包问题** ，每个物品 $c_i$ ，要么不选，要么放入背包1，要么放入背包2。<br>

    定义 $dp[i][j_1][j_2]$ 表示从前 $i$ 个数中能否选出两个不相交的子集 $A$ 和 $B$，使得 $\sum A = j_1$ 且 $\sum B = j_2$ 。<br>
    则有 $dp[i][j_1][j_2] = dp[i-1][j_1][j_2] \parallel dp[i-1][j_1-c_i][j_2] \parallel dp[i-1][j_1][j_2-c_i]$ 。<br>
    所有使得 $dp[n][x][k-x] = true$ 的 $x$ 都是可能值。<br>
    实现时优化掉第一个维度，注意循环枚举顺序。

    === "Go"
        ```go
        --8<-- "knapsack/multi-knapsacks/go/cf687c.go"
        ```
    === "Java"
        ```java
        --8<-- "knapsack/multi-knapsacks/java/cf687c/Main.java"
        ```

### 动态物品多背包

每个物品的费用或价值根据一定条件动态变化。

## 背包问题变种
