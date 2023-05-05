# 最短路

## Floyd 算法

Floyd 算法适合用来求出图中 **任意两点** 间的最短路径。

适用于任何图，无论有向无向，边权正负。

第一维度压缩后的状态转移方程：

$$f_{x,y} = \min(f_{x,y}, f_{x,k} + f_{k,y})$$

简单来说就是，允许经过节点 `k` 的情况下，节点 `x` 到节点 `y` 的最短路径可能变小。

状态转移时可以按照任意顺序枚举节点 `k` 。

???+ tip "算法实现"
    === "Go"
        ```go
        for k := 1; k <= n; k++ {
            for x := 1; x <= n; x++ {
                for y := 1; y <= n; y++ {
                    f[x][y] = min(f[x][y], f[x][k] + f[k][y])
                }
            }
        }
        ```

### CF295B. Greg and Graph

???+ note "问题描述"
    第一行输入 `n(1≤n≤500)` 表示 `n` 个点的有向完全图，节点编号从 1 开始。<br>
    接下来 `n` 行每行输入 `n` 个整数，表示 `n x n` 的邻接矩阵 `a` ，其中 `a[i][j]` 表示 `i` 到 `j` 的边权。<br>
    除 `a[i][i] = 0` 外，边权值的范围：`[1, 1e5]` 。<br>
    最后一行输入 `1 ~ n` 的排列，表示要一个个地删除图上的点，每删除一个点，这个点的出边和入边都会被删除。<br>
    输出 `n` 个数，第 `i` 个数表示第 `i` 次删除之前，所有剩余点对的最短路之和。

    在 [Codeforces](https://codeforces.com/problemset/problem/295/B "1700")
    或 [洛谷](https://www.luogu.com.cn/problem/CF295B "普及+/提高") 查看该题。

??? info "解题思路"
    正难则反，正序删点转化为倒序加点。

    === "Go"
        ```go
        --8<-- "shortest-path/floyd/go/cf295b.go"
        ```
    === "Java"
        ```java
        --8<-- "shortest-path/floyd/java/cf295b/Main.java"
        ```

### LC2642. 求最短路径的图类

???+ note "问题描述"
    给你一个有 `n` 个节点的 **有向带权** 图，节点编号为 `0` 到 `n - 1` 。<br>
    图中的初始边用数组 `edges` 表示，其中 $edges[i] = [from_i, to_i, cost_i]$ 表示从 $from_i$ 到 $to_i$ 有一条代价为 $cost_i$ 的边。

    请你实现一个 `Graph` 类：

    - `Graph(int n, int[][] edges)` 初始化图有 `n` 个节点，并输入初始边。
    - `addEdge(int[] edge)` 向边集中添加一条边，其中 `edge = [from, to, cost]` 。数据保证添加这条边之前对应的两个节点之间没有有向边。
    - `int shortestPath(int x, int y)` 返回从节点 `x` 到 `y` 的路径 **最小** 代价。如果路径不存在，返回 `-1` 。一条路径的代价是路径中所有边代价之和。

    ??? quote "约束条件"
        - `1 ≤ n ≤ 100`
        - `0 ≤ edges.length ≤ n * (n - 1)`
        - `edges[i].length == edge.length == 3`
        - $1 ≤ cost_i, cost ≤ 1e6$
        - 图中任何时候都不会有重边和自环。
        - 调用 `addEdge` 至多 `100` 次。
        - 调用 `shortestPath` 至多 `100` 次。

    在 [LeetCode主站](https://leetcode.com/problems/design-graph-with-shortest-path-calculator "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/design-graph-with-shortest-path-calculator "困难：1811") 查看该题。

??? info "解题思路"
    **方法一：[朴素（暴力）的 Dijkstra 算法](#lc2642-求最短路径的图类_1)**

    **方法二：[优先队列优化的 Dijkstra 算法](#lc2642-求最短路径的图类_1)**

    **方法三：Floyd 算法**

    本题调用 `addEdge` 和 `shortestPath` 的次数都比较少，使用 Floyd 算法或者 Dijkstra 算法都可以。<br>
    如果调用 `shortestPath` 次数比较多，则使用 Floyd 算法更合适。

    === "Go"
        ```go
        --8<-- "shortest-path/floyd/go/lc2642.go"
        ```

## Dijkstra 算法

Dijkstra 算法适合用来求出无负权边图中的单源最短路径。其中：

* **无负权边** 表示图中所有边的权值必须为非负数；
* **单源最短路径** 表示 Dijkstra 算法可以求出从某一个节点到其余所有节点的最短路径。

记 $V$ 是图中的节点数，$E$ 是图中的边数。<br>
朴素（暴力）的 Dijkstra 算法的时间复杂度为：$O(V^2)$ ，在 **稠密图** 中效率更优。<br>
优先队列优化的 Dijkstra 算法的时间复杂度为：$O(E\log E)$ ，在 **稀疏图** 中效率更优。

### LC2642. 求最短路径的图类

???+ note "问题描述"
    给你一个有 `n` 个节点的 **有向带权** 图，节点编号为 `0` 到 `n - 1` 。<br>
    图中的初始边用数组 `edges` 表示，其中 $edges[i] = [from_i, to_i, cost_i]$ 表示从 $from_i$ 到 $to_i$ 有一条代价为 $cost_i$ 的边。

    请你实现一个 `Graph` 类：

    - `Graph(int n, int[][] edges)` 初始化图有 `n` 个节点，并输入初始边。
    - `addEdge(int[] edge)` 向边集中添加一条边，其中 `edge = [from, to, cost]` 。数据保证添加这条边之前对应的两个节点之间没有有向边。
    - `int shortestPath(int x, int y)` 返回从节点 `x` 到 `y` 的路径 **最小** 代价。如果路径不存在，返回 `-1` 。一条路径的代价是路径中所有边代价之和。

    ??? quote "约束条件"
        - `1 ≤ n ≤ 100`
        - `0 ≤ edges.length ≤ n * (n - 1)`
        - `edges[i].length == edge.length == 3`
        - $1 ≤ cost_i, cost ≤ 1e6$
        - 图中任何时候都不会有重边和自环。
        - 调用 `addEdge` 至多 `100` 次。
        - 调用 `shortestPath` 至多 `100` 次。

    在 [LeetCode主站](https://leetcode.com/problems/design-graph-with-shortest-path-calculator "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/design-graph-with-shortest-path-calculator "困难：1811") 查看该题。

??? info "解题思路"
    **方法一：[Floyd 算法](#lc2642-求最短路径的图类)**

    ??? tip "方法二：朴素（暴力）的 Dijkstra 算法"
        === "Go"
            ```go
            --8<-- "shortest-path/dijkstra/go/lc2642_1.go"
            ```

    ??? tip "方法三：优先队列优化的 Dijkstra 算法"
        === "Go"
            ```go
            --8<-- "shortest-path/dijkstra/go/lc2642_2.go"
            ```

### LC2662. 前往目标的最小代价

???+ note "问题描述"
    给你一个整数数组 `start` ，其中 `start = [startX, startY]` 表示初始位置。<br>
    给你一个整数数组 `target` ，其中 `target = [targetX, targetY]` 表示目标位置。<br>
    从位置 `(x1, y1)` 到任一其他位置 `(x2, y2)` 的代价是 `|x2 - x1| + |y2 - y1|` 。<br>
    给你一个二维整数数组 `specialRoads` ，表示空间中存在的一些特殊路径。你可以使用每条特殊路径任意次数。<br>
    其中 $specialRoads[i] = [x1_i, y1_i, x2_i, y2_i, cost_i]$ 表示第 `i` 条特殊路径可以从 $(x1_i, y1_i)$ 到 $(x2_i, y2_i)$ ，但代价等于 $cost_i$ 。<br>
    返回从 `(startX, startY)` 到 `(targetX, targetY)` 所需的最小代价。

    ??? quote "约束条件"
        - `1 <= startX <= targetX <= 1e5`
        - `1 <= startY <= targetY <= 1e5`
        - `1 <= specialRoads.length <= 200`
        - $startX <= x1_i, x2_i <= targetX$
        - $startY <= y1_i, y2_i <= targetY$
        - $1 <= cost_i <= 1e5$

    在 [LeetCode主站](https://leetcode.com/problems/minimum-cost-of-a-path-with-special-roads "Medium")
    或 [力扣中文社区](https://leetcode.cn/problems/minimum-cost-of-a-path-with-special-roads "中等") 查看该题。

??? info "解题思路"
    === "Go"
        ```go
        --8<-- "shortest-path/dijkstra/go/lc2662.go"
        ```

## 0-1 BFS（双端队列BFS）

0-1 BFS 适用于解决边权值只有 $0$ 和 $1$（或者能够转化为这种情况）的最短路问题。

[这里](https://codeforces.com/blog/entry/22276)有一篇很详细的教程。

0-1 BFS 的时间复杂度为：$O(E+V)$ ，其中 $V$ 和 $E$ 分别是图中的节点数和边数。

### LC1368. 至少有一条有效路径的最小代价 { data-toc-label='LC1368. 至少有一条有效路径...' }

???+ note "问题描述"
    给你一个 `m x n (1≤m,n≤100)` 的网格图 `grid` 。<br>
    每个单元格的数字表示当前单元格允许的前进方向，`1, 2, 3, 4` 分别对应右、左、下、上四个方向。<br>
    每个单元格的数字可以修改一次，修改的代价为 `1` 。<br>
    求使得网格图中至少有一条从 `(0, 0)` 到 `(m-1, n-1)` 的有效路径的最小修改代价。

    在 [LeetCode主站](https://leetcode.com/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid "困难：2069") 查看该题。

??? info "解题思路"
    根据题意建立有向图，每个单元格作为一个节点。<br>
    每个节点向与其相邻的单元格节点连出有向边，如果方向相同则边权为 `0` ；方向不同则边权为 `1` 。<br>
    问题转化为求从 `(0, 0)` 到 `(m-1, n-1)` 的最短路径。

    **方法一：0-1 BFS**

    === "Go"
        ```go
        --8<-- "shortest-path/01bfs/go/lc1368_1.go"
        ```

    **方法二：Dijkstra 算法**

    Dijkstra 算法不是本题的重点，[这里](#dijkstra-算法)是更适合使用Dijkstra 算法的题目。

    === "Go"
        ```go
        --8<-- "shortest-path/01bfs/go/lc1368_2.go"
        ```

### LC2290. 移除障碍物的最小数目 { data-toc-label='LC2290. 移除障碍物的最小数...' }

???+ note "问题描述"
    给你一个 `m x n` `1≤m,n≤1e5` 且 `2≤m*n≤1e5` 的网格图 `grid` 。每个单元格都是两个值之一：

    * `0` 表示一个空单元格；
    * `1` 表示一个可以移除的障碍物。
    
    你可以向上、下、左、右移动，从一个空单元格移动到另一个空单元格。<br>
    现在你需要从左上角 `(0, 0)` 移动到右下角 `(m-1, n-1)` ，返回需要移除的障碍物的最小数目。<br>
    题目保证左上角 `(0, 0)` 和右下角 `(m-1, n-1)` 为空单元格。

    在 [LeetCode主站](https://leetcode.com/problems/minimum-obstacle-removal-to-reach-corner "Hard")
    或 [力扣中文社区](https://leetcode.cn/problems/minimum-obstacle-removal-to-reach-corner "困难：2138") 查看该题。

??? info "解题思路"
    根据题意建立有向图，每个单元格作为一个节点。<br>
    每个节点向与其相邻的单元格节点连出有向边，相邻单元格的值即为边权。<br>
    问题转化为求从 `(0, 0)` 到 `(m-1, n-1)` 的最短路径。

    **方法一：0-1 BFS**

    === "Go"
        ```go
        --8<-- "shortest-path/01bfs/go/lc2290_1.go"
        ```

    **方法二：Dijkstra 算法**

    Dijkstra 算法不是本题的重点，[这里](#dijkstra-算法)是更适合使用Dijkstra 算法的题目。

    === "Go"
        ```go
        --8<-- "shortest-path/01bfs/go/lc2290_2.go"
        ```
