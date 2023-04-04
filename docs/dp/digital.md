# 数位 DP

## LC357. 各位数字都不同的数字

???+ note "问题描述"
    给你一个整数 $n(0≤n≤8)$ ，统计并返回各位数字都不同的数字 $x(0≤x<10^n)$ 的个数。

    在 [LeetCode主站](https://leetcode.com/problems/count-numbers-with-unique-digits "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/count-numbers-with-unique-digits "中等") 查看该题。

??? info "解题思路"
    **数位 DP + 记忆化搜索**

    === "Go"
        ```go
        --8<-- "digital/go/lc357.go"
        ```

## LC1012. 至少有 1 位重复的数字

???+ note "问题描述"
    给定正整数 $n(1≤n≤1e9)$，返回在 $[1, n]$ 范围内具有 **至少 $1$ 位** 重复数字的正整数的个数。

    在 [LeetCode主站](https://leetcode.com/problems/numbers-with-repeated-digits "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/numbers-with-repeated-digits "困难：2230") 查看该题。

    如果要统计每个数位互不相同的正整数的个数呢？参考 [LC2376. 统计特殊整数](https://leetcode.cn/problems/count-special-integers "困难：2120")。

??? info "解题思路"
    **数位 DP + 记忆化搜索**

    === "Go"
        ```go
        --8<-- "digital/go/lc1012.go"
        ```

## LC600. 不含连续 1 的非负整数

???+ note "问题描述"
    给定一个正整数 $n(1≤n≤1e9)$ ，统计在 $[0, n]$ 范围的整数中，二进制表示不存在 **连续的 $1$** 的整数的个数。

    在 [LeetCode主站](https://leetcode.com/problems/non-negative-integers-without-consecutive-ones "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/non-negative-integers-without-consecutive-ones "困难") 查看该题。

??? info "解题思路"
    **方法一：2进制 + 数位 DP + 记忆化搜索**

    === "Go"
        ```go
        --8<-- "digital/go/lc600_1.go"
        ```

    **方法二：4进制 + 数位 DP + 记忆化搜索**

    === "Go"
        ```go
        --8<-- "digital/go/lc600_2.go"
        ```

## LC902. 最大为 N 的数字组合

???+ note "问题描述"
    给定一个按 **严格递增顺序** 排列的数字数组 `digits` ，`'1'≤digits[i]≤'9'` 。<br>
    你可以用任意次数的 `digits[i]` 来写数字。例如，如果 `digits = ['1','3','5']` ，我们可以写数字，如 `'13'`, `'551'`, 和 `'1351315'` 。<br>
    返回可以生成的小于或等于给定整数 `n` 的正整数的个数 。

    在 [LeetCode主站](https://leetcode.com/problems/numbers-at-most-n-given-digit-set "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/numbers-at-most-n-given-digit-set "困难：1990") 查看该题。

??? info "解题思路"
    **数位 DP + 记忆化搜索**

    === "Go"
        ```go
        --8<-- "digital/go/lc902.go"
        ```

## LC788. 旋转数字

???+ note "问题描述"
    我们称一个数 X 为好数, 如果它的每位数字逐个地被旋转 180 度后，我们仍可以得到一个有效的，且和 X 不同的数。要求每位数字都要被旋转。<br>
    如果一个数的每位数字被旋转以后仍然还是一个数字， 则这个数是有效的。0, 1, 和 8 被旋转后仍然是它们自己；2 和 5 可以互相旋转成对方（在这种情况下，它们以不同的方向旋转，换句话说，2 和 5 互为镜像）；6 和 9 同理，除了这些以外其他的数字旋转以后都不再是有效的数字。<br>
    现在我们有一个正整数 `N(1≤N≤1e4)` , 计算从 `1` 到 `N` 中有多少个数 X 是好数？

    在 [LeetCode主站](https://leetcode.com/problems/rotated-digits "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/rotated-digits "中等：1397") 查看该题。

??? info "解题思路"
    **数位 DP + 记忆化搜索**

    题目数据范围较小，可以直接枚举统计。但如果数据范围增大至 `1e9` ，则需要通过数位 DP 求解。

    === "Go"
        ```go
        --8<-- "digital/go/lc788.go"
        ```

## LC233. 数字 1 的个数

???+ note "问题描述"
    给定一个整数 $n(0≤n≤1e9)$ ，计算所有小于等于 $n$ 的非负整数中数字 $1$ 出现的个数。

    在 [LeetCode主站](https://leetcode.com/problems/number-of-digit-one "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/number-of-digit-one "困难") 查看该题。

    如果要统计数字 $2$ 出现的个数呢？参考 [LC面试题 17.06. 2出现的次数](https://leetcode.cn/problems/number-of-2s-in-range-lcci "困难")。

??? info "解题思路"
    **方法一：数学**

    累加每一个数位上数字 $d$ 出现的个数，本题中 $d = 1$ 。<br>
    以 $n = 1234567$ 为例，计算百位上出现数字 $d$ 的次数。<br>
	首先 $n$ 有 $1234$ 个完整的 $[000,999]$ 的循环，每个循环百位上都会出现 $100$ 次数字 $d$ 。<br>
	对于剩余的 $m = 567$ ，分情况讨论：

    * 如果 $m<d*100$ ，则百位上不会出现数字 $d$ ；
	* 如果 $d*100≤m<(d+1)*100$ ，则百位上会出现 $m-d*100+1$ 次数字 $d$ ；
	* 如果 $m≥(d+1)*100$ ，则百位上会出现全部的 $100$ 次数字 $d$ 。

    === "Go"
        ```go
        --8<-- "digital/go/lc233_1.go"
        ```

    **方法二：数位 DP + 记忆化搜索**

    === "Go"
        ```go
        --8<-- "digital/go/lc233_2.go"
        ```

## LG2602. [ZJOI2010] 数字计数

???+ note "问题描述"
    输入两个正整数 $a$ 和 $b$ ，$1≤a≤b≤1e12$，求在 $[a,b]$ 中的所有整数中，每个数码(digit)各出现了多少次。

    在 [洛谷](https://www.luogu.com.cn/problem/P2602 "提高+/省选-") 查看该题。

??? info "解题思路"
    **数位 DP + 记忆化搜索**

    === "Go"
        ```go
        --8<-- "digital/go/lg2602.go"
        ```
    === "Java"
        ```java
        --8<-- "digital/java/lg2602/Main.java"
        ```

## LC1397. 找到所有好字符串

???+ note "问题描述"
    给你两个长度为 `n(1≤n≤500)` 的字符串 `s1` 和 `s2` ，`s1<=s2` ，以及一个字符串 `evil` 。<br>
    **好字符串** 的定义为：长度为 `n` ，字典序大于等于 `s1` 小于等于 `s2` ，且不包含 `evil` 为子字符串。<br>
    请你返回好字符串的数目。由于答案可能很大，请你返回答案对 `1e9 + 7` 取余的结果。

    在 [LeetCode主站](https://leetcode.com/problems/find-all-good-strings "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/find-all-good-strings "困难：2667") 查看该题。

??? info "解题思路"
    **KMP + 数位 DP + 记忆化搜索**

    === "Go"
        ```go
        --8<-- "digital/go/lc1397.go"
        ```
