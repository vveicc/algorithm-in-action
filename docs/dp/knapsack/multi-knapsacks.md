# 多背包问题

这里整理一些涉及到多个背包的问题。

## 静态物品多背包

每个物品的费用和价值固定。

### CF687C: The Values You Can Make

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

## 动态物品多背包

每个物品的费用或价值根据一定条件动态变化。状态转移时计算具体的费用或价值即可。

### CF1525D: Armchairs

???+ note "问题描述"
    第一行输入 $n(2≤n≤5000)$ ，第二行输入长为 $n$ 的数组 $a$ ，其中只有 0 和 1。保证 1 的数量不超过 $n/2$ 。<br>
    $a[i]=0$ 表示位置 $i$ 处有一把椅子，$a[i]=1$ 表示位置 $i$ 处有一个人。<br>
    一把椅子只能坐一个人。一个人从 $i$ 移动到 $j$ 的代价为 $abs(i-j)$ 。<br>
    问所有人都坐到椅子上的总代价和最小是多少？

    在 [Codeforces](https://codeforces.com/problemset/problem/1525/D "1800")
    或 [洛谷](https://www.luogu.com.cn/problem/CF1525D "普及+/提高") 查看该题。

    进阶：如果一把椅子可以坐多个人呢？参考 [LC2463: 最小移动总距离](#lc2463-最小移动总距离)。

??? info "解题思路"
    根据输入可以得到人的位置序列 $peoples$ 和椅子的位置序列 $chairs$ ，两个序列严格递增。<br>
    假设最优方案中，第 $i$ 个人坐到椅子 $c_i$ 上，则 **存在最优方案，使得 $c_i$ 是严格单调递增的** 。出现交叉的方案不会更优。

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

### LC2463: 最小移动总距离

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
    与 [CF1525D: Armchairs](#cf1525d-armchairs) 类似，唯一的不同在于多个机器人可以进入相同的工厂，所以 $f_i$ 是不严格单调递增的。

    如果把每个工厂当做容量为 $factory[j][1]$ 的背包，机器人当做体积为 $1$ 的物品，则转化为动态物品多背包问题。<br>
    定义 $dp[i][j]$ 表示前 $j$ 个机器人进入前 $i$ 个工厂维修的最小移动总距离。<br>
    实现时可以优化掉第一个维度，注意循环枚举顺序。

    === "Go"
        ```go
        --8<-- "knapsack/multi-knapsacks/go/lc2463.go"
        ```

### LC1478: 安排邮筒

???+ note "问题描述"
    给你一个长度为 $n(1≤n≤100)$ 的房屋数组 $houses$ 和一个整数 $k(1≤k≤n)$ 。<br>
    其中 $houses[i]$ 是第 $i$ 栋房子在一条街上的位置，范围：$[1,1e4]$ 。<br>
    现需要在这条街上安排 $k$ 个邮筒。请你返回每栋房子与离它最近的邮筒之间的距离的最小总和。

    在 [LeetCode主站](https://leetcode.com/problems/allocate-mailboxes "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/allocate-mailboxes "困难：2190") 查看该题。

??? info "解题思路"
    给一组房屋安排一个邮筒，选择房屋的中位数作为目标点，距离总和最小。

    如果把每个邮筒当做容量不限的背包，房子当做物品，则转化为动态物品多背包问题。<br>
    定义 $dp[i][j]$ 表示给 $houses[:j+1]$ 安排编号为 $0\dots i$ 的 $i+1$ 个邮筒的最小距离总和。<br>
    实现时可以优化掉第一个维度，注意循环枚举顺序。

    === "Go"
        ```go
        --8<-- "knapsack/multi-knapsacks/go/lc1478.go"
        ```
