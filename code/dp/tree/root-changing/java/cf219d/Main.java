import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(new BufferedWriter(new OutputStreamWriter(System.out)));

    static List<Integer>[] ug;
    static List<Integer> xs = new ArrayList<>();
    static int min;

    public static void main(String[] args) throws Exception {
        int n = Integer.parseInt(in.readLine());
        ug = new List[n + 1];
        for (int i = 1; i <= n; i++) {
            ug[i] = new ArrayList<>();
        }
        for (; n > 1; n--) {
            StringTokenizer tokenizer = new StringTokenizer(in.readLine());
            int v = Integer.parseInt(tokenizer.nextToken());
            int w = Integer.parseInt(tokenizer.nextToken());
            ug[v].add(w << 1);     // 最低位 0 表示正向
            ug[w].add(v << 1 | 1); // 最低位 1 表示反向
        }
        // 第一遍DFS计算以节点1为根时，需要反向的边的数量
        min = dfs1(1, 0);
        // 第二遍DFS通过换根DP计算每个节点为根时，需要反向的边的数量
        dfs(1, 0, min);
        Collections.sort(xs);
        out.println(min);
        for (int x : xs) {
            out.printf("%d ", x);
        }
        out.flush();
        out.close();
        in.close();
    }

    public static int dfs1(int x, int fa) {
        int count = 0;
        for (int y : ug[x]) {
            if ((y >> 1) != fa) {
                count += (y & 1) + dfs1(y >> 1, x);
            }
        }
        return count;
    }

    public static void dfs(int x, int fa, int c) {
        if (c < min) {
            min = c;
            xs.clear();
            xs.add(x);
        } else if (c == min) {
            xs.add(x);
        }
        for (int y : ug[x]) {
            if ((y >> 1) != fa) {
                // 如果x->y为正向，则以节点y为根时，需要反向的边的数量+1；如果为反向则-1
                dfs(y >> 1, x, c - ((y & 1) << 1) + 1);
            }
        }
    }
}
