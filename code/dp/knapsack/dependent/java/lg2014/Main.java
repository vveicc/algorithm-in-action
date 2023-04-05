import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(System.out);

    public static void main(String[] args) throws Exception {
        // 所有课程的依赖关系组成森林结构
        // 新增一门编号为0且学分为0的课程作为根节点，将依赖关系转为一棵树
        // 从编号从0开始的 n+1 门课程中选择 m+1 门课程
        StringTokenizer nm = new StringTokenizer(in.readLine());
        n = Integer.parseInt(nm.nextToken()) + 1;
        m = Integer.parseInt(nm.nextToken()) + 1;

        s = new int[n];
        og = new List[n];
        for (int i = 0; i < n; i++) {
            og[i] = new ArrayList<Integer>();
        }
        for (int i = 1; i < n; i++) {
            StringTokenizer t = new StringTokenizer(in.readLine());
            int k = Integer.parseInt(t.nextToken());
            s[i] = Integer.parseInt(t.nextToken());
            og[k].add(i);
        }

        // dp[x][i]表示选择以x为根的子树中的i门课程所能获得的最大学分
        dp = new int[n][m + 1];
        dfs(0);
        out.println(dp[0][Math.min(n, m)]);
        out.flush();
        out.close();
        in.close();
    }

    static int n, m;
    static int[] s;
    static List<Integer>[] og;
    static int[][] dp;

    public static int dfs(int x) {
        int cx = 1;
        dp[x][1] = s[x]; // 仅选择课程x
        for (int y : og[x]) {
            int cy = dfs(y);
            for (int i = Math.min(cx, m); i != 0; i--) { // 从已合并过的子树中选择i门课程
                for (int j = 1; j <= cy && i + j <= m; j++) { // 从子树y中选择j门课程
                    dp[x][i + j] = Math.max(dp[x][i + j], dp[x][i] + dp[y][j]);
                }
            }
            cx += cy;
        }
        return cx;
    }
}