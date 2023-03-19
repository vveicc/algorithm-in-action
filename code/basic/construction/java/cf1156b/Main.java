import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(System.out);

    public static void main(String[] args) throws Exception {
        int t = Integer.parseInt(in.readLine());
        for (; t > 0; t--) {
            char[] s = in.readLine().toCharArray();
            // 排序
            Arrays.sort(s);
            int[] idx = new int[2];
            char[][] a = new char[2][s.length];
            for (char c : s) {
                int j = c & 1;
                a[j][idx[j]] = c; // 按序奇偶分组
                idx[j]++;
            }
            int m = idx[0], n = idx[1];
            String x = new String(a[0], 0, m);
            String y = new String(a[1], 0, n);
            if (m == 0) {
                out.println(y);
            } else if (n == 0) {
                out.println(x);
            } else if (Math.abs(a[0][m - 1] - a[1][0]) != 1) {
                out.printf("%s%s%n", x, y); // x+y
            } else if (Math.abs(a[1][n - 1] - a[0][0]) != 1) {
                out.printf("%s%s%n", y, x); // y+x
            } else {
                // 当且仅当字符串s仅由2个或3个相邻字母组成时，无法完成构造
                out.println("No answer");
            }
        }
        out.flush();
        out.close();
        in.close();
    }
}
