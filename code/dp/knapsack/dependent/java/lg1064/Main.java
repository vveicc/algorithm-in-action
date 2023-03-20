import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(System.out);

    public static void main(String[] args) throws Exception {
        StringTokenizer nm = new StringTokenizer(in.readLine());
        int n = Integer.parseInt(nm.nextToken()) / 10;
        int m = Integer.parseInt(nm.nextToken());
        int[][] items = new int[m + 1][3];
        List<Integer>[] og = new List[m + 1];
        for (int i = 1; i <= m; i++) {
            og[i] = new ArrayList<>();
        }
        for (int i = 1; i <= m; i++) {
            StringTokenizer vpq = new StringTokenizer(in.readLine());
            items[i][0] = Integer.parseInt(vpq.nextToken()) / 10;
            items[i][1] = Integer.parseInt(vpq.nextToken()) * items[i][0];
            items[i][2] = Integer.parseInt(vpq.nextToken());
            if (items[i][2] != 0) {
                og[items[i][2]].add(i);
            }
        }
        int[] dp = new int[n + 1];
        for (int x = 1; x <= m; x++) { // 枚举主件与其附件组成的分组
            if (items[x][2] == 0) {
                int vx = items[x][0];
                for (int i = n; i >= vx; i--) { // 枚举背包容量
                    int c = og[x].size();
                    for (int s = (1 << c) - 1; s >= 0; s--) { // 枚举分组内的物品（附件选择状态）
                        int v = vx, w = items[x][1]; // 选择主件x
                        for (int j = 0; j < c; j++) {
                            if (((s >> j) & 1) == 1) { // 选择状态s表示的附件
                                int y = og[x].get(j);
                                v += items[y][0];
                                w += items[y][1];
                            }
                        }
                        if (i >= v) {
                            dp[i] = Math.max(dp[i], dp[i - v] + w);
                        }
                    }
                }
            }
        }

        out.println(dp[n] * 10);
        out.flush();
        out.close();
        in.close();
    }
}
