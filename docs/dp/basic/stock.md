# 买卖股票

## LC121: 买卖股票的最佳时机

???+ note "问题描述：单次买卖能获得的最大利润"
    给定一个长度为 `n(1≤n≤1e5)` 的整数数组 `prices` ，`0≤prices[i]≤1e4` 表示一支股票第 `i` 天的价格。<br>
    你只能选择 **某一天** 买入这只股票，并选择在 **未来的某一个不同的日子** 卖出该股票。<br>
    返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 `0` 。

    在 [LeetCode主站](https://leetcode.com/problems/best-time-to-buy-and-sell-stock "Easy")
    或 [力扣中文社区](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock "简单") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "basic/stock/go/lc121.go"
        ```

## LC122: 买卖股票的最佳时机 II

???+ note "问题描述：多次买卖能获得的最大利润"
    给定一个长度为 `n(1≤n≤3e4)` 的整数数组 `prices` ，`0≤prices[i]≤1e4` 表示一支股票第 `i` 天的价格。<br>
    在每一天，你可以决定是否购买、出售股票。你在任何时候 **最多** 只能持有 **一股** 股票。你也可以先购买，然后在 **同一天** 出售。<br>
    返回你能获得的最大利润。

    在 [LeetCode主站](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii "中等") 查看该题。

??? info "解题思路"
    **方法一：动态规划**

    === "Go"
        ```go
        --8<-- "basic/stock/go/lc122_1.go"
        ```

    **方法二：贪心**

    === "Go"
        ```go
        --8<-- "basic/stock/go/lc122_2.go"
        ```

## LC123: 买卖股票的最佳时机 III

???+ note "问题描述：最多买卖两次能获得的最大利润"
    给定一个长度为 `n(1≤n≤1e5)` 的整数数组 `prices` ，`0≤prices[i]≤1e5` 表示一支股票第 `i` 天的价格。<br>
    设计一个算法来计算你所能获取的最大利润。你最多可以完成 **两笔** 交易。<br>
    注意：你不能同时参与多笔交易，必须在再次购买前出售掉之前的股票。

    在 [LeetCode主站](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii "困难") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "basic/stock/go/lc123.go"
        ```

## LC188: 买卖股票的最佳时机 IV

???+ note "问题描述：最多买卖k次能获得的最大利润"
    给定一个长度为 `n(0≤n≤1e3)` 的整数数组 `prices` ，`0≤prices[i]≤1e3` 表示一支股票第 `i` 天的价格。<br>
    设计一个算法来计算你所能获取的最大利润。你最多可以完成 `k(0≤k≤100)` 笔交易。<br>
    注意：你不能同时参与多笔交易，必须在再次购买前出售掉之前的股票。

    在 [LeetCode主站](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv "困难") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "basic/stock/go/lc188.go"
        ```

## LC309: 买卖股票的最佳时机含冷冻期

???+ note "问题描述"
    给定一个长度为 `n(1≤n≤5e3)` 的整数数组 `prices` ，`0≤prices[i]≤1e3` 表示一支股票第 `i` 天的价格。<br>
    设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:

    * 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。

    注意：你不能同时参与多笔交易，必须在再次购买前出售掉之前的股票。

    在 [LeetCode主站](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown "中等") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "basic/stock/go/lc309.go"
        ```

## LC714: 买卖股票的最佳时机含手续费

???+ note "问题描述"
    给定一个长度为 `n(1≤n≤5e4)` 的整数数组 `prices` ，`1≤prices[i]≤5e4` 表示一支股票第 `i` 天的价格。整数 `fee(0≤fee≤5e4)` 代表了交易股票的手续费用。<br>
    你可以无限次地完成交易，但是你每笔交易都需要付手续费。你不能同时参与多笔交易，必须在再次购买前出售掉之前的股票。<br>
    返回你能获得的最大利润。

    在 [LeetCode主站](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee "中等") 查看该题。

??? info "解题思路"
    **方法一：动态规划**

    === "Go"
        ```go
        --8<-- "basic/stock/go/lc714_1.go"
        ```

    **方法二：贪心**

    === "Go"
        ```go
        --8<-- "basic/stock/go/lc714_2.go"
        ```
