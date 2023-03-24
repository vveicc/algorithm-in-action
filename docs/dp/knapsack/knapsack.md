# 单背包问题

## 0-1 背包

### [[T]](#0-1-背包 "模板") AcWing2: 0-1 背包

???+ note "问题描述"
    第一行输入 $n$ 和 $m$ 两个整数，$1≤n,m≤1e3$ 。表示有 $n$ 件物品和一个容量为 $m$ 的背包。<br>
    接下来 $n$ 行，每行输入 $v_i$ 和 $w_i$ 两个整数，表示第 $i$ 件物品的体积是 $v_i$ ，价值是 $w_i$ ，$1≤v_i,w_i≤1e3$ 。<br>
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

### [[T]](#完全背包 "模板") AcWing3: 完全背包

???+ note "问题描述"
    第一行输入 $n$ 和 $m$ 两个整数，$1≤n,m≤1e3$ 。表示有 $n$ 种物品和一个容量为 $m$ 的背包。<br>
    接下来 $n$ 行，每行输入 $v_i$ 和 $w_i$ 两个整数，表示第 $i$ 种物品的体积是 $v_i$ ，价值是 $w_i$ ，$1≤v_i,w_i≤1e3$ 。<br>
    每种物品都有无限件可供选择。<br>
    求解将哪些物品装入背包，可使这些物品的总体积不超过背包容量，且总价值最大。输出最大总价值。

    在 [AcWing](https://www.acwing.com/problem/content/3 "简单") 查看该题。

??? info "解题思路"
    定义 $f[i][j]$ 表示只从前 $i$ 种物品中选择的情况下，容量为 $j$ 的背包所能达到的最大总价值。<br>
    根据第 $i$ 件物品放入 $k$ 件，状态转移如下：
    
    $$f[i][j] = {max}\{f[i-1][j-k \times v_i] + k \times w_i\} \qquad 0≤k \times v_i≤m$$

    当 $k = 0$ 时，$f[i][j] = f[i-1][j]$ ；当 $k > 1$ 时，有

    $$
    \begin{align*}
        f[i][j] &= {max}\{f[i-1][j-v_i-(k-1) \times v_i] + (k-1) \times w_i + w_i\} \\
                \\
                &= f[i][j-v_i] + w_i
    \end{align*}
    $$

    因此，状态转移方程可以优化为：

    $$f[i][j] = {max}(f[i-1][j], f[i][j-v_i])$$

    实现时可以优化掉第一个维度，注意循环枚举顺序。

    === "Go"
        ```go
        --8<-- "knapsack/complete/go/acwing3.go"
        ```
    === "Java"
        ```java
        --8<-- "knapsack/complete/java/acwing3/Main.java"
        ```

## 多重背包

### [[T]](#多重背包 "模板") AcWing4: 多重背包 I

???+ note "问题描述"
    第一行输入 $n$ 和 $m$ 两个整数，$1≤n,m≤100$ 。表示有 $n$ 种物品和一个容量为 $m$ 的背包。<br>
    接下来 $n$ 行，每行输入 $v_i\ \ w_i\ \ c_i$ 三个整数，表示第 i 种物品的体积是 $v_i$ ，价值是 $w_i$ ，有 $c_i$ 件，$1≤v_i,w_i,c_i≤100$ 。<br>
    求解将哪些物品装入背包，可使这些物品的总体积不超过背包容量，且总价值最大。输出最大总价值。

    在 [AcWing](https://www.acwing.com/problem/content/4 "简单") 查看该题。

??? info "解题思路"
    由于每种物品的数量有限，不能按照完全背包进行优化，原因是状态转移方程不一定成立。

    ??? tip "转化为 0-1 背包问题"
        **每种物品选 $c_i$ 次** 等价于 **有 $c_i$ 件相同物品，每件物品最多选一次** 。<br>
        问题转换为 0-1 背包问题，时间复杂度：$O(m\sum{c_i})$ 。

        === "Go"
            ```go
            --8<-- "knapsack/multi-items/go/acwing4_1.go"
            ```
        === "Java"
            ```java
            --8<-- "knapsack/multi-items/java/acwing4_1/Main.java"
            ```
    
    ??? tip "转化为分组背包问题"
        把每种物品看成一个分组，分组内包含分别由 $1, 2, ..., c_i$ 件该种物品组成的 $c_i$ 件物品，最多选择一件。<br>
        问题转换为[分组背包问题](#分组背包)，时间复杂度：$O(m\sum{c_i})$ 。

        === "Go"
            ```go
            --8<-- "knapsack/multi-items/go/acwing4_2.go"
            ```
        === "Java"
            ```java
            --8<-- "knapsack/multi-items/java/acwing4_2/Main.java"
            ```
    
    ??? tip "进阶一：二进制分组优化"
        [AcWing5: 多重背包问题 II](https://www.acwing.com/problem/content/5 "中等") 将数据范围增大为：$1≤n≤1e3$ ；$1≤m≤2e3$ ；$1≤v_i,w_i,c_i≤2e3$ 。<br>
        采用上述方法将会超出时间限制，可以通过二进制分组进行优化，时间复杂度：$O(m\sum{\log{c_i}})$ 。

        === "Go"
            ```go
            --8<-- "knapsack/multi-items/go/acwing5.go"
            ```
        === "Java"
            ```java
            --8<-- "knapsack/multi-items/java/acwing5/Main.java"
            ```
    
    ??? tip "进阶二：单调队列优化"
        [AcWing6: 多重背包问题 III](https://www.acwing.com/problem/content/6 "困难") 将数据范围增大为：$1≤n≤1e3$ ；$1≤m≤2e4$ ；$1≤v_i,w_i,c_i≤2e4$ 。<br>
        采用二进制分组优化也会超出时间限制，需要通过单调队列进行优化，时间复杂度：$O(nm)$ 。

        === "Go"
            ```go
            --8<-- "knapsack/multi-items/go/acwing6.go"
            ```
        === "Java"
            ```java
            --8<-- "knapsack/multi-items/java/acwing6/Main.java"
            ```


## 混合背包

### [[T]](#混合背包 "模板") AcWing7: 混合背包

???+ note "问题描述"
    第一行输入 $n$ 和 $m$ 两个整数，$1≤n,m≤1e3$ 。表示有 $n$ 种物品和一个容量为 $m$ 的背包。<br>
    接下来 $n$ 行，每行输入 $v_i\ \ w_i\ \ c_i$ 三个整数，表示第 $i$ 种物品的体积是 $v_i$ ，价值是 $w_i$ ，$1≤v_i,w_i≤1e3$ ，$-1≤c_i≤1e3$ 的值则分以下三种情况：<br>

    1. $c_i = -1$ 表示第 i 种物品只有 $1$ 件；
    2. $c_i = 0$ 表示第 i 种物品有无限件；
    3. $c_i > 0$ 表示第 i 种物品有 $c_i$ 件。

    求解将哪些物品装入背包，可使这些物品的总体积不超过背包容量，且总价值最大。输出最大总价值。

    在 [AcWing](https://www.acwing.com/problem/content/7 "中等") 查看该题。

??? info "解题思路"
    混合背包就是将[0-1 背包](#0-1-背包)、[完全背包](#完全背包)、[多重背包](#多重背包)三种背包混合起来，求解时分情况处理即可。

    === "Go"
        ```go
        --8<-- "knapsack/mixed/go/acwing7.go"
        ```
    === "Java"
        ```java
        --8<-- "knapsack/mixed/java/acwing7/Main.java"
        ```

## 二维费用背包

### [[T]](#二维费用背包 "模板") AcWing8: 二维费用背包

???+ note "问题描述"
    第一行输入 $N\ \ C\ \ M$ 三个整数，$1≤N≤1e3$ ，$1≤C,M≤1e2$ 。表示有 $N$ 件物品和一个容量为 $C$ 最大承重为 $M$ 的背包。<br>
    接下来 $N$ 行，每行输入 $v_i\ \ m_i\ \ w_i$ 三个整数，表示第 $i$ 件物品的体积是 $v_i$ ，重量是 $m_i$ ，价值是 $w_i$ ，$1≤v_i,m_i≤1e2$ ，$1≤w_i≤1e3$ 。<br>
    求解将哪些物品装入背包，可使这些物品的总体积不超过背包容量，总重量不超过背包最大承重，且总价值最大。输出最大总价值。

    在 [AcWing](https://www.acwing.com/problem/content/8 "中等") 查看该题。

??? info "解题思路"
    在 0-1 背包的基础上增加一个费用维度即可。

    === "Go"
        ```go
        --8<-- "knapsack/two-cost/go/acwing8.go"
        ```
    === "Java"
        ```java
        --8<-- "knapsack/two-cost/java/acwing8/Main.java"
        ```

## 分组背包

### [[T]](#分组背包 "模板") AcWing9: 分组背包

???+ note "问题描述"
    第一行输入 $n$ 和 $m$ 两个整数，$1≤n,m≤1e2$ 。表示有 $N$ 组物品和一个容量为 $m$ 的背包。<br>
    接下来输入 $n$ 组物品的数据：

    * 每组数据第一行输入一个整数 $c_i(1≤c_i≤1e2)$ ，表示第 $i$ 组有 $c_i$ 件物品；
    * 接下来 $c_i$ 行，每行输入 $v_{ij}$ 和 $w_{ij}$ 两个整数，表示第 $i$ 组第 $j$ 件物品的体积是 $v_{ij}$ ，价值是 $w_{ij}$ ，$1≤v_{ij},w_{ij}≤1e2$ 。

    每组物品有若干件，同一组内的物品最多只能选一件。<br>
    求解将哪些物品装入背包，可使这些物品的总体积不超过背包容量，且总价值最大。输出最大总价值。

    在 [AcWing](https://www.acwing.com/problem/content/9 "中等") 查看该题。

??? info "解题思路"
    和 0-1 背包类似，对每一组进行一次 0-1 背包即可。

    === "Go"
        ```go
        --8<-- "knapsack/group/go/acwing9.go"
        ```
    === "Java"
        ```java
        --8<-- "knapsack/group/java/acwing9/Main.java"
        ```

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

### [[T]](#有依赖的背包 "模板") AcWing10: 有依赖的背包

???+ note "问题描述"
    第一行输入 $n$ 和 $m$ 两个整数，$1≤n,m≤1e2$ 。表示有编号为 $1\dots n$ 的 $n$ 件物品和一个容量为 $m$ 的背包。<br>
    物品之间存在依赖关系，所有物品的依赖关系形成一棵树。如果选择某件物品，则必须同时选择它的父节点。<br>
    接下来 $n$ 行，每行输入 $v_i\ \ w_i\ \ p_i$ 三个整数，表示编号为 $i$ 的物品的体积是 $v_i$ ，价值是 $w_i$ ，$1≤v_i,w_i≤1e2$ ，依赖编号为 $p_i$ 的物品，$p_i = -1$ 表示该物品为根节点。<br>
    求解将哪些物品装入背包，可使这些物品的总体积不超过背包容量，且总价值最大。输出最大总价值。

    在 [AcWing](https://www.acwing.com/problem/content/10 "困难") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "knapsack/dependent/go/acwing10.go"
        ```
    === "Java"
        ```java
        --8<-- "knapsack/dependent/java/acwing10/Main.java"
        ```

### LG1064: 金明的预算方案

???+ note "问题描述"
    第一行输入 $n(1≤n≤3.2e4)$ 和 $m(1≤m≤60)$ 两个整数。表示有 $n$ 元钱，想要买 $m$ 个物品。<br>
    接下来 $m$ 行，每行输入 $v_i\ \ p_i\ \ q_i$ 三个整数，表示第 $i(1≤i≤m)$ 个物品的价格为 $v_i(1≤v_i≤1e4)$ ，重要度为 $p_i(1≤p_i≤5)$ ，$v_i$ 是 $10$ 的整数倍。如果 $q_i = 0$ ，表示该物品是主件；否则，表示该物品是第 $q_i$ 个物品（主件）的附件。<br>
    主件可以单独购买，购买附件则必须同时购买它的主件，每个主件最多有 $2$ 个附件。<br>
    求解购买哪些物品，可使这些物品的 $v_i \times p_i$ 之和最大。输出最大和。题目保证答案 $≤2e5$ 。

    在 [洛谷](https://www.luogu.com.cn/problem/P1064 "普及+/提高") 查看该题。

??? info "解题思路"
    将主件与其附件的所有组合当做一个分组，转化为[分组背包](#分组背包)。
    
    === "Go"
        ```go
        --8<-- "knapsack/dependent/go/lg1064.go"
        ```
    === "Java"
        ```java
        --8<-- "knapsack/dependent/java/lg1064/Main.java"
        ```

## 动态物品背包

每个物品的费用或价值根据一定条件动态变化。状态转移时计算具体的费用或价值即可。

## 背包问题变种

### 求方案数

#### [[T]](#求方案数 "模板") AcWing11: 求方案数

???+ note "问题描述"
    第一行输入 $n$ 和 $m$ 两个整数，$1≤n,m≤1e3$ 。表示有 $n$ 件物品和一个容量为 $m$ 的背包。<br>
    接下来 $n$ 行，每行输入 $v_i$ 和 $w_i$ 两个整数，表示第 $i$ 件物品的体积是 $v_i$ ，价值是 $w_i$ ，$1≤v_i,w_i≤1e3$ 。<br>
    求解将哪些物品装入背包，可使这些物品的总体积不超过背包容量，且总价值最大。<br>
    输出 **最优选法的方案数**。注意答案可能很大，请输出答案模 $1e9+7$ 的结果。

    在 [AcWing](https://www.acwing.com/problem/content/11 "中等") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "knapsack/variant/go/acwing11.go"
        ```
    === "Java"
        ```java
        --8<-- "knapsack/variant/java/acwing11/Main.java"
        ```

### 求具体方案

#### [[T]](#求具体方案 "模板") AcWing12: 求具体方案

???+ note "问题描述"
    第一行输入 $n$ 和 $m$ 两个整数，$1≤n,m≤1e3$ 。表示有 $n$ 件物品和一个容量为 $m$ 的背包。<br>
    接下来 $n$ 行，每行输入 $v_i$ 和 $w_i$ 两个整数，表示第 $i$ 件物品的体积是 $v_i$ ，价值是 $w_i$ ，$1≤v_i,w_i≤1e3$ 。<br>
    求解将哪些物品装入背包，可使这些物品的总体积不超过背包容量，且总价值最大。<br>
    输出 **字典序最小的方案**。这里的字典序是指：所选物品的编号所构成的序列。物品的编号范围是 $1 \dots n$ 。

    在 [AcWing](https://www.acwing.com/problem/content/12 "中等") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "knapsack/variant/go/acwing12.go"
        ```
    === "Java"
        ```java
        --8<-- "knapsack/variant/java/acwing12/Main.java"
        ```
