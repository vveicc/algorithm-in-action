# 构造

## CF1554C: Mikasa

???+ note "问题描述"
    输入 $t(1≤t≤3e4)$ 表示 $t$ 组数据，每组数据输入两个整数 $n$ 和 $m$，均在 $[0,1e9]$ 范围内。<br>
    定义数组 $a = [n \oplus 0, n \oplus 1, n \oplus 2, ..., n \oplus m]$。<br>
    输出不在 $a$ 中的最小非负整数。

    在 [Codeforces](https://codeforces.com/problemset/problem/1554/C "1800")
    或 [洛谷](https://www.luogu.com.cn/problem/CF1554C "普及/提高-") 上查看该题。

??? info "解题思路"
    **方法一：从高到低按位构造**
    
    根据题意，需要找到最小的 $x$ ，使得 $n \oplus x > m$ 。<br>
    显然，当 $n > m$ 时，答案为 0 。<br>
    令 $k = m+1$ ，则 $x$ 需要满足 $n \oplus x ≥ k$。<br>
    分别使用 $n_i, k_i, x_i$ 表示 $n, k, x$ 的二进制第 $i$ 位。<br>
    从高位到低位按位构造，对于第 $i$ 位，有以下四种情况：
    
    | $n_i$ | $k_i$ | 如何处理能够保证 $n \oplus x ≥ k$ 且 $x$ 最小？|
    | :-: | :-: | :--------- |
    | 0 | 0 | 将 $x_i$ 赋为 0 |
    | 1 | 1 | 将 $x_i$ 赋为 0 |
    | 0 | 1 | 将 $x_i$ 赋为 1 ，如果赋为 0 将会导致 $n \oplus x < k$ |
    | 1 | 0 | 将 $x_i$ 赋为 0 ，此时已经能保证 $n \oplus x ≥ k$，为了使 $x$ 最小，需要将剩余的二进制位全部赋为 0 |

    === "Go"
        ```go
        --8<-- "construction/go/cf1554c_1.go"
        ```
    === "Java"
        ```java
        --8<-- "construction/java/cf1554c_1/Main.java"
        ```

    **方法二：基于 $n \oplus m$ 构造**
    
    根据题意，需要找到最小的 $x$ ，使得 $n \oplus x > m$ 。<br>
    显然，当 $n > m$ 时，答案为 0 。<br>
    令 $x = n \oplus m$ ，此时 $x$ 不满足 $n \oplus x > m$ ，考虑如何修改 $x$ 能够满足条件。<br>
    分别使用 $n_i, m_i, x_i$ 表示 $n, m, x$ 的二进制第 $i$ 位。<br>
    对于 $x$ 的每一个二进制位，有以下四种情况：

    | $n_i$ | $m_i$ | 如何修改 $x$ 能够增大 $n \oplus x$ 且保证 $x$ 最小？|
    | :-: | :-: | :--------- |
    | 0 | 1 | 不能修改 $x_i$ ，将 $x_i$ 改为 0 会导致 $n \oplus x$ 减小 |
    | 1 | 1 | 不能修改 $x_i$ ，将 $x_i$ 改为 1 会导致 $n \oplus x$ 减小 |
    | 1 | 0 | 将 $x_i$ 改为 0 能够增大 $n \oplus x$ ，同时导致 $x$ 减小，为了使 $x$ 最小，需要修改该情况的最高位，并将其后低位全部置 0 |
    | 0 | 0 | 将 $x_i$ 改为 1 能够增大 $n \oplus x$ ，同时导致 $x$ 增大，为了使 $x$ 最小，需要修改该情况的最低位，并将其后低位全部置 0 |

    注意，第三种情况是最优的，但可能不存在这种情况，此时需要通过第四种情况修改 $x$ 。

    === "Go"
        ```go
        --8<-- "construction/go/cf1554c_2.go"
        ```
    === "Java"
        ```java
        --8<-- "construction/java/cf1554c_2/Main.java"
        ```

## LC2375: 根据模式串构造最小数字

???+ note "问题描述"
    给你下标从 `0` 开始、长度为 `n(1≤n≤8)` 的字符串 `pattern` ，只包含 `'I'` 和 `'D'` 。<br>
    你需要构造一个下标从 `0` 开始长度为 `n + 1` 的字符串，且它要满足以下条件：

    * `num` 包含数字 `'1'` 到 `'9'` ，其中每个数字 至多 使用一次。
    * 如果 `pattern[i] == 'I'` ，那么 `num[i] < num[i + 1]` 。
    * 如果 `pattern[i] == 'D'` ，那么 `num[i] > num[i + 1]` 。
    
    请你返回满足上述条件字典序最小的字符串 `num` 。

    在 [LeetCode主站](https://leetcode.com/problems/construct-smallest-number-from-di-string "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/construct-smallest-number-from-di-string "中等：1642") 上查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "construction/go/lc2375.go"
        ```

## CF1304D: Shortest and Longest LIS

???+ note "问题描述"
    输入 $t(1≤t≤1e4)$ 表示 $t$ 组数据。所有数据的 $n$ 之和 $≤2e5$。<br>
    每组数据输入 $n(2≤n≤2e5)$ 和长为 $n-1$ 的字符串 $s$，仅包含 `'<'` 和 `'>'`。
    
    * `s[i] = '<'` 表示 `a[i] < a[i+1]` ；
    * `s[i] = '>'` 表示 `a[i] > a[i+1]` 。
    
    请构造两个 $1$ ~ $n$ 的排列，符合字符串 $s$，且第一个数组的 LIS 最短，第二个数组的 LIS 最长。<br>
    如果有多种构造方案，输出任意一种。

    在 [Codeforces](https://codeforces.com/problemset/problem/1304/D "1800")
    或 [洛谷](https://www.luogu.com.cn/problem/CF1304D "普及+/提高") 上查看该题。

    如果觉得这题比较难，可以先做简单版：[LC2375: 根据模式串构造最小数字](#lc2375-根据模式串构造最小数字)。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "construction/go/cf1304d.go"
        ```
    === "Java"
        ```java
        --8<-- "construction/java/cf1304d/Main.java"
        ```
