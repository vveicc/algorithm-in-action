import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(new BufferedWriter(new OutputStreamWriter(System.out)));

    public static void main(String[] args) throws Exception {
        int t = Integer.parseInt(in.readLine());
        for (; t > 0; t--) {
            StringTokenizer tokenizer = new StringTokenizer(in.readLine());
            int n = Integer.parseInt(tokenizer.nextToken());
            int m = Integer.parseInt(tokenizer.nextToken());
            if (n > m) {
                out.println(0);
            } else {
                int i = 0;
                int x = n ^ m;
                if ((n & x) == 0) {
                    for (int y = n | x; (y >> i & 1) == 1; i++) {
                    }
                    out.println((x >> i | 1) << i); // 通过第四种情况修改x
                } else {
                    for (int y = n & x; (y >> i) != 0; i++) {
                    }
                    out.println(x >> i << i); // 通过第三种情况修改x
                }
            }
        }
        out.flush();
        out.close();
        in.close();
    }
}
