import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));

    public static void main(String[] args) throws Exception {
        StringTokenizer nm = new StringTokenizer(in.readLine());
        int n = Integer.parseInt(nm.nextToken());
        int m = Integer.parseInt(nm.nextToken());
        int[][] items = new int[n][3];
        for (int i = 0; i < n; i++) {
            StringTokenizer tokenizer = new StringTokenizer(in.readLine());
            items[i][0] = Integer.parseInt(tokenizer.nextToken());
            items[i][1] = Integer.parseInt(tokenizer.nextToken());
            items[i][2] = Integer.parseInt(tokenizer.nextToken());
        }

        int[] dp = new int[m + 1];
        for (int[] item : items) {
            int v = item[0], w = item[1], c = item[2];
            if (c * v >= m) { // 转化为完全背包
                for (int j = v; j <= m; j++) {
                    dp[j] = Math.max(dp[j], dp[j - v] + w);
                }
            } else { // 不能转化为完全背包，使用二进制分组优化
                // 每个分组看作独立的一件物品，转化为0-1背包，组合分组可以组合出[1, c]范围内的任一件数
                for (int k = 1; k <= c; k <<= 1) {
                    c -= k;
                    int kv = k * v, kw = k * w;
                    for (int j = m; j >= kv; j--) {
                        dp[j] = Math.max(dp[j], dp[j - kv] + kw);
                    }
                }
                if (c != 0) {
                    int cv = c * v, cw = c * w;
                    for (int j = m; j >= cv; j--) {
                        dp[j] = Math.max(dp[j], dp[j - cv] + cw);
                    }
                }
            }
        }
        System.out.println(dp[m]);
    }
}
