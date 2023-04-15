import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(System.out);

    public static void main(String[] args) throws IOException {
        int T = Integer.parseInt(in.readLine());
        while (T-- > 0) {
            StringTokenizer nk = new StringTokenizer(in.readLine());
            int n = Integer.parseInt(nk.nextToken());
            int k = Integer.parseInt(nk.nextToken());
            char[] s = in.readLine().toCharArray();
            int r = Math.min(k, Integer.toBinaryString(n - k + 1).length());
            int l = k - r;
            int mask = (1 << r) - 1;
            boolean[] has = new boolean[1 << r];
            int i = 0, l1 = 0, rt = 0;
            for (; i < l; i++) {
                l1 += s[i] & 1;
            }
            for (; i < k; i++) {
                rt = (rt << 1) | (s[i] & 1);
            }
            has[rt] = l1 == l;
            for (; i < n; i++) {
                l1 += (s[i - r] & 1) - (s[i - k] & 1);
                rt = ((rt << 1) | (s[i] & 1)) & mask;
                if (l1 == l) {
                    has[rt] = true;
                }
            }
            rt = mask;
            for (; rt != -1 && has[rt]; rt--) {
            }
            if (rt == -1) {
                out.println("NO");
            } else {
                rt ^= mask;
                String ans = Integer.toBinaryString(rt);
                out.println("YES");
                for (i = k - ans.length(); i != 0; i--) {
                    out.print(0);
                }
                out.println(ans);
            }
        }
        out.flush();
        out.close();
        in.close();
    }
}
