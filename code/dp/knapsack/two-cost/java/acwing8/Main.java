import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(System.out);

    public static void main(String[] args) throws Exception {
        StringTokenizer ncm = new StringTokenizer(in.readLine());
        int N = Integer.parseInt(ncm.nextToken());
        int C = Integer.parseInt(ncm.nextToken());
        int M = Integer.parseInt(ncm.nextToken());
        int[][] items = new int[N][3];
        for (int i = 0; i < N; i++) {
            StringTokenizer vmw = new StringTokenizer(in.readLine());
            items[i][0] = Integer.parseInt(vmw.nextToken());
            items[i][1] = Integer.parseInt(vmw.nextToken());
            items[i][2] = Integer.parseInt(vmw.nextToken());
        }

        int[][] dp = new int[C + 1][M + 1];
        for (int[] item : items) { // 枚举物品
            int v = item[0], m = item[1], w = item[2];
            for (int i = C; i >= v; i--) { // 枚举第一维度费用
                for (int j = M; j >= m; j--) { // 枚举第二维度费用
                    dp[i][j] = Math.max(dp[i][j], dp[i - v][j - m] + w);
                }
            }
        }
        out.println(dp[C][M]);
        out.flush();
        out.close();
        in.close();
    }
}
