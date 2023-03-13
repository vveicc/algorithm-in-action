# 动态规划基础

## 二维 DP

### 最小总代价

#### CF1525D: Armchairs

???+ note "问题描述"
    输入 $n(2≤n≤5000)$ 和长为 $n$ 的数组 $a$ ，其中只有 0 和 1。保证 1 的数量不超过 $n/2$ 。<br>
    $a[i]=0$ 表示位置 $i$ 处有一把椅子，$a[i]=1$ 表示位置 $i$ 处有一个人。<br>
    一把椅子只能坐一个人。一个人从 $i$ 移动到 $j$ 的代价为 $abs(i-j)$ 。<br>
    问所有人都坐到椅子上的总代价和最小是多少？

    在 [Codeforces](https://codeforces.com/problemset/problem/1525/D "1800")
    或 [洛谷](https://www.luogu.com.cn/problem/CF1525D "普及+/提高") 上查看该题。

    进阶：如果一把椅子可以坐多个人呢？参考 [LC2463: 最小移动总距离](#lc2463-最小移动总距离)。

??? info "解题思路"
    根据输入可以得到人的位置序列 $peoples$ 和椅子的位置序列 $chairs$ ，两个序列严格递增。<br>
    假设最优方案中，第 $i$ 个人坐到椅子 $c_i$ 上，则 **存在最优方案，使得 $c_i$ 是严格单调递增的** 。<br>
    出现交叉的方案不会更优。

    定义 $dp[i][j]$ 表示前 $j$ 个人在前 $i$ 把椅子中找到椅子坐的最小总代价。
    
    1. 如果第 $j$ 个人不坐在第 $i$ 把椅子，则 $dp[i][j] = dp[i-1][j]$ ；
    2. 如果第 $j$ 个人坐在第 $i$ 把椅子上，则 $dp[i][j] = dp[i-1][j] + abs(chairs[i-1] - peoples[j-1])$ 。
    
    实现时可以优化掉第一个维度。

    === "Go"
        ```go
        --8<-- "basic/two-dimensional/min-cost/go/cf1525d.go"
        ```
    === "Java"
        ```java
        --8<-- "basic/two-dimensional/min-cost/java/cf1525d/Main.java"
        ```

#### LC2463: 最小移动总距离

???+ note "问题描述"
    长度为 $m(1≤m≤100)$ 的机器人数组 $robot$ 给出第 $i$ 个机器人的位置 $robot[i]\ ([-1e9,1e9])$ 。<br>
    长度为 $n(1≤n≤100)$ 的工厂数组 $factory$ 给出第 $j$ 个工厂的：
    
    1. 位置：$factory[j][0]\ ([-1e9,1e9])$ ；
    2. 维修机器人的上限：$factory[j][1]\ ([0,m])$ 。
    
    每个机器人所在的位置互不相同。每个工厂所在的位置也互不相同。<br>
    注意一个机器人可能一开始跟一个工厂在相同的位置。<br>
    机器人从位置 $x$ 到位置 $y$ 的移动距离为 $|x-y|$ 。<br>
    问所有机器人都进入工厂维修的移动距离总和最小是多少？测试数据保证所有机器人都可以被维修。<br>

    在 [LeetCode主站](https://leetcode.com/problems/minimum-total-distance-traveled "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/minimum-total-distance-traveled "困难：2454") 上查看该题。

    如果觉得这题比较难，可以先做简单版：[CF1525D: Armchairs](#cf1525d-armchairs)。

??? info "解题思路"
    不失一般性地，假设机器人的坐标是递增的，工厂的坐标也是递增的。<br>
    设最优方案中，机器人 $i$ 进入工厂 $f_i$，则 **存在最优方案，使得 $f_i$ 是不严格单调递增的** 。<br>
    与 [CF1525D: Armchairs](#cf1525d-armchairs) 类似，唯一的不同在于多个机器人可以进入相同的工厂，所以 $f_i$ 是不严格单调递增的。<br>
    定义 $dp[i][j]$ 表示前 $j$ 个机器人进入前 $i$ 个工厂维修的最小移动总距离，实现时优化掉第一个维度。

    === "Go"
        ```go
        --8<-- "basic/two-dimensional/min-cost/go/lc2463.go"
        ```

#### LC1478: 安排邮筒

???+ note "问题描述"
    给你一个长度为 $n(1≤n≤100)$ 的房屋数组 $houses$ 和一个整数 $k(1≤k≤n)$ 。<br>
    其中 $houses[i]\ ([1,1e4])$ 是第 $i$ 栋房子在一条街上的位置，现需要在这条街上安排 $k$ 个邮筒。<br>
    请你返回每栋房子与离它最近的邮筒之间的距离的最小总和。

    在 [LeetCode主站](https://leetcode.com/problems/allocate-mailboxes "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/allocate-mailboxes "困难：2190") 上查看该题。

??? info "解题思路"
    给一组房屋安排一个邮筒，选择房屋的中位数作为目标点，距离总和最小。

    === "Go"
        ```go
        --8<-- "basic/two-dimensional/min-cost/go/lc1478.go"
        ```
