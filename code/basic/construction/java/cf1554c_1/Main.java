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
            int k = Integer.parseInt(tokenizer.nextToken()) + 1;
            if (n >= k) {
                out.println(0);
            } else {
                int x = 0;
                int i = 1;
                for (; k >> i != 0; i++) {
                }
                for (; i >= 0; i--) {
                    int ni = (n >> i) & 1;
                    int ki = (k >> i) & 1;
                    if (ni == 0 && ki == 1) {
                        x |= 1 << i;
                    } else if (ni == 1 && ki == 0) {
                        break;
                    }
                }
                out.println(x);
            }
        }
        out.flush();
        out.close();
        in.close();
    }
}
