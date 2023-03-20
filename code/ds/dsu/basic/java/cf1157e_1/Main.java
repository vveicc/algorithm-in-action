import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(System.out);

    public static void main(String[] args) throws Exception {
        int n = Integer.parseInt(in.readLine());
        int[] a = new int[n];
        int[] b = new int[n + 1];
        StringTokenizer at = new StringTokenizer(in.readLine());
        StringTokenizer bt = new StringTokenizer(in.readLine());
        for (int i = 0; i < n; i++) {
            a[i] = Integer.parseInt(at.nextToken());
            int x = Integer.parseInt(bt.nextToken());
            b[x]++;
        }

        fa = new int[n + 1];
        for (int i = n, next = 0; i >= 0; i--) {
            if (b[i] != 0) {
                next = i;
            }
            fa[i] = next;
        }

        for (int x : a) {
            int y = find(n - x);
            if (--b[y] == 0) {
                union(y, y + 1); // 指向下一个未使用的数字
            }
            out.printf("%d ", (x + y) % n);
        }
        out.flush();
        out.close();
        in.close();
    }

    // 并查集
    static int[] fa;

    public static int find(int x) {
        if (fa[x] != x) {
            fa[x] = find(fa[x]); // 路径压缩
        }
        return fa[x];
    }

    public static void union(int f, int t) {
        f = find(f);
        t = find(t);
        if (f != t) {
            fa[f] = t;
        }
    }
}
