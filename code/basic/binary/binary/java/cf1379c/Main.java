import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(new BufferedWriter(new OutputStreamWriter(System.out)));

    public static void main(String[] args) throws Exception {
        int t = Integer.parseInt(in.readLine());
        for (; t > 0; t--) {
            StringTokenizer nm = new StringTokenizer(in.readLine());
            int n = Integer.parseInt(nm.nextToken());
            int m = Integer.parseInt(nm.nextToken());
            Integer[] as = new Integer[m];
            int[][] abs = new int[m][2];
            for (int i = 0; i < m; i++) {
                StringTokenizer tokenizer = new StringTokenizer(in.readLine());
                abs[i][0] = Integer.parseInt(tokenizer.nextToken());
                abs[i][1] = Integer.parseInt(tokenizer.nextToken());
                as[i] = abs[i][0];
            }

            // 按照a降序排序
            Arrays.sort(as, (o1, o2) -> Integer.compare(o2, o1));
            // 前缀和
            long[] s = new long[m + 1];
            for (int i = 0; i < m; i++) {
                s[i + 1] = s[i] + as[i];
            }
            long ans = 0;
            for (int[] ab : abs) {
                int a = ab[0];
                int b = ab[1];
                // 前i个物品的a大于当前物品的b，如果当前物品取多个（多于1个），则前i个物品需要每个取1个
                int i = Arrays.binarySearch(as, b, (o1, o2) -> Integer.compare(o2, o1));
                if (i < 0) {
                    i = -i - 1;
                }
                if (i >= n) {
                    ans = Math.max(ans, s[n]);
                } else if (a > b) { // 前i个物品包含当前物品
                    ans = Math.max(ans, s[i] + ((long) b) * (n - i));
                } else { // 前i个物品不包含当前物品
                    ans = Math.max(ans, s[i] + a + ((long) b) * (n - i - 1));
                }
            }
            out.println(ans);
            if (t > 1) {
                in.readLine(); // 每组数据由空行分隔
            }
        }
        out.flush();
        out.close();
        in.close();
    }
}
