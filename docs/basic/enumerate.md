# 枚举

## LC1615. 最大网络秩

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

## LC829. 连续整数求和

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

## LC1638. 统计只差一个字符的子串数目 { data-toc-label='LC1638. 统计只差一个字符的...' }

???+ note "问题描述"
    给你两个长度为 `n(1≤n≤100)` 且仅包含小写英文字母的字符串 `s` 和 `t` ，请你找出 `s` 中的非空子串的数目，这些子串满足替换 **一个不同字符** 以后，是 `t` 串的子串。换言之，请你找到 `s` 和 `t` 串中 **恰好** 只有一个字符不同的子字符串对的数目。

    在 [LeetCode主站](https://leetcode.com/problems/consecutive-numbers-sum "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/consecutive-numbers-sum "中等：1745") 查看该题。

??? info "解题思路"
    两个子串长度相等，恰好有一个字符不同。

    **方法一：暴力枚举**

    === "Go"
        ```go
        --8<-- "enumerate/go/lc1638_1.go"
        ```

    **方法二：动态规划 + 枚举**

    === "Go"
        ```go
        --8<-- "enumerate/go/lc1638_2.go"
        ```

    **方法三：总结规律 + 枚举**

    === "Go"
        ```go
        --8<-- "enumerate/go/lc1638_3.go"
        ```

## CF118C. Fancy Number

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

## LC1625. 操作后字典序最小的字符串 { data-toc-label='LC1625. 操作后字典序最小的...' }

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

## CF1626D. Martial Arts Tournament { data-toc-label='CF1626D. Martial Arts Tour...' }

???+ note "问题描述"
    第一行输入 $t(1≤t≤1e4)$ 表示 $t$ 组数据。所有数据的 $n$ 之和 $≤2e5$。<br>
    每组数据第一行输入 $n(1≤n≤2e5)$ ，第二行输入长为 $n$ 的整数数组 $a(1≤a[i]≤n)$ 。<br>
    你需要选择两个整数 $x\ y\ (x<y)$ ，把 $a$ 中小于 $x$ 的数分为一组，大于等于 $y$ 的分为一组，其余的分为一组，一共三组。
    对每一组，如果组的大小不是 $2$ 的幂次，则增加到最近的 $2$ 的幂次，花费为增量。比如 $5$ 补齐到 $8$ ，花费为 $8-5=3$ 。如果已经是 $2$ 的幂次，则花费为 $0$ 。<br>
    计算花费之和的最小值。

    在 [Codeforces](https://codeforces.com/problemset/problem/1626/D "2100")
    或 [洛谷](https://www.luogu.com.cn/problem/CF1626D "普及+/提高") 查看该题。

??? info "解题思路"
    枚举 $x$ 和 $y$ 显然会超时，不妨换个思路，枚举与组的大小最接近的 $2$ 的幂次。

    === "Go"
        ```go
        --8<-- "enumerate/go/cf1626d.go"
        ```
    === "Java"
        ```java
        --8<-- "enumerate/java/cf1626d/Main.java"
        ```

## CF1181C. Flag

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

## LC2552. 统计上升四元组

???+ note "问题描述"
    给你一个长度为 `4≤n≤4e3` 下标从 `0` 开始的整数数组 `nums` ，其中包含 `1` 到 `n` 的 **所有** 数字。<br>
    如果一个四元组 `(i, j, k, l)` 满足以下条件，我们称它是上升的：

    - `0 <= i < j < k < l < n` 且
    - `nums[i] < nums[k] < nums[j] < nums[l]` 。

    请你返回上升四元组的数目。

    在 [LeetCode主站](https://leetcode.com/problems/count-increasing-quadruplets "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/count-increasing-quadruplets "困难：2433") 查看该题。

??? info "解题思路"

    === "Go"
        ```go
        --8<-- "enumerate/go/lc2552.go"
        ```

## LC1330. 翻转子数组得到最大的数组值 { data-toc-label='LC1330. 翻转子数组得到最大...' }

???+ note "问题描述"
    给你一个长度为 `1≤n≤3e4` 的整数数组 `nums` ，`-1e5≤nums[i]≤1e5` 。<br>
    **数组值** 定义为所有满足 `0≤i<n-1` 的 `|nums[i]-nums[i+1]|` 的和。<br>
    你可以选择给定数组的任意子数组，并将该子数组翻转。但你只能执行这个操作 **一次** 。<br>
    请你计算可以得到的最大 **数组值** 。

    在 [LeetCode主站](https://leetcode.com/problems/reverse-subarray-to-maximize-array-value "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/reverse-subarray-to-maximize-array-value "困难：2482") 查看该题。

??? info "解题思路"
    翻转 $nums[i..j]$ ，先不考虑边界情况，记 $a=nums[i-1], b=nums[i], c=nums[j], d=nums[j+1]$ 。<br>
    很明显，翻转后数组值的变化量为 $x = |a-c| + |b-d| - |a-b| - |c-d|$ ，问题转化为将 $x$ 最大化。

    根据 $a,b,c,d$ 的大小关系，共有 $4! = 24$ 种情况，下面分情况讨论：
    
    **第 1 类** $\max(a,b) ≤ \min(c,d)$
    
    $$
    \begin{align*}
        x &= |a-c| + |b-d| - |a-b| - |c-d| \\
          &= (c-a) + (d-b) - |a-b| - |c-d| \\
          &= (c+d-|c-d|) - (a+b+|a-b|) \\
          &= 2 \times (\min(c,d) - \max(a,b)) ≥ 0
    \end{align*}
    $$

    根据对称性，$\max(c,d) ≤ \min(a,b)$ 时，也有 $x = 2 \times (\min(a,b) - \max(c,d)) ≥ 0$ 。

    **第 2 类** $\max(a,c) ≤ \min(b,d)$

    $$
    \begin{align*}
        x &= |a-c| + |b-d| - |a-b| - |c-d| \\
          &= |a-c| + |b-d| - (b-a) - (d-c) \\
          &= (a+c+|a-c|) - (b+d-|b-d|) \\
          &= 2 \times (\max(a,c) - \min(b,d)) ≤ 0
    \end{align*}
    $$

    根据对称性，$\max(b,d) ≤ \min(a,c)$ 时，也有 $x = 2 \times (\max(b,d) - \min(a,c)) ≤ 0$ 。

    **第 3 类** $\max(a,d) ≤ \min(b,c)$

    $$
    \begin{align*}
        x &= |a-c| + |b-d| - |a-b| - |c-d| \\
          &= (c-a) + (b-d) - (b-a) - (c-d) = 0
    \end{align*}
    $$

    根据对称性，$\max(b,d) ≤ \min(a,c)$ 时，也有 $x = 0$ 。

    综合以上三类情况，仅第 1 类会使数组值变大，此时

    $$\max x = 2 \times (\max_{i=0}^{n-1}\{\min(nums[i],nums[i+1])\} - \min_{j=0}^{n-1}\{\max(nums[j],nums[j+1])\})$$

    实现时注意处理边界情况。

    === "Go"
        ```go
        --8<-- "enumerate/go/lc1330.go"
        ```
