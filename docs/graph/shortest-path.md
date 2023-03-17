# 最短路

## Dijkstra 算法

Dijkstra 算法适合用来求出无负权边图中的单源最短路径。其中：

* **无负权边** 表示图中所有边的权值必须为非负数；
* **单源最短路径** 表示 Dijkstra 算法可以求出从某一个节点到其余所有节点的最短路径。

使用二叉堆实现的优先队列优化的 Dijkstra 算法的时间复杂度为：$O((E+V)\log V)$ ，其中 $V$ 和 $E$ 分别是图中的节点数和边数。

## 0-1 BFS（双端队列BFS）

0-1 BFS 适用于解决边权值只有 $0$ 和 $1$（或者能够转化为这种情况）的最短路问题。

[这里](https://codeforces.com/blog/entry/22276)有一篇很详细的教程。

0-1 BFS 的时间复杂度为：$O(E+V)$ ，其中 $V$ 和 $E$ 分别是图中的节点数和边数。

### LC1368: 使网格图至少有一条有效路径的最小代价

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

### LC2290: 到达角落需要移除障碍物的最小数目

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
