# 模拟

## LC43: 字符串相乘

???+ note "问题描述"
    给定两个以字符串形式表示的非负整数 $num1$ 和 $num2$ ，字符串长度范围：$[1,200]$ 。<br>
    $num1$ 和 $num2$ 都不包含任何前导零，除了数字 $0$ 本身。<br>
    返回 $num1$ 和 $num2$ 的乘积，以字符串形式表示。

    在 [LeetCode主站](https://leetcode.com/problems/multiply-strings "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/multiply-strings "中等") 上查看该题。

??? info "解题思路"
    **方法一：做加法**

    === "Go"
        ```go
        --8<-- "simulate/go/lc43_1.go"
        ```

    **方法二：做乘法**

    === "Go"
        ```go
        --8<-- "simulate/go/lc43_2.go"
        ```

## LC2532: 过桥的时间

???+ note "问题描述"
    共有 $k$ 位工人计划将 $n$ 个箱子从旧仓库移动到新仓库。<br>
    给你两个整数 $n([1,1e4])$ 和 $k([1,1e4])$，以及一个二维整数数组 $time$ ，数组的大小为 $k * 4$ 。<br>
    其中 $time[i] = [leftToRight_i, pickOld_i, rightToLeft_i, putNew_i]$ 且 $1≤time[i][j]≤1e3$ 。

    ??? quote "箱子搬运规则"
        一条河将两座仓库分隔，只能通过一座桥通行。旧仓库位于河的右岸，新仓库在河的左岸。<br>
        开始时，所有 $k$ 位工人都在桥的左侧等待。为了移动这些箱子，第 $i$ 位工人（下标从 $0$ 开始）可以：

        * 从左岸（新仓库）跨过桥到右岸（旧仓库），用时 $leftToRight_i$ 分钟。
        * 从旧仓库选择一个箱子，并返回到桥边，用时 $pickOld_i$ 分钟。不同工人可以同时搬起所选的箱子。
        * 从右岸（旧仓库）跨过桥到左岸（新仓库），用时 $rightToLeft_i$ 分钟。
        * 将箱子放入新仓库，并返回到桥边，用时 $putNew_i$ 分钟。不同工人可以同时放下所选的箱子。
        
        如果满足下面任一条件，则认为工人 $i$ 的效率低于工人 $j$ ：

        * $leftToRight_i + rightToLeft_i > leftToRight_j + rightToLeft_j$
        * $leftToRight_i + rightToLeft_i == leftToRight_j + rightToLeft_j$ 且 $i > j$
        
        工人通过桥时需要遵循以下规则：

        * 如果工人 $x$ 到达桥边时，工人 $y$ 正在过桥，那么工人 $x$ 需要在桥边等待。
        * 如果没有正在过桥的工人，那么在桥右边等待的工人可以先过桥。如果同时有多个工人在右边等待，那么效率最低的工人会先过桥。
        * 如果没有正在过桥的工人，且桥右边也没有在等待的工人，同时旧仓库还剩下至少一个箱子需要搬运，此时在桥左边的工人可以过桥。如果同时有多个工人在左边等待，那么效率最低的工人会先过桥。
    
    所有 $n$ 个箱子都需要放入新仓库，请你返回最后一个搬运箱子的工人到达河左岸的时间。

    在 [LeetCode主站](https://leetcode.com/problems/time-to-cross-a-bridge "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/time-to-cross-a-bridge "困难：2589") 上查看该题。

??? info "解题思路"
    **方法一：四堆模拟**

    === "Go"
        ```go
        --8<-- "simulate/go/lc2532.go"
        ```
