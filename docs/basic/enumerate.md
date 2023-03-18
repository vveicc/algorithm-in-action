# 枚举

## LC1615: 最大网络秩

???+ note "问题描述"
    给你 `n(2≤n≤100)` 座城市，编号从 `0` 开始。<br>
    给你一个数组 `roads` ，其中 $roads[i] = [a_i, b_i]$ 都表示在城市 $a_i$ 和 $b_i$ 之间有一条双向道路。<br>
    题目保证每对城市之间最多只有一条道路相连。<br>
    任意两座不同城市构成的城市对的网络秩定义为：与这两座城市直接相连的道路总数。<br>
    如果存在一条道路直接连接这两座城市，则这条道路只计算一次。<br>
    请返回所有不同城市对中的最大网络秩。

    在 [LeetCode主站](https://leetcode.com/problems/maximal-network-rank "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/maximal-network-rank "中等：1522") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "enumerate/go/lc1615.go"
        ```

## LC829: 连续整数求和

???+ note "问题描述"
    给定一个正整数 `n(1≤n≤1e9)` ，返回连续正整数满足所有数字之和为 `n` 的组数 。

    在 [LeetCode主站](https://leetcode.com/problems/consecutive-numbers-sum "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/consecutive-numbers-sum "困难：1694") 查看该题。

??? info "解题思路"
    枚举和为 `n` 的连续正整数的长度 `k` 。

    === "Go"
        ```go
        --8<-- "enumerate/go/lc829.go"
        ```

## CF118C: Fancy Number

???+ note "问题描述"
    第一行输入 $n$ 和 $k$ 两个整数，$(2≤k≤n≤1e4)$ 。<br>
    第二行输入长度为 $n$ 的字符串 $s$，仅包含 '0'~'9'。<br>
    每次操作你可以把一个 s[i] 修改成任意 '0'~'9'，假设修改成 $b$ ，则花费为 $abs(s[i]-b)$ 。<br>
    要使 $s$ 中至少有 $k$ 个相同字符，且在总花费最小的前提下，让修改后的 $s$ 的字典序尽量小。<br>
    输出最小总花费以及修改后的 $s$ 。

    在 [Codeforces](https://codeforces.com/problemset/problem/118/C "1900")
    或 [洛谷](https://www.luogu.com.cn/problem/CF118C "普及+/提高") 查看该题。

??? info "解题思路"
    枚举要修改出 $k$ 个的字符 $b$ 。

    === "Go"
        ```go
        --8<-- "enumerate/go/cf118c.go"
        ```
    === "Java"
        ```java
        --8<-- "enumerate/java/cf118c/Main.java"
        ```
