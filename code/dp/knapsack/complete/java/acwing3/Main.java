import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));

    public static void main(String[] args) throws Exception {
        StringTokenizer nm = new StringTokenizer(in.readLine());
        int n = Integer.parseInt(nm.nextToken());
        int m = Integer.parseInt(nm.nextToken());
        int[][] items = new int[n][2];
        for (int i = 0; i < n; i++) {
            StringTokenizer tokenizer = new StringTokenizer(in.readLine());
            items[i][0] = Integer.parseInt(tokenizer.nextToken());
            items[i][1] = Integer.parseInt(tokenizer.nextToken());
        }
        int[] dp = new int[m + 1];
        for (int[] item : items) { // 枚举物品
            int v = item[0], w = item[1];
            for (int j = v; j <= m; j++) { // 枚举背包容量
                dp[j] = Math.max(dp[j], dp[j - v] + w);
            }
        }
        System.out.println(dp[m]);
        in.close();
    }
}