import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(System.out);

    public static void main(String[] args) throws Exception {
        StringTokenizer nm = new StringTokenizer(in.readLine());
        int n = Integer.parseInt(nm.nextToken());
        int m = Integer.parseInt(nm.nextToken());
        char[][] a = new char[n][];
        for (int i = 0; i < n; i++) {
            a[i] = in.readLine().toCharArray();
        }

        int ans = 0;
        int[] lec = new int[n];
        for (int j = 0; j < m; j++) { // 枚举右边界
            // lec[i]表示第i行第j列往左连续相同字符的个数
            for (int i = 0; i < n; i++) {
                if (j != 0 && a[i][j] == a[i][j - 1]) {
                    lec[i]++;
                } else {
                    lec[i] = 1;
                }
            }
            next: for (int i = 0; i < n;) {
                int i0 = i; // 中间h行的第一行
                int mn = lec[i]; // 左侧最短同色长度
                // 处理中间h行
                for (i++; i < n && a[i][j] == a[i0][j]; i++) {
                    mn = Math.min(mn, lec[i]);
                }
                // 此时i指向后h行的第一行
                int h = i - i0;
                if (i0 < h || i + h > n) {
                    continue; // 前面或后面不够h行
                } else {
                    // 处理前h行
                    for (int k = --i0; k > i0 - h; k--) {
                        if (a[k][j] != a[i0][j]) {
                            continue next; // 前h行不同色
                        }
                        mn = Math.min(mn, lec[k]);
                    }
                    // 处理后h行
                    for (int k = i; k < i + h; k++) {
                        if (a[k][j] != a[i][j]) {
                            continue next; // 后h行不同色
                        }
                        mn = Math.min(mn, lec[k]);
                    }
                    ans += mn;
                }
            }
        }
        out.println(ans);
        out.flush();
        out.close();
        in.close();
    }
}
