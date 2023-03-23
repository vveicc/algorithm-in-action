import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(System.out);

    public static void main(String[] args) throws Exception {
        int T = Integer.parseInt(in.readLine());
        for (; T > 0; T--) {
            in.readLine(); // 空行
            StringTokenizer nm = new StringTokenizer(in.readLine());
            int n = Integer.parseInt(nm.nextToken());
            int m = Integer.parseInt(nm.nextToken());
            char[][] a = new char[n][m];
            for (int i = 0; i < n; i++) {
                a[i] = in.readLine().toCharArray();
            }
            char[] s = in.readLine().toCharArray();

            // 预处理a中所有长度为2和3的子串的位置
            int[][] f2 = new int[100][3];
            int[][] f3 = new int[1000][3];
            for (int i = n-1; i >= 0; i--) {
                if (m > 1) {
                    f2[fi2(a[i], m - 2)] = new int[] { m - 1, m, i + 1 };
                }
                for (int j = m - 3; j >= 0; j--) {
                    f2[fi2(a[i], j)] = new int[] { j + 1, j + 2, i + 1 };
                    f3[fi3(a[i], j)] = new int[] { j + 1, j + 3, i + 1 };
                }
            }

            // f[i]表示s[i:]是否可以划分
            boolean[] f = new boolean[m + 1];
            f[m] = true;
            if (m > 1) {
                f[m - 2] = f2[fi2(s, m - 2)][0] != 0;
            }
            for (int i = m - 3; i >= 0; i--) {
                f[i] = (f[i + 2] && f2[fi2(s, i)][0] != 0) || (f[i + 3] && f3[fi3(s, i)][0] != 0);
            }

            if (!f[0]) {
                out.println(-1);
            } else {
                int idx = 0;
                int[][] arr = new int[m >> 1][3];
                for (int i = 0; i < m;) {
                    int x = fi2(s, i);
                    if (f[i + 2] && f2[x][0] != 0) {
                        arr[idx++] = f2[x];
                        i += 2;
                    } else {
                        arr[idx++] = f3[fi3(s, i)];
                        i += 3;
                    }
                }
                out.println(idx);
                for (int i = 0; i < idx; i++) {
                    out.printf("%d %d %d%n", arr[i][0], arr[i][1], arr[i][2]);
                }
            }
        }
        out.flush();
        out.close();
        in.close();
    }

    public static int fi2(char[] s, int i) {
        return 10 * (s[i] & 15) + (s[i + 1] & 15);
    }

    public static int fi3(char[] s, int i) {
        return 100 * (s[i] & 15) + 10 * (s[i + 1] & 15) + (s[i + 2] & 15);
    }
}
