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

## LC1625: 执行操作后字典序最小的字符串

???+ note "问题描述"
    给你一个长度为 `n(2≤n≤100)` 的字符串 `s` 以及两个整数 `a(0≤a≤9)` 和 `b(1≤b<n)` 。<br>
    其中，`n` 为偶数，且字符串 `s` 仅由数字 `0` 到 `9` 组成。你可以在 `s` 上按任意顺序多次执行下面两个操作之一：

    * 累加：将 `a` 加到 `s` 中所有下标为奇数的元素上（下标从 `0` 开始）。数字一旦超过 `9` 就会变成 `0`，如此循环往复。例如，`s = "3456"` 且 `a = 5`，则执行此操作后 `s` 变成 `"3951"`。
    * 轮转：将 `s` 向右轮转 `b` 位。例如，`s = "3456"` 且 `b = 1`，则执行此操作后 `s` 变成 `"6345"`。
    
    请你返回在 `s` 上执行上述操作任意次后可以得到的字典序最小的字符串。

    在 [LeetCode主站](https://leetcode.com/problems/consecutive-numbers-sum "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/consecutive-numbers-sum "中等：1992") 查看该题。

??? info "解题思路"
    累加操作执行 `10 / GCD(10, a)` 次后复原。<br>
    轮转操作执行 `n / GCD(n, b)` 次后复原。<br>
    `b` 为偶数时，只有奇数索引位置可以执行累加操作；<br>
    `b` 为奇数时，奇数索引位置和偶数索引位置可以独立执行累加操作。

    === "Go"
        ```go
        --8<-- "enumerate/go/lc1625.go"
        ```

## CF1181C: Flag

???+ note "问题描述"
    第一行输入 $n$ 和 $m$ 两个整数，$(1≤n,m≤1e3)$ 。<br>
    接下来 $n$ 行每行输入一个长度为 $m$ 的字符串，形成一个 $n \times m$ 的字符矩阵，元素都是小写字母。<br>
    定义「国旗」为一个 $3 \times h$ 行的子矩阵，前 $h$ 行的字符都相同，中间 $h$ 行的字符都相同，后 $h$ 行的字符都相同，它们分别记作 A B C，要求 A 和 B 的字符不同，B 和 C 的字符不同（A 和 C 无要求）。<br>
    输出是国旗的子矩阵的数量。

    在 [Codeforces](https://codeforces.com/problemset/problem/1181/C "1900")
    或 [洛谷](https://www.luogu.com.cn/problem/CF1181C "提高+/省选-") 查看该题。

??? info "解题思路"
    按列枚举国旗右边界，按行枚举中间 $h$ 行。

    === "Go"
        ```go
        --8<-- "enumerate/go/cf1181c.go"
        ```
    === "Java"
        ```java
        --8<-- "enumerate/java/cf1181c/Main.java"
        ```
