# 背包 DP

## 0-1 背包

### [[T]](#t-acwing2-0-1-背包问题 "模板") AcWing2: 0-1 背包问题

???+ note "问题描述"
    第一行输入 $n$ 和 $m$ 两个整数，$1≤n,m≤1e3$ 。表示有 $n$ 件物品和一个容量为 $m$ 的背包。<br>
    接下来 $n$ 行，每行输入 $v_i$ 和 $w_i$ 两个整数，表示第 i 件物品的体积是 $v_i$ ，价值是 $w_i$ ，$1≤v_i,w_i≤1e3$ 。<br>
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

### [[T]](#t-acwing3-完全背包问题 "模板") AcWing3: 完全背包问题

???+ note "问题描述"
    第一行输入 $n$ 和 $m$ 两个整数，$1≤n,m≤1e3$ 。表示有 $n$ 种物品和一个容量为 $m$ 的背包。<br>
    接下来 $n$ 行，每行输入 $v_i$ 和 $w_i$ 两个整数，表示第 i 种物品的体积是 $v_i$ ，价值是 $w_i$ ，$1≤v_i,w_i≤1e3$ 。<br>
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

### [[T]](#t-acwing4-多重背包问题-i "模板") AcWing4: 多重背包问题 I

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

#### CF1525D: Armchairs

???+ note "问题描述"
    输入 $n(2≤n≤5000)$ 和长为 $n$ 的数组 $a$ ，其中只有 0 和 1。保证 1 的数量不超过 $n/2$ 。<br>
    $a[i]=0$ 表示位置 $i$ 处有一把椅子，$a[i]=1$ 表示位置 $i$ 处有一个人。<br>
    一把椅子只能坐一个人。一个人从 $i$ 移动到 $j$ 的代价为 $abs(i-j)$ 。<br>
    问所有人都坐到椅子上的总代价和最小是多少？

    在 [Codeforces](https://codeforces.com/problemset/problem/1525/D "1800")
    或 [洛谷](https://www.luogu.com.cn/problem/CF1525D "普及+/提高") 查看该题。

    进阶：如果一把椅子可以坐多个人呢？参考 [LC2463: 最小移动总距离](#lc2463-最小移动总距离)。

??? info "解题思路"
    根据输入可以得到人的位置序列 $peoples$ 和椅子的位置序列 $chairs$ ，两个序列严格递增。<br>
    假设最优方案中，第 $i$ 个人坐到椅子 $c_i$ 上，则 **存在最优方案，使得 $c_i$ 是严格单调递增的** 。<br>
    出现交叉的方案不会更优。

    如果把椅子当做容量为 $1$ 的背包，人当做体积为 $1$ 的物品，则转化为动态物品多背包问题。<br>
    定义 $dp[i][j]$ 表示前 $j$ 个人在前 $i$ 把椅子中找到椅子坐的最小总代价。
    
    1. 如果第 $j$ 个人不坐在第 $i$ 把椅子，则 $dp[i][j] = dp[i-1][j]$ ；
    2. 如果第 $j$ 个人坐在第 $i$ 把椅子上，则 $dp[i][j] = dp[i-1][j-1] + abs(chairs[i-1] - peoples[j-1])$ 。
    
    实现时可以优化掉第一个维度，注意循环枚举顺序。

    === "Go"
        ```go
        --8<-- "knapsack/multi-knapsacks/go/cf1525d.go"
        ```
    === "Java"
        ```java
        --8<-- "knapsack/multi-knapsacks/java/cf1525d/Main.java"
        ```

#### LC2463: 最小移动总距离

???+ note "问题描述"
    给你一个长度为 $m(1≤m≤100)$ 的机器人数组 $robot$ ，$robot[i]$ 是第 $i$ 个机器人的位置，范围：$[-1e9,1e9]$。<br>
    给你一个长度为 $n(1≤n≤100)$ 的工厂数组 $factory$ ，$factory[j] = [pos_j, limit_j]$ 表示第 `j` 个工厂的位置和维修机器人的上限，其中 $-1e9≤pos_j≤1e9$ ，$0≤limit_j≤m$ 。<br>
    每个机器人所在的位置互不相同。每个工厂所在的位置也互不相同。<br>
    注意一个机器人可能一开始跟一个工厂在相同的位置。<br>
    机器人从位置 $x$ 到位置 $y$ 的移动距离为 $|x-y|$ 。<br>
    问所有机器人都进入工厂维修的移动距离总和最小是多少？测试数据保证所有机器人都可以被维修。<br>

    在 [LeetCode主站](https://leetcode.com/problems/minimum-total-distance-traveled "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/minimum-total-distance-traveled "困难：2454") 查看该题。

    如果觉得这题比较难，可以先做简单版：[CF1525D: Armchairs](#cf1525d-armchairs)。

??? info "解题思路"
    不失一般性地，假设机器人的坐标是递增的，工厂的坐标也是递增的。<br>
    设最优方案中，机器人 $i$ 进入工厂 $f_i$，则 **存在最优方案，使得 $f_i$ 是不严格单调递增的** 。<br>
    与 [CF1525D: Armchairs](#cf1525d-armchairs) 类似，唯一的不同在于多个机器人可以进入相同的工厂，所以 $f_i$ 是不严格单调递增的。<br>
    如果把每个工厂当做容量为 $factory[j][1]$ 的背包，机器人当做体积为 $1$ 的物品，则转化为动态物品多背包问题。<br>
    定义 $dp[i][j]$ 表示前 $j$ 个机器人进入前 $i$ 个工厂维修的最小移动总距离。<br>
    实现时可以优化掉第一个维度，注意循环枚举顺序。

    === "Go"
        ```go
        --8<-- "knapsack/multi-knapsacks/go/lc2463.go"
        ```

#### LC1478: 安排邮筒

???+ note "问题描述"
    给你一个长度为 $n(1≤n≤100)$ 的房屋数组 $houses$ 和一个整数 $k(1≤k≤n)$ 。<br>
    其中 $houses[i]$ 是第 $i$ 栋房子在一条街上的位置，范围：$[1,1e4]$ 。<br>
    现需要在这条街上安排 $k$ 个邮筒。请你返回每栋房子与离它最近的邮筒之间的距离的最小总和。

    在 [LeetCode主站](https://leetcode.com/problems/allocate-mailboxes "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/allocate-mailboxes "困难：2190") 查看该题。

??? info "解题思路"
    给一组房屋安排一个邮筒，选择房屋的中位数作为目标点，距离总和最小。<br>
    如果把每个邮筒当做容量不限的背包，房子当做物品，则转化为动态物品多背包问题。<br>
    定义 $dp[i][j]$ 表示给 $houses[:j+1]$ 安排编号为 $[0..i]$ 的 $i+1$ 个邮筒的最小距离总和。<br>
    实现时可以优化掉第一个维度，注意循环枚举顺序。

    === "Go"
        ```go
        --8<-- "knapsack/multi-knapsacks/go/lc1478.go"
        ```

## 背包问题变种
