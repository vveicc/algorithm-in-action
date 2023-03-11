import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(new BufferedWriter(new OutputStreamWriter(System.out)));

    public static void main(String[] args) throws Exception {
        int n = Integer.parseInt(in.readLine());
        int[] rs = new int[n];
        int[][] arr = new int[n][];
        for (int i = 0; i < n; i++) {
            StringTokenizer tokenizer = new StringTokenizer(in.readLine());
            int l = Integer.parseInt(tokenizer.nextToken());
            int r = Integer.parseInt(tokenizer.nextToken());
            arr[i] = new int[] { l, r, i };
            rs[i] = r;
        }

        // 按照左端点降序排序
        Arrays.sort(arr, (a, b) -> Integer.compare(b[0], a[0]));

        int[] ans = new int[n];
        // 离散化树状数组计算前序有多少个区间右端点小于当前区间右端点
        Arrays.sort(rs);
        BIT bit = new BIT(n + 1);
        for (int[] lri : arr) {
            int x = Arrays.binarySearch(rs, lri[1]);
            ans[lri[2]] = bit.query(x);
            bit.update(x + 1, 1);
        }

        for (int c : ans) {
            out.println(c);
        }
        out.flush();
        out.close();
        in.close();
    }

    // 树状数组
    private static class BIT {

        int n;
        int[] arr;

        BIT(int n) {
            this.n = n;
            this.arr = new int[n];
        }

        // 单点更新
        void update(int i, int v) {
            for (; i < n; i += (i & -i)) {
                arr[i] += v;
            }
        }

        // 前缀查询 [1, i]
        int query(int i) {
            int sum = 0;
            for (; i > 0; i &= (i - 1)) {
                sum += arr[i];
            }
            return sum;
        }
    }
}
