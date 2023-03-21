import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(System.out);

    static int mod = (int) 1e9 + 7;

    public static void main(String[] args) throws Exception {
        StringTokenizer nm = new StringTokenizer(in.readLine());
        int n = Integer.parseInt(nm.nextToken());
        int m = Integer.parseInt(nm.nextToken());
        int[][] items = new int[n][2];
        for (int[] item : items) {
            StringTokenizer vw = new StringTokenizer(in.readLine());
            item[0] = Integer.parseInt(vw.nextToken());
            item[1] = Integer.parseInt(vw.nextToken());
        }

        int[] dp = new int[m + 1];
        // cnt[i]表示容量为i的背包装入总价值为dp[i]的物品的方案数
        int[] cnt = new int[m + 1];
        for (int i = 0; i <= m; i++) {
            cnt[i] = 1; // 初始化为1，表示什么都不装的方案
        }
        for (int[] item : items) {
            int v = item[0], w = item[1];
            for (int i = m; i >= v; i--) {
                int x = dp[i - v] + w;
                if (x > dp[i]) {
                    dp[i] = x;
                    cnt[i] = cnt[i - v];
                } else if (x == dp[i]) {
                    cnt[i] = (cnt[i] + cnt[i - v]) % mod;
                }
            }
        }
        out.println(cnt[m]);
        out.flush();
        out.close();
        in.close();
    }
}
