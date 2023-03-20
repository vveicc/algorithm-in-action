import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(System.out);

    static int n, m, root;
    static List<Integer>[] og;
    static int[][] items;
    // dp[x][i]表示选择以x为根的子树中的物品，且容量不超过i时能获得的最大价值
    static int[][] dp;

    public static void dfs(int x) {
        int v = items[x][0], w = items[x][1];
        for (int i = v; i <= m; i++) {
            dp[x][i] = w; // 物品x必须选，初始化为物品x的价值
        }
        for (int y : og[x]) {
            dfs(y); // 先计算子树y
            for (int i = m; i >= v; i--) { // 枚举背包容量
                for (int j = 0; j <= i - v; j++) { // 枚举留给子树y的背包容量
                    dp[x][i] = Math.max(dp[x][i], dp[x][i - j] + dp[y][j]);
                }
            }
        }
    }

    public static void main(String[] args) throws Exception {
        StringTokenizer nm = new StringTokenizer(in.readLine());
        n = Integer.parseInt(nm.nextToken());
        m = Integer.parseInt(nm.nextToken());
        og = new List[n + 1];
        for (int i = 1; i <= n; i++) {
            og[i] = new ArrayList<>();
        }
        items = new int[n + 1][2];
        for (int i = 1; i <= n; i++) {
            StringTokenizer vwp = new StringTokenizer(in.readLine());
            items[i][0] = Integer.parseInt(vwp.nextToken());
            items[i][1] = Integer.parseInt(vwp.nextToken());
            int p = Integer.parseInt(vwp.nextToken());
            if (p == -1) {
                root = i;
            } else {
                og[p].add(i);
            }
        }
        dp = new int[n + 1][m + 1];
        dfs(root);
        out.println(dp[root][m]);
        out.flush();
        out.close();
        in.close();
    }
}
