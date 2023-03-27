import java.io.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(System.out);

    public static void main(String[] args) throws Exception {
        int T = Integer.parseInt(in.readLine());
        for (; T > 0; T--) {
            char[] s = in.readLine().toCharArray();
            long ans = 0;
            // prev[0]表示前一个与索引奇偶性相同的非'?'字符（'0'或'1'）的索引
            // prev[1]表示前一个与索引奇偶性不同的非'?'字符（'0'或'1'）的索引
            int[] prev = new int[] { -1, -1 };
            for (int i = 0, n = s.length; i < n; i++) {
                if (s[i] != '?') {
                    prev[(i & 1) ^ (s[i] & 1)] = i;
                }
                // (prev[0], i]区间内的非'?'字符（'0'或'1'）与索引奇偶性均不同
                // (prev[1], i]区间内的非'?'字符（'0'或'1'）与索引奇偶性均相同
                ans += i - Math.min(prev[0], prev[1]);
            }
            out.println(ans);
        }
        out.flush();
        out.close();
        in.close();
    }
}
