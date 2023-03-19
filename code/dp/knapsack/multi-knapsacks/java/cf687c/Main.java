import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(System.out);

    public static void main(String[] args) throws Exception {
        StringTokenizer nk = new StringTokenizer(in.readLine());
        int n = Integer.parseInt(nk.nextToken());
        int k = Integer.parseInt(nk.nextToken());
        int[] c = new int[n];
        StringTokenizer tokenizer = new StringTokenizer(in.readLine());
        for (int i = 0; i < n; i++) {
            c[i] = Integer.parseInt(tokenizer.nextToken());
        }

        // dp[i][j1][j2]表示从前i个数中是否能选出两个不相交的子集，其中子集1的元素和为j1，子集2的元素和为j2
        boolean[][] dp = new boolean[k + 1][k + 1];
        dp[0][0] = true;

        Arrays.sort(c); // 排序优化枚举效率

        int s = 0, ms = 0;
        for (int x : c) {
            s += x;
            ms = Math.min(s, k); // 子集元素和从当前最大和开始枚举
            for (int j1 = ms; j1 >= 0; j1--) {
                for (int j2 = ms; j2 >= 0; j2--) {
                    // 不选择x，或者选择x放入子集1，或者选择x放入子集2
                    dp[j1][j2] |= (j1 >= x && dp[j1 - x][j2]) || (j2 >= x && dp[j1][j2 - x]);
                }
            }
        }

        int cnt = 0;
        int[] ans = new int[k + 1];
        for (int x = 0; x <= k; x++) {
            if (dp[x][k - x]) {
                ans[cnt++] = x;
            }
        }
        out.println(cnt);
        for (int i = 0; i < cnt; i++) {
            out.printf("%d ", ans[i]);
        }
        out.flush();
        out.close();
        in.close();
    }
}
