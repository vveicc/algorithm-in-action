import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));

    static final int inf = (int) -1e9;

    static int d;
    static boolean[] monster;
    static List<Integer>[] ug;
    static Tuple[] dis;
    static int cnt;

    public static void main(String[] args) throws Exception {
        StringTokenizer nmd = new StringTokenizer(in.readLine());
        int n = Integer.parseInt(nmd.nextToken());
        int m = Integer.parseInt(nmd.nextToken());
        d = Integer.parseInt(nmd.nextToken());
        monster = new boolean[n + 1];
        StringTokenizer ms = new StringTokenizer(in.readLine());
        for (int i = 0; i < m; i++) {
            int o = Integer.parseInt(ms.nextToken());
            monster[o] = true;
        }
        ug = new List[n + 1];
        for (int i = 1; i <= n; i++) {
            ug[i] = new ArrayList<>();
        }
        for (int i = 1; i < n; i++) {
            StringTokenizer uv = new StringTokenizer(in.readLine());
            int u = Integer.parseInt(uv.nextToken());
            int v = Integer.parseInt(uv.nextToken());
            ug[u].add(v);
            ug[v].add(u);
        }
        dis = new Tuple[n + 1];
        // 第一遍DFS以节点1为根，计算每个节点与其子树中的最远怪物距离、次远怪物距离及最远怪物距离所在子树节点
        dfs1(1, 0);
        // 第二遍DFS通过换根DP计算每个节点为根时，与其子树中的最远怪物距离，判断是否可以作为传送门
        dfs(1, 0, inf);
        System.out.println(cnt);
    }

    public static int dfs1(int x, int fa) {
        // 距离初始化为inf，方便处理子树中没有怪物的情况
        int f = inf, s = inf, o = 0;
        for (int y : ug[x]) {
            if (y != fa) {
                int yf = dfs1(y, x) + 1;
                if (yf > f) {
                    o = y;
                    s = f;
                    f = yf;
                } else if (yf > s) {
                    s = yf;
                }
            }
        }
        dis[x] = new Tuple(f, s, o);
        if (f < 0 && monster[x]) {
            return 0;
        } else {
            return f;
        }
    }

    public static void dfs(int x, int fa, int df) {
        if (df <= d) {
            Tuple dx = dis[x];
            if (dx.f <= d) {
                cnt++;
            }
            if (df < 0 && monster[x]) {
                df = 0;
            }
            for (int y : ug[x]) {
                if (y != fa) {
                    if (y == dx.o) {
                        dfs(y, x, Math.max(df, dx.s) + 1);
                    } else {
                        dfs(y, x, Math.max(df, dx.f) + 1);
                    }
                }
            }
        }
    }

    private static class Tuple {
        int f; // 子树中最远怪物的距离
        int s; // 子树中次远怪物的距离
        int o; // 最远怪物距离所在子树节点

        Tuple(int f, int s, int o) {
            this.f = f;
            this.s = s;
            this.o = o;
        }
    }
}
