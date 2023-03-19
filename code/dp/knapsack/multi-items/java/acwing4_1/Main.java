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
        for (int[] item : items) { // 枚举每种物品
            int v = item[0], w = item[1], c = item[2];
            for (; c > 0; c--) { // 枚举该种商品的每一件，转化为0-1背包
                for (int j = m; j >= v; j--) { // 枚举背包容量
                    dp[j] = Math.max(dp[j], dp[j - v] + w);
                }
            }
        }
        System.out.println(dp[m]);
        in.close();
    }
}
