import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(System.out);

    public static void main(String[] args) throws Exception {
        StringTokenizer nm = new StringTokenizer(in.readLine());
        int n = Integer.parseInt(nm.nextToken());
        int m = Integer.parseInt(nm.nextToken());
        int[] dp = new int[m + 1];
        for (; n > 0; n--) { // 枚举分组
            int c = Integer.parseInt(in.readLine());
            int lower = 100;
            int[][] items = new int[c][2];
            for (int i = 0; i < c; i++) {
                StringTokenizer vw = new StringTokenizer(in.readLine());
                items[i][0] = Integer.parseInt(vw.nextToken());
                items[i][1] = Integer.parseInt(vw.nextToken());
                lower = Math.min(lower, items[i][0]);
            }
            for (int i = m; i >= lower; i--) { // 枚举背包容量
                for (int[] item : items) { // 枚举分组内的每一件物品
                    int v = item[0], w = item[1];
                    if (i >= v) {
                        dp[i] = Math.max(dp[i], dp[i - v] + w);
                    }
                }
            }
        }
        out.println(dp[m]);
        out.flush();
        out.close();
        in.close();
    }
}
