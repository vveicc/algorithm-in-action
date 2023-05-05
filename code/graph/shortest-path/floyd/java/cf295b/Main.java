import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(System.out);

    public static void main(String[] args) throws IOException {
        int n = Integer.parseInt(in.readLine());
        int[][] a = new int[n][n];
        for (int i = 0; i < n; i++) {
            StringTokenizer at = new StringTokenizer(in.readLine());
            for (int j = 0; j < n; j++) {
                a[i][j] = Integer.parseInt(at.nextToken());
            }
        }
        long[] b = new long[n];
        StringTokenizer bt = new StringTokenizer(in.readLine());
        for (int i = 0; i < n; i++) {
            b[i] = Integer.parseInt(bt.nextToken()) - 1;
        }
        boolean[] vis = new boolean[n];
        for (int i = n - 1; i >= 0; i--) {
            int k = (int) b[i];
            vis[k] = true;
            b[i] = 0;
            for (int x = 0; x < n; x++) {
                for (int y = 0; y < n; y++) {
                    a[x][y] = Math.min(a[x][y], a[x][k] + a[k][y]);
                    if (vis[x] && vis[y]) {
                        b[i] += a[x][y];
                    }
                }
            }
        }
        for (long x : b) {
            out.printf("%d ", x);
        }
        out.flush();
        out.close();
        in.close();
    }
}
