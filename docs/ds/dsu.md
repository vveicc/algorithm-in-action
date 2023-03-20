# 并查集

## 并查集基础

### CF1157E: Minimum Array

???+ note "问题描述"
    第一行输入 $n(1≤n≤2e5)$ ，第二行和第三行分别输入两个长为 $n$ 的数组 $a$ 和 $b$ ，元素范围：$[0,n-1]$ 。<br>
    你可以重排数组 $b$ 。<br>
    还有一个长为 $n$ 的数组 $c$ ，其中 $c[i] = (a[i] + b[i]) \% n$ 。<br>
    输出字典序最小的数组 $c$ 。

    在 [Codeforces](https://codeforces.com/problemset/problem/1157/E "1700")
    或 [洛谷](https://www.luogu.com.cn/problem/CF1157E "普及+/提高") 查看该题。

??? info "解题思路"
    对于 $a[i]$ ，需要去数组 $b$ 中找 $(n-a[i])%n$ ，如果不存在就找更大的值，如果找不到，就再从 $0$ 开始找。<br>
    直接暴力查找会超时，有多种优化方法：

    - 类似 [Java TreeMap](https://docs.oracle.com/javase/8/docs/api/java/util/TreeMap.html) 这样的平衡树，只维护存在的；
    - 并查集，如果 $x$ 不存在，则把 $x$ 和 $x+1$ 合并，这样可以快速找到下一个存在的。

    **方法一：并查集**

    === "Go"
        ```go
        --8<-- "dsu/basic/go/cf1157e_1.go"
        ```
    === "Java"
        ```java
        --8<-- "dsu/basic/java/cf1157e_1/Main.java"
        ```

    **方法二：有序字典**

    === "Java"
        ```java
        --8<-- "dsu/basic/java/cf1157e_2/Main.java"
        ```
