import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));

    public static void main(String[] args) throws Exception {
        StringTokenizer nm = new StringTokenizer(in.readLine());
        int n = Integer.parseInt(nm.nextToken());
        int m = Integer.parseInt(nm.nextToken());
        // groups[i][j]表示从第i个dq中取出j+1个数的最大和
        int[][] groups = new int[n][];
        for (int i = 0; i < n; i++) {
            StringTokenizer tokenizer = new StringTokenizer(in.readLine());
            int k = Integer.parseInt(tokenizer.nextToken());
            int[] s = new int[k + 1]; // 前缀和
            for (int j = 0; j < k; j++) {
                s[j + 1] = s[j] + Integer.parseInt(tokenizer.nextToken());
            }
            groups[i] = new int[k];
            groups[i][k - 1] = s[k];
            for (int j = 1; j < k; j++) {
                for (int l = 0; l <= j; l++) { // 在dq头部取l个数，尾部取j-l个数
                    groups[i][j - 1] = Math.max(groups[i][j - 1], s[l] + s[k] - s[k - (j - l)]);
                }
            }
        }

        int t = 0;
        int[] dp = new int[m + 1];
        for (int[] group : groups) { // 枚举分组
            int k = group.length;
            t += k;
            for (int i = Math.min(t, m); i > 0; i--) { // 枚举背包容量
                for (int j = 0; j < k && j < i; j++) { // 枚举分组内的每一件物品
                    // 使用j+1的容量容纳group[j]
                    dp[i] = Math.max(dp[i], dp[i - 1 - j] + group[j]);
                }
            }
        }
        System.out.println(dp[m]);
        in.close();
    }
}
