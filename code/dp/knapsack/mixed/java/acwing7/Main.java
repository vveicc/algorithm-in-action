import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(System.out);

    public static void main(String[] args) throws Exception {
        StringTokenizer nm = new StringTokenizer(in.readLine());
        int n = Integer.parseInt(nm.nextToken());
        int m = Integer.parseInt(nm.nextToken());
        int[][] items = new int[n][3];
        for (int i = 0; i < n; i++) {
            StringTokenizer vwc = new StringTokenizer(in.readLine());
            items[i][0] = Integer.parseInt(vwc.nextToken());
            items[i][1] = Integer.parseInt(vwc.nextToken());
            items[i][2] = Integer.parseInt(vwc.nextToken());
        }

        // 滚动数组优化
        int i = 1;
        int[][] f = new int[2][m + 1];
        // 单调递减索引队列
        int[] q = new int[m + 1];
        for (int[] item : items) {
            int v = item[0], w = item[1], c = item[2];
            int cv = c * v;
            if (c == -1 || c == 1) { // 0-1 背包
                int[] cur = f[i & 1];
                for (int j = m; j >= v; j--) {
                    cur[j] = Math.max(cur[j], cur[j - v] + w);
                }
            } else if (c == 0 || cv >= m) { // 完全背包
                int[] cur = f[i & 1];
                for (int j = v; j <= m; j++) {
                    cur[j] = Math.max(cur[j], cur[j - v] + w);
                }
            } else { // 多重背包，使用单调队列优化
                // 滚动数组
                int[] cur = f[++i & 1], pre = f[(i & 1) ^ 1];
                // 枚举余数
                for (int j = 0; j < v; j++) {
                    int head = 0, tail = -1;
                    // 枚举余数j的方案：j+v, j+2*v, j+3*v, ...
                    for (int k = j; k <= m; k += v) {
                        // 初始化为上次计算的结果
                        cur[k] = pre[k];
                        // 将不在窗口内的索引从队首出队
                        for (; head <= tail && k - q[head] > cv; head++) {
                        }
                        // 使用队首的最大值更新结果
                        if (head <= tail) {
                            // 使用k-q[head]的容量装入当前物品
                            cur[k] = Math.max(cur[k], pre[q[head]] + (k - q[head]) / v * w);
                        }
                        // 将队尾的价值小于当前方案的索引移除，保证单调递减
                        for (; head <= tail && pre[q[tail]] + (k - q[tail]) / v * w <= pre[k]; tail--) {
                        }
                        // 将当前方案的索引入队尾
                        q[++tail] = k;
                    }
                }
            }
        }

        out.println(f[i & 1][m]);
        out.flush();
        out.close();
        in.close();
    }
}
