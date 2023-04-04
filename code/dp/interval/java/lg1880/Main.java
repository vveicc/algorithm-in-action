import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(System.out);

    public static void main(String[] args) throws Exception {
        int n = Integer.parseInt(in.readLine());
        int[] a = new int[n];
        int m = n << 1;
        int[] s = new int[m + 1];
        StringTokenizer at = new StringTokenizer(in.readLine());
        for (int i = 0; i < n; i++) {
            a[i] = Integer.parseInt(at.nextToken());
            s[i + 1] = s[i] + a[i];
        }
        for (int i = 0; i < n; i++) {
            int j = i + n;
            s[j + 1] = s[j] + a[i];
        }
        int[][] mn = new int[m][m];
        int[][] mx = new int[m][m];
        for (int i = m - 1; i >= 0; i--) {
            for (int j = i + 1; j < i + n && j < m; j++) {
                mn[i][j] = (int) 1e6;
                for (int k = i; k < j; k++) {
                    mn[i][j] = Math.min(mn[i][j], mn[i][k] + mn[k + 1][j]);
                    mx[i][j] = Math.max(mx[i][j], mx[i][k] + mx[k + 1][j]);
                }
                // 合并成一堆
                int score = s[j + 1] - s[i];
                mn[i][j] += score;
                mx[i][j] += score;
            }
        }
        int x = mn[0][n - 1], y = mx[0][n - 1];
        for (int i = 1; i < n; i++) {
            x = Math.min(x, mn[i][i + n - 1]);
            y = Math.max(y, mx[i][i + n - 1]);
        }
        out.printf("%d%n%d", x, y);
        out.flush();
        out.close();
        in.close();
    }
}