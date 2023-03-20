import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(System.out);

    public static void main(String[] args) throws Exception {
        int n = Integer.parseInt(in.readLine());
        int[] a = new int[n];
        TreeMap<Integer, Integer> b = new TreeMap<>();
        StringTokenizer at = new StringTokenizer(in.readLine());
        StringTokenizer bt = new StringTokenizer(in.readLine());
        for (int i = 0; i < n; i++) {
            a[i] = Integer.parseInt(at.nextToken());
            int x = Integer.parseInt(bt.nextToken());
            b.compute(x, (k, v) -> v == null ? 1 : v + 1);
        }
        for (int x : a) {
            Integer y = b.ceilingKey(n - x);
            if (y == null) {
                y = b.ceilingKey(0);
            }
            b.compute(y, (k, v) -> v == 1 ? null : v - 1);
            out.printf("%d ", (x + y) % n);
        }
        out.flush();
        out.close();
        in.close();
    }
}
