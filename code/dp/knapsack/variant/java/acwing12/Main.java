import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(System.out);

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

        // 倒序转移最大总价值
        int[][] dp = new int[n + 1][m + 1];
        for (int i = n - 1; i >= 0; i--) {
            int v = items[i][0], w = items[i][1];
            for (int j = 0; j <= m; j++) {
                dp[i][j] = dp[i + 1][j];
                if (j >= v) {
                    dp[i][j] = Math.max(dp[i][j], dp[i + 1][j - v] + w);
                }
            }
        }

        // 正序输出字典序最小的方案
        int left = m;
        for (int i = 0; i < n && left > 0; i++) {
            int v = items[i][0], w = items[i][1];
            if (left >= v && dp[i][left] == dp[i + 1][left - v] + w) {
                out.printf("%d ", i + 1);
                left -= v;
            }
        }
        out.flush();
        out.close();
        in.close();
    }
}
