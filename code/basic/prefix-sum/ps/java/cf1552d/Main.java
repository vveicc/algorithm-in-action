import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(System.out);

    public static void main(String[] args) throws Exception {
        int t = Integer.parseInt(in.readLine());
        for (; t > 0; t--) {
            int n = Integer.parseInt(in.readLine());
            StringTokenizer tokenizer = new StringTokenizer(in.readLine());
            int[] a = new int[n];
            for (int i = 0; i < n; i++) {
                a[i] = Integer.parseInt(tokenizer.nextToken());
            }
            boolean ans = false;
            Set<Integer> set = new HashSet<>();
            set.add(0);
            for (int i = (1 << n) - 1; i > 0; i--) {
                int sum = 0;
                for (int j = 0; j < n; j++) {
                    if (((i >> j) & 1) == 1) {
                        sum += a[j];
                    }
                }
                if (set.contains(sum)) {
                    ans = true;
                    break;
                } else {
                    set.add(sum);
                }
            }
            out.println(ans ? "YES" : "NO");
        }
        out.flush();
        out.close();
        in.close();
    }
}
